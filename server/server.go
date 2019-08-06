package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/ecator/gomeeting/config"
	"github.com/ecator/gomeeting/db"
	"github.com/ecator/gomeeting/fun"
	"github.com/ecator/gomeeting/loggers"
	"github.com/ecator/gomeeting/msg"
	"github.com/julienschmidt/httprouter"
)

var (
	conf        *config.Config
	logger      *loggers.Logger
	dbConn      *db.DB
	frontDir    string
	onlineUsers map[string]uint32
)

// StartServer starts server
func StartServer(listenAddr string, listenPort uint, frontendPath string, configPath string, logFile string) {
	var err error
	frontDir = frontendPath
	if logger, err = loggers.New(logFile, ""); err != nil {
		panic(err)
	}
	if conf, err = config.ParseConfig(configPath); err != nil {
		logger.Error(err.Error())
		logger.Close()
		os.Exit(1)
	}
	if dbConn, err = db.New(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password); err != nil {
		logger.Error(err.Error())
		logger.Close()
		os.Exit(1)
	}
	// connect to db
	if err = dbConn.Open(); err != nil {
		logger.Error(err.Error())
		logger.Close()
		os.Exit(1)
	}

	// set routers
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(handleNotFound)
	// handler
	router.GET("/", decorateAuth(handleIndex))
	router.GET("/login", decorateAuth(handleLogin))
	router.POST("/login", decorateAuth(handleLogin))
	router.GET("/logout", decorateAuth(handleLogout))
	router.GET("/user", decorateAuth(handleShowUser))
	router.GET("/password", decorateAuth(handlePassword))

	// the handler for api
	router.POST("/api/:target", decorateAuth(apiPost))
	router.DELETE("/api/:target/:id", decorateAuth(apiDel))
	router.DELETE("/api/:target", decorateAuth(apiDel))
	router.PUT("/api/:target/:id", decorateAuth(apiPut))
	router.GET("/api/:target/:id", decorateAuth(apiGet))
	router.GET("/api/:target", decorateAuth(apiGet))

	// static resources
	router.ServeFiles("/js/*filepath", http.Dir(filepath.Join(frontendPath, "js")))
	router.ServeFiles("/img/*filepath", http.Dir(filepath.Join(frontendPath, "img")))
	router.ServeFiles("/css/*filepath", http.Dir(filepath.Join(frontendPath, "css")))

	// make the super token to control all
	token := fun.GetMd5Str(time.Now().String() + "root")
	onlineUsers = make(map[string]uint32)
	onlineUsers[token] = 0
	// start http server
	logger.Info("Starting service...")
	errCh := make(chan error)
	var wg sync.WaitGroup
	listenAddr = listenAddr + ":" + strconv.Itoa(int(listenPort))
	go func() {
		wg.Add(1)
		errCh <- http.ListenAndServe(listenAddr, router)
		wg.Done()
	}()
	select {
	case err := <-errCh:
		logger.Error(err.Error())
		logger.Close()
		os.Exit(1)
	case <-time.After(time.Second * 3):
		logger.Info("Listening at " + listenAddr)
		logger.Info("The super token is " + token)
		wg.Wait()

	}
}

func decorateAuth(fn func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ret := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var (
			resp   jsonResp
			status int
		)
		logger.Info(r.Method + " " + r.RequestURI + " from " + r.RemoteAddr)
		// not neet to auth
		switch r.RequestURI {
		case "/login", "/logout":
			fn(w, r, ps)
			return
		}
		// auth
		if _, err := getThisUserID(r); err == nil {
			// auth ok
			fn(w, r, ps)
			return
		}
		// auth error
		status = 9000
		resp = jsonResp{status, msg.GetMsg(status)}
		// redirect
		redirectLocation(w, "/login")
		// response
		responseJSON(w, &resp)

	}
	return ret
}

func responseFile(w http.ResponseWriter, filePath string) {
	if b, err := ioutil.ReadFile(filePath); err == nil {
		w.Write(b)
	} else {
		logger.Error(err.Error())
	}
}

func responseJSON(w http.ResponseWriter, resp *jsonResp) {
	if b, err := json.Marshal(resp); err == nil {
		w.Header().Set("Content-type", "application/json")
		w.Write(b)
	} else {
		logger.Error(err.Error())
	}
}

func redirectLocation(w http.ResponseWriter, urlstr string) {
	w.Header().Set("Location", urlstr)
	w.WriteHeader(302)
}

func getThisUserID(r *http.Request) (uint32, error) {
	var status int
	if c, err := r.Cookie("auth"); err == nil {
		if id, has := onlineUsers[c.Value]; has == true {
			return id, nil
		}
		// auth expired
		status = 9001
	} else {
		// no auth
		status = 9000
	}
	return 0, errors.New(msg.GetMsg(status))
}
