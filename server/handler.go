package server

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/ecator/gomeeting/fun"
	"github.com/ecator/gomeeting/msg"
	"github.com/julienschmidt/httprouter"
)

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	responseFile(w, filepath.Join(frontDir, "404.html"))
}

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responseFile(w, filepath.Join(frontDir, "index.html"))
}

func handlePassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// show password modify page
	responseFile(w, filepath.Join(frontDir, "password.html"))
}

func handleAdmin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// show admin page
	if id1, err := getThisUserID(r); err != nil {
		handleNotFound(w, r)
	} else if id1 != 0 {
		handleNotFound(w, r)
	} else {
		responseFile(w, filepath.Join(frontDir, "admin.html"))
	}
}

func handleShowUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// show user profile page
	responseFile(w, filepath.Join(frontDir, "user.html"))
}

func handleBrowerErr(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(403)
	responseFile(w, filepath.Join(frontDir, "brower_err.html"))
}

func handleLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp   jsonResp
		status int
	)
	if c, err := r.Cookie("auth"); err == nil {
		delete(onlineUsers, c.Value)
		// delete cookie
		c.Expires = time.Now().Add(time.Hour * 24 * -365)
		http.SetCookie(w, c)
		redirectLocation(w, "/")
	}
	status = 0
	resp = jsonResp{status, msg.GetMsg(1000, "logout")}
	responseJSON(w, &resp)
}

func handleLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp   jsonResp
		status int
		u      *user
		u1     *user
		err    error
	)
	switch r.Method {
	case "GET":
		// show login page
		responseFile(w, filepath.Join(frontDir, "login.html"))
	case "POST":
		// check auth
		u = new(user)
		parseReqToObj(r, ps, u)
		if u.Username == "" || u.Pw == "" {
			// login parameter error
			status = 9004
			resp = jsonResp{status, msg.GetMsg(status, "login parameter")}
		} else {
			//logger.Info("get user(" + u.Username + "," + u.Pw + ")")
			u1 = new(user)
			u1.Username = u.Username
			if err = selectObjByUsername(u1); err == nil {
				if u1.Pw == u.Pw {
					//login ok
					token := fun.GetMd5Str(time.Now().String() + u1.Username)
					onlineUsers[token] = u1.ID
					// add cookie
					cookie := new(http.Cookie)
					cookie.Name = "auth"
					cookie.Value = token
					cookie.Expires = time.Now().Add(time.Hour * 24 * 365 * 100)
					http.SetCookie(w, cookie)
					status = 0
					resp = jsonResp{status, msg.GetMsg(1000, "login")}
				} else {
					// password wrong
					status = 9002
					resp = jsonResp{status, msg.GetMsg(status, "password")}
				}
			} else {
				// user not found
				status = 9003
				resp = jsonResp{status, msg.GetMsg(status, "user")}
			}
		}
		// response
		responseJSON(w, &resp)
	}
}
