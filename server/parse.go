package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ecator/gomeeting/fun"

	"github.com/julienschmidt/httprouter"
)

func parseReqToObj(r *http.Request, ps httprouter.Params, o interface{}) {
	var (
		id        string
		username  string
		password  string
		level     string
		orgID     string
		name      string
		email     string
		roomID    string
		startTime string
		endTime   string
		maker     string
		memo      string
		makeDate  string
		message   string
		b         []byte
	)

	// get from params first
	id = ps.ByName("id")
	username = ps.ByName("username")
	password = ps.ByName("password")
	level = ps.ByName("level")
	orgID = ps.ByName("org_id")
	name = ps.ByName("name")
	email = ps.ByName("email")
	roomID = ps.ByName("room_id")
	startTime = ps.ByName("start_time")
	endTime = ps.ByName("end_time")
	maker = ps.ByName("maker")
	memo = ps.ByName("memo")
	makeDate = ps.ByName("makeDate")
	message = ps.ByName("message")

	// check if not found then get from formvalue
	if id == "" {
		id = r.FormValue("id")
	}
	if username == "" {
		username = r.FormValue("username")
	}
	if password == "" {
		password = r.FormValue("password")
	}
	if level == "" {
		level = r.FormValue("level")
	}
	if orgID == "" {
		orgID = r.FormValue("org_id")
	}
	if name == "" {
		name = r.FormValue("name")
	}
	if email == "" {
		email = r.FormValue("email")
	}
	if roomID == "" {
		roomID = r.FormValue("room_id")
	}
	if startTime == "" {
		startTime = r.FormValue("start_time")
	}
	if endTime == "" {
		endTime = r.FormValue("end_time")
	}
	if maker == "" {
		maker = r.FormValue("maker")
	}
	if memo == "" {
		memo = r.FormValue("memo")
	}
	if makeDate == "" {
		makeDate = r.FormValue("make_date")
	}
	if message == "" {
		message = r.FormValue("message")
	}
	// set value to obj

	if id != "" {
		fun.SetUint32ByName(o, "ID", fun.Str2Uint32(id))
	}
	if username != "" {
		fun.SetStrByName(o, "Username", username)
	}
	if password != "" {
		fun.SetStrByName(o, "Pw", password)
	}
	if level != "" {
		fun.SetUint32ByName(o, "Level", fun.Str2Uint32(level))
	}
	if orgID != "" {
		fun.SetUint32ByName(o, "OrgID", fun.Str2Uint32(orgID))
	}
	if name != "" {
		fun.SetStrByName(o, "Name", name)
	}
	if email != "" {
		fun.SetStrByName(o, "Email", email)
	}

	if roomID != "" {
		fun.SetUint32ByName(o, "RoomID", fun.Str2Uint32(roomID))
	}
	if startTime != "" {
		fun.SetUint32ByName(o, "StartTime", fun.Str2Uint32(startTime))
	}
	if endTime != "" {
		fun.SetUint32ByName(o, "EndTime", fun.Str2Uint32(endTime))
	}
	if maker != "" {
		fun.SetUint32ByName(o, "Maker", fun.Str2Uint32(maker))
	}
	if memo != "" {
		fun.SetStrByName(o, "Memo", memo)
	}
	if makeDate != "" {
		fun.SetUint32ByName(o, "MakeDate", fun.Str2Uint32(makeDate))
	}
	if message != "" {
		fun.SetStrByName(o, "Message", message)
	}

	// try to parse from body
	b, _ = ioutil.ReadAll(r.Body)
	json.Unmarshal(b, o)
}
