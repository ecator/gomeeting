package server

import (
	"errors"
	"fmt"

	"github.com/ecator/gomeeting/msg"

	"github.com/ecator/gomeeting/fun"
)

func insertObj(o interface{}) error {
	var (
		sql string
		err error
	)
	switch o.(type) {
	case *user:
		if fun.GetStrByName(o, "Username") == "" {
			err = errors.New(msg.GetMsg(9007, "username"))
		} else if fun.GetStrByName(o, "Pw") == "" {
			err = errors.New(msg.GetMsg(9007, "password"))
		} else {
			sql = fmt.Sprintf("insert into user values (%d,'%s','%s',%d,%d,'%s','%s')", fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Username"), fun.GetStrByName(o, "Pw"), fun.GetUint32ByName(o, "Level"), fun.GetUint32ByName(o, "OrgID"), fun.GetStrByName(o, "Name"), fun.GetStrByName(o, "Email"))
		}
	case *org:
		if fun.GetStrByName(o, "Name") == "" {
			err = errors.New(msg.GetMsg(9007, "name"))
		} else {
			sql = fmt.Sprintf("insert into org values (%d,'%s')", fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name"))
		}
	case *room:
		if fun.GetStrByName(o, "Name") == "" {
			err = errors.New(msg.GetMsg(9007, "name"))
		} else {
			sql = fmt.Sprintf("insert into room values (%d,'%s')", fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name"))
		}
	case *meeting:
		if fun.GetUint32ByName(o, "RoomID") == 0 {
			err = errors.New(msg.GetMsg(9010, "room_id"))
		} else if fun.GetUint32ByName(o, "StartTime") >= fun.GetUint32ByName(o, "EndTime") {
			err = errors.New(msg.GetMsg(9011, "end_time", "start_time"))
		} else if fun.GetUint32ByName(o, "Maker") == 0 {
			err = errors.New(msg.GetMsg(9010, "maker"))
		} else if fun.GetStrByName(o, "Memo") == "" {
			err = errors.New(msg.GetMsg(9007, "memo"))
		} else if fun.GetUint32ByName(o, "MakeDate") < 19000101 || fun.GetUint32ByName(o, "MakeDate") > 99991231 {
			err = errors.New(msg.GetMsg(9012, "make_date", "19000101", "99991231"))
		} else if _, err = selectMeetings(fmt.Sprintf("select * from meeting where room_id=%[1]d and  start_time <%[2]d and end_time>=%[2]d and make_date=%[3]d", fun.GetUint32ByName(o, "RoomID"), fun.GetUint32ByName(o, "EndTime"), fun.GetUint32ByName(o, "MakeDate"))); err == nil {
			err = errors.New(msg.GetMsg(9009, "end_time"))
		} else if _, err = selectMeetings(fmt.Sprintf("select * from meeting where room_id=%[1]d and  start_time <=%[2]d and end_time>%[2]d and make_date=%[3]d", fun.GetUint32ByName(o, "RoomID"), fun.GetUint32ByName(o, "StartTime"), fun.GetUint32ByName(o, "MakeDate"))); err == nil {
			err = errors.New(msg.GetMsg(9009, "start_time"))
		} else if _, err = selectMeetings(fmt.Sprintf("select * from meeting where room_id=%d and  start_time >%d and end_time<%d and make_date=%d", fun.GetUint32ByName(o, "RoomID"), fun.GetUint32ByName(o, "StartTime"), fun.GetUint32ByName(o, "EndTime"), fun.GetUint32ByName(o, "MakeDate"))); err == nil {
			err = errors.New(msg.GetMsg(9009, "start_time - end_time"))
		} else {
			sql = fmt.Sprintf("insert into meeting values (%d,%d,%d,%d,'%s',%d)", fun.GetUint32ByName(o, "RoomID"), fun.GetUint32ByName(o, "StartTime"), fun.GetUint32ByName(o, "EndTime"), fun.GetUint32ByName(o, "Maker"), fun.GetStrByName(o, "Memo"), fun.GetUint32ByName(o, "MakeDate"))
		}
	default:
		err = errors.New(msg.GetMsg(9004, "type"))
	}

	if sql != "" {
		_, err = dbConn.Exec(sql)
	}
	return err
}

func deleteObj(o interface{}) error {
	var (
		sql string
		err error
	)
	switch o.(type) {
	case *user:
		sql = fmt.Sprintf("delete from user where id=%d", fun.GetUint32ByName(o, "ID"))
	case *org:
		sql = fmt.Sprintf("delete from org where id=%d", fun.GetUint32ByName(o, "ID"))
	case *room:
		sql = fmt.Sprintf("delete from room where id=%d", fun.GetUint32ByName(o, "ID"))
	case *meeting:
		sql = fmt.Sprintf("delete from meeting where room_id=%d and start_time=%d and end_time=%d and make_date=%d", fun.GetUint32ByName(o, "RoomID"), fun.GetUint32ByName(o, "StartTime"), fun.GetUint32ByName(o, "EndTime"), fun.GetUint32ByName(o, "MakeDate"))
	default:
		err = errors.New(msg.GetMsg(9004, "type"))
	}

	if sql != "" {
		_, err = dbConn.Exec(sql)
	}
	return err
}

func updateObj(o interface{}) error {
	var (
		sql string
		err error
	)
	switch o.(type) {
	case *user:
		if fun.GetStrByName(o, "Username") == "" {
			err = errors.New("username can not be empty")
		} else if fun.GetStrByName(o, "Pw") == "" {
			err = errors.New("password can not be empty")
		} else {
			sql = fmt.Sprintf("update user set username='%s',pw='%s',level=%d,org_id=%d,name='%s',email='%s' where id=%d", fun.GetStrByName(o, "Username"), fun.GetStrByName(o, "Pw"), fun.GetUint32ByName(o, "Level"), fun.GetUint32ByName(o, "OrgID"), fun.GetStrByName(o, "Name"), fun.GetStrByName(o, "Email"), fun.GetUint32ByName(o, "ID"))
		}
	case *org:
		if fun.GetStrByName(o, "Name") == "" {
			err = errors.New("name can not be empty")
		} else {
			sql = fmt.Sprintf("update org set name='%s' where id=%d", fun.GetStrByName(o, "Name"), fun.GetUint32ByName(o, "ID"))
		}
	case *room:
		if fun.GetStrByName(o, "Name") == "" {
			err = errors.New("name can not be empty")
		} else {
			sql = fmt.Sprintf("update room set name='%s' where id=%d", fun.GetStrByName(o, "Name"), fun.GetUint32ByName(o, "ID"))
		}
	default:
		err = errors.New(msg.GetMsg(9004, "type"))
	}
	if sql != "" {
		_, err = dbConn.Exec(sql)
	}
	return err
}

func selectObjByID(o interface{}) error {
	var (
		sql string
		err error
		ro  interface{}
	)
	switch o.(type) {
	case *user:
		sql = fmt.Sprintf("select * from user where id=%d", fun.GetUint32ByName(o, "ID"))
		ro, err = selectUser(sql)
	case *org:
		sql = fmt.Sprintf("select * from org where id=%d", fun.GetUint32ByName(o, "ID"))
		ro, err = selectOrg(sql)
	case *room:
		sql = fmt.Sprintf("select * from room where id=%d", fun.GetUint32ByName(o, "ID"))
		ro, err = selectRoom(sql)
	default:
		err = errors.New(msg.GetMsg(9004, "type"))
	}
	if err == nil {
		fun.SetByObj(o, ro)
	}
	return err
}

func selectObjByUsername(o interface{}) error {
	var (
		sql string
		err error
		ro  interface{}
	)
	switch o.(type) {
	case *user:
		sql = fmt.Sprintf("select * from user where username='%s'", fun.GetStrByName(o, "Username"))
		ro, err = selectUser(sql)
	default:
		err = errors.New(msg.GetMsg(9004, "type"))
	}
	if err == nil {
		fun.SetByObj(o, ro)
	}
	return err
}

func getNewObjID(o interface{}) uint32 {
	var sql string
	switch o.(type) {
	case user, *user:
		sql = "select ifnull(max(id),999)+1 as newid from user"
	case org, *org:
		sql = "select ifnull(max(id),999)+1 as newid from org"
	case room, *room:
		sql = "select ifnull(max(id),999)+1 as newid from room"
	default:
		return 1000
	}

	if r, err := dbConn.Query(sql); err == nil {
		return fun.Str2Uint32(r[0]["newid"])
	}
	return 1000
}

func makeJSONrespMeetings(makeDate uint32) []jsonRespMeeting {
	var (
		ms  []jsonRespMeeting
		m   jsonRespMeeting
		sql string
	)
	sql = `
	SELECT 
		 meeting.room_id as room_id
		,room.name as room_name
		,meeting.start_time as start_time
		,meeting.end_time as end_time
		,meeting.make_date as make_date
		,meeting.maker as maker
		,meeting.memo as memo
		,user.username as user_username
		,user.name as user_name
		,user.email as user_email
		,user.level as user_level
		,user.org_id as user_org_id
		,org.name as org_name
		FROM meeting
			inner join user on meeting.maker = user.id
			inner join room on meeting.room_id=room.id
			inner join org on user.org_id=org.id
		where meeting.make_date=%d
	`
	sql = fmt.Sprintf(sql, makeDate)
	ms = []jsonRespMeeting{}
	m = jsonRespMeeting{}
	if r, err := dbConn.Query(sql); err == nil {
		for i := 0; i < len(r); i++ {
			m.Room.ID = fun.Str2Uint32(r[i]["room_id"])
			m.Room.Name = r[i]["room_name"]
			m.StartTime = fun.Str2Uint32(r[i]["start_time"])
			m.EndTime = fun.Str2Uint32(r[i]["end_time"])
			m.MakeDate = fun.Str2Uint32(r[i]["make_date"])
			m.Maker.ID = fun.Str2Uint32(r[i]["maker"])
			m.Maker.Username = r[i]["user_username"]
			m.Maker.Name = r[i]["user_name"]
			m.Maker.Level = fun.Str2Uint32(r[i]["user_level"])
			m.Maker.Org.ID = fun.Str2Uint32(r[i]["user_org_id"])
			m.Maker.Org.Name = r[i]["org_name"]
			m.Maker.Email = r[i]["user_email"]
			m.Memo = r[i]["memo"]
			ms = append(ms, m)
		}
	}
	return ms
}

func makeJSONrespRooms() []jsonRespRoom {
	var (
		sql string
		r   jsonRespRoom
		rs  []jsonRespRoom
	)
	sql = "select * from room"
	if ret, err := dbConn.Query(sql); err == nil {
		r = jsonRespRoom{}
		rs = []jsonRespRoom{}
		for i := 0; i < len(ret); i++ {
			r.ID = fun.Str2Uint32(ret[i]["id"])
			r.Name = ret[i]["name"]
			rs = append(rs, r)
		}
	}
	return rs
}
func makeJSONrespOrgs() []jsonRespOrg {
	var (
		sql string
		r   jsonRespOrg
		rs  []jsonRespOrg
	)
	sql = "select * from org"
	if ret, err := dbConn.Query(sql); err == nil {
		r = jsonRespOrg{}
		rs = []jsonRespOrg{}
		for i := 0; i < len(ret); i++ {
			r.ID = fun.Str2Uint32(ret[i]["id"])
			r.Name = ret[i]["name"]
			rs = append(rs, r)
		}
	}
	return rs
}
func makeJSONrespUsers() []jsonRespUser {
	var (
		sql string
		r   jsonRespUser
		rs  []jsonRespUser
	)
	sql = `
	SELECT
		user.id as id
	   ,user.username as username
	   ,user.level as level
	   ,user.name as name
	   ,user.email as email
	   ,user.org_id as org_id
	   ,org.name as org_name
	   FROM user
	   left join org on user.org_id = org.id
	`
	if ret, err := dbConn.Query(sql); err == nil {
		r = jsonRespUser{}
		rs = []jsonRespUser{}
		for i := 0; i < len(ret); i++ {
			r.ID = fun.Str2Uint32(ret[i]["id"])
			r.Username = ret[i]["username"]
			r.Level = fun.Str2Uint32(ret[i]["level"])
			r.Name = ret[i]["name"]
			r.Email = ret[i]["email"]
			r.Username = ret[i]["username"]
			r.Org.ID = fun.Str2Uint32(ret[i]["org_id"])
			r.Org.Name = ret[i]["org_name"]
			rs = append(rs, r)
		}
	}
	return rs
}
