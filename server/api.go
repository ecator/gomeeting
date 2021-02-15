package server

import (
	"net/http"

	"github.com/ecator/gomeeting/fun"

	"github.com/ecator/gomeeting/msg"

	"github.com/julienschmidt/httprouter"
)

func apiPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp             jsonResp
		respRoom         jsonRespRoom
		respUser         jsonRespUser
		respOrg          jsonRespOrg
		respNotification jsonRespNotification
		respMeeting      jsonRespMeeting
		objOrg           *org
		objRoom          *room
		objUser          *user
		status           int
		o                interface{}
		target           string
	)
	target = ps.ByName("target")
	switch target {
	case "meeting":
		o = new(meeting)
	case "user":
		o = new(user)
	case "room":
		o = new(room)
	case "org":
		o = new(org)
	case "notification":
		o = new(notification)
	default:
		status = 9004
		resp = jsonResp{status, msg.GetMsg(status, "target")}
	}
	if status == 0 {
		id1, _ := getThisUserID(r)
		parseReqToObj(r, ps, o)
		// check
		switch target {
		case "meeting":
			if id1 > 0 && fun.GetUint32ByName(o, "Maker") != id1 {
				fun.SetUint32ByName(o, "Maker", id1)
			}
		case "user", "room", "org", "notification":
			// must root
			if id1 != 0 {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			} else if fun.GetStrByName(o, "Pw") != "" {
				fun.SetStrByName(o, "Pw", fun.GetMd5Str(fun.GetStrByName(o, "Pw")))
			}
		}
		if status == 0 {
			// insert to db
			if status == 0 {
				fun.SetUint32ByName(o, "ID", getNewObjID(o))
				if err := insertObj(o); err == nil {
					// success
					status = 0
					switch target {
					case "user":
						respUser = jsonRespUser{}
						respUser.ID = fun.GetUint32ByName(o, "ID")
						respUser.Level = fun.GetUint32ByName(o, "Level")
						respUser.Username = fun.GetStrByName(o, "Username")
						respUser.Name = fun.GetStrByName(o, "Name")
						respUser.Email = fun.GetStrByName(o, "Email")
						respUser.Org.ID = fun.GetUint32ByName(o, "OrgID")
						respUser.Org.Name = ""
						// search org_name
						objOrg = new(org)
						objOrg.ID = respUser.Org.ID
						if err := selectObjByID(objOrg); err == nil {
							respUser.Org.Name = objOrg.Name
						} else {
							logger.Warn(err.Error())
						}
						status = 0
						resp = jsonResp{status, respUser}
					case "meeting":
						respMeeting = jsonRespMeeting{}
						respMeeting.Room.ID = fun.GetUint32ByName(o, "RoomID")
						respMeeting.StartTime = fun.GetUint32ByName(o, "StartTime")
						respMeeting.EndTime = fun.GetUint32ByName(o, "EndTime")
						respMeeting.Maker.ID = fun.GetUint32ByName(o, "Maker")
						respMeeting.MakeDate = fun.GetUint32ByName(o, "MakeDate")
						respMeeting.Memo = fun.GetStrByName(o, "Memo")
						// search room_name
						objRoom = new(room)
						objRoom.ID = respMeeting.Room.ID
						if err := selectObjByID(objRoom); err == nil {
							respMeeting.Room.Name = objRoom.Name
						} else {
							logger.Warn(err.Error())
						}
						// search user
						objUser = new(user)
						objUser.ID = respMeeting.Maker.ID
						if err := selectObjByID(objUser); err == nil {
							respMeeting.Maker.Username = objUser.Username
							respMeeting.Maker.Level = objUser.Level
							respMeeting.Maker.Org.ID = objUser.OrgID
							respMeeting.Maker.Name = objUser.Name
							respMeeting.Maker.Email = objUser.Email
						} else {
							logger.Warn(err.Error())
						}
						// search org_name
						objOrg = new(org)
						objOrg.ID = respMeeting.Maker.Org.ID
						if err := selectObjByID(objOrg); err == nil {
							respMeeting.Maker.Org.Name = objOrg.Name
						} else {
							logger.Warn(err.Error())
						}
						resp = jsonResp{status, respMeeting}
					case "room":
						respRoom = jsonRespRoom{fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name")}
						resp = jsonResp{status, respRoom}
					case "org":
						respOrg = jsonRespOrg{fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name")}
						resp = jsonResp{status, respOrg}
					case "notification":
						respNotification = jsonRespNotification{fun.GetStrByName(o, "Message")}
						resp = jsonResp{status, respNotification}
					}
				} else {
					// fail
					logger.Error(err.Error())
					status = 9006
					resp = jsonResp{status, err.Error()}
				}
			}
		}
	}
	// response
	responseJSON(w, &resp)
}

func apiDel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp   jsonResp
		status int
		o      interface{}
		target string
	)
	target = ps.ByName("target")
	switch target {
	case "meeting":
		o = new(meeting)
	case "user":
		o = new(user)
	case "room":
		o = new(room)
	case "org":
		o = new(org)
	case "notification":
		o = new(notification)
	default:
		status = 9004
		resp = jsonResp{status, msg.GetMsg(status, "target")}
	}
	if status == 0 {
		id1, _ := getThisUserID(r)
		parseReqToObj(r, ps, o)
		switch target {
		case "meeting":
			if id1 > 0 && fun.GetUint32ByName(o, "Maker") != id1 {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
		case "user", "room", "org":
			if id1 == 0 {
				if fun.GetUint32ByName(o, "ID") == 0 {
					// id wrong
					status = 9004
					resp = jsonResp{status, msg.GetMsg(status, "id")}
				}
			} else {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
		case "notification":
			// must root
			if id1 != 0 {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
		}
		// delete from db
		if status == 0 {
			if err := deleteObj(o); err == nil {
				// success
				status = 0
				resp = jsonResp{status, msg.GetMsg(1000, "delete "+target)}
			} else {
				// fail
				logger.Error(err.Error())
				status = 9006
				resp = jsonResp{status, err.Error()}
			}
		}
	}
	// response
	responseJSON(w, &resp)
}

func apiPut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp     jsonResp
		respUser jsonRespUser
		status   int
		o        interface{}
		objUser  *user
		objOrg   *org
		objRoom  *room
		target   string
	)
	target = ps.ByName("target")
	switch target {
	case "user":
		o = new(user)
	case "room":
		o = new(room)
	case "org":
		o = new(org)
	default:
		status = 9004
		resp = jsonResp{status, msg.GetMsg(status, "target")}
	}
	if status == 0 {
		id1, _ := getThisUserID(r)
		parseReqToObj(r, ps, o)
		switch target {
		case "user":
			if id1 > 0 && fun.GetUint32ByName(o, "ID") != id1 {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
			if status == 0 {
				objUser = new(user)
				objUser.ID = fun.GetUint32ByName(o, "ID")
				if err := selectObjByID(objUser); err == nil {
					if fun.GetStrByName(o, "Username") == "" {
						fun.SetStrByName(o, "Username", objUser.Username)
					}
					if fun.GetStrByName(o, "Pw") == "" {
						fun.SetStrByName(o, "Pw", objUser.Pw)
					} else if id1 == 0 {
						// only root can use the plain text
						fun.SetStrByName(o, "Pw", fun.GetMd5Str(fun.GetStrByName(o, "Pw")))
					}
					if fun.GetUint32ByName(o, "Level") == 0 {
						fun.SetUint32ByName(o, "Level", objUser.Level)
					}
					if fun.GetUint32ByName(o, "OrgID") == 0 {
						fun.SetUint32ByName(o, "OrgID", objUser.OrgID)
					}
					if fun.GetStrByName(o, "Name") == "" {
						fun.SetStrByName(o, "Name", objUser.Name)
					}
					if fun.GetStrByName(o, "Email") == "" {
						fun.SetStrByName(o, "Email", objUser.Email)
					}
				} else {
					// user not exists
					status = 9003
					resp = jsonResp{status, msg.GetMsg(status, "user")}
				}
			}
		case "room", "org":
			// only root
			if id1 == 0 {
				if fun.GetUint32ByName(o, "ID") == 0 {
					// id wrong
					status = 9004
					resp = jsonResp{status, msg.GetMsg(status, "id")}
				}
			} else {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
			if status == 0 {
				// check if existed
				if target == "room" {
					objRoom = new(room)
					objRoom.ID = fun.GetUint32ByName(o, "ID")
					if err := selectObjByID(objRoom); err != nil {
						//  not found
						status = 9003
					}
				} else {
					// check org if existed
					objOrg = new(org)
					objOrg.ID = fun.GetUint32ByName(o, "ID")
					if err := selectObjByID(objOrg); err != nil {
						//  not found
						status = 9003
					}
				}
				if status == 9003 {
					resp = jsonResp{status, msg.GetMsg(status, target)}
				}
			}
		}
		// update db
		if status == 0 {
			if err := updateObj(o); err == nil {
				// success
				status = 0
				switch target {
				case "room":
					resp = jsonResp{status, jsonRespRoom{fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name")}}
				case "org":
					resp = jsonResp{status, jsonRespOrg{fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name")}}
				case "user":
					respUser = jsonRespUser{}
					respUser.ID = fun.GetUint32ByName(o, "ID")
					respUser.Level = fun.GetUint32ByName(o, "Level")
					respUser.Username = fun.GetStrByName(o, "Username")
					respUser.Name = fun.GetStrByName(o, "Name")
					respUser.Email = fun.GetStrByName(o, "Email")
					respUser.Org.ID = fun.GetUint32ByName(o, "OrgID")
					respUser.Org.Name = ""
					// search org_name
					objOrg = new(org)
					objOrg.ID = respUser.Org.ID
					if err := selectObjByID(objOrg); err == nil {
						respUser.Org.Name = objOrg.Name
					} else {
						logger.Warn(err.Error())
					}
					status = 0
					resp = jsonResp{status, respUser}
				}

			} else {
				// fail
				logger.Error(err.Error())
				status = 9006
				resp = jsonResp{status, err.Error()}
			}
		}

	}
	// response
	responseJSON(w, &resp)
}

func apiGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp     jsonResp
		respUser jsonRespUser
		status   int
		o        interface{}
		objOrg   *org
		target   string
	)
	target = ps.ByName("target")
	switch target {
	case "meeting":
		o = new(meeting)
	case "user", "users":
		o = new(user)
	case "room", "rooms":
		o = new(room)
	case "org", "orgs":
		o = new(org)
	case "notification":
		o = new(notification)
	default:
		status = 9004
		resp = jsonResp{status, msg.GetMsg(status, "target")}
	}
	if status == 0 {
		id1, _ := getThisUserID(r)
		parseReqToObj(r, ps, o)
		// check
		switch target {
		case "user", "users":
			if fun.GetUint32ByName(o, "ID") == 0 {
				fun.SetUint32ByName(o, "ID", id1)
			}
			if id1 > 0 && fun.GetUint32ByName(o, "ID") != id1 {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
		case "room", "org":
			if id1 == 0 {
				if fun.GetUint32ByName(o, "ID") == 0 {
					// id wrong
					status = 9004
					resp = jsonResp{status, msg.GetMsg(status, "id")}
				}
			} else {
				// no privilege
				status = 9004
				resp = jsonResp{status, msg.GetMsg(status, "privilege")}
			}
		}
		// search
		if status == 0 {
			switch target {
			case "meeting":
				status = 0
				resp = jsonResp{status, makeJSONrespMeetings(fun.GetUint32ByName(o, "MakeDate"))}
			case "rooms":
				status = 0
				resp = jsonResp{status, makeJSONrespRooms()}
			case "orgs":
				status = 0
				resp = jsonResp{status, makeJSONrespOrgs()}
			case "users":
				status = 0
				resp = jsonResp{status, makeJSONrespUsers()}
			case "notification":
				if o, err := selectNotification("select * from notification limit 1"); err == nil {
					status = 0
					resp = jsonResp{status, jsonRespNotification{fun.GetStrByName(o, "Message")}}
				} else {
					// no notification
					status = 9003
					resp = jsonResp{status, msg.GetMsg(status, target)}
				}
			default:
				if err := selectObjByID(o); err == nil {
					// success
					status = 0
					switch target {
					case "room":
						resp = jsonResp{status, jsonRespRoom{fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name")}}
					case "org":
						resp = jsonResp{status, jsonRespOrg{fun.GetUint32ByName(o, "ID"), fun.GetStrByName(o, "Name")}}
					case "user":
						respUser = jsonRespUser{}
						respUser.ID = fun.GetUint32ByName(o, "ID")
						respUser.Level = fun.GetUint32ByName(o, "Level")
						respUser.Username = fun.GetStrByName(o, "Username")
						respUser.Name = fun.GetStrByName(o, "Name")
						respUser.Email = fun.GetStrByName(o, "Email")
						respUser.Org.ID = fun.GetUint32ByName(o, "OrgID")
						respUser.Org.Name = ""
						// search org_name
						objOrg = new(org)
						objOrg.ID = respUser.Org.ID
						if err := selectObjByID(objOrg); err == nil {
							respUser.Org.Name = objOrg.Name
						} else {
							logger.Warn(err.Error())
						}
						if conf.LDAP.Enable && respUser.Org.ID == conf.LDAP.OrgID {
							respUser.Ldap = true
						} else {
							respUser.Ldap = false
						}
						status = 0
						resp = jsonResp{status, respUser}
					}
				} else {
					// not found
					logger.Error(err.Error())
					status = 9003
					resp = jsonResp{status, msg.GetMsg(status, target)}
				}

			}
		}

	}
	// response
	responseJSON(w, &resp)
}
