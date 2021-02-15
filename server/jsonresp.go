package server

type jsonResp struct {
	Status  int         `json:"status"`
	Results interface{} `json:"results"`
}
type jsonRespOrg struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type jsonRespRoom struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type jsonRespUser struct {
	ID       uint32      `json:"id"`
	Username string      `json:"username"`
	Level    uint32      `json:"level"`
	Org      jsonRespOrg `json:"org"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Ldap     bool        `json:"ldap"`
}

type jsonRespMeeting struct {
	Room      jsonRespRoom `json:"room"`
	StartTime uint32       `json:"start_time"`
	EndTime   uint32       `json:"end_time"`
	Maker     jsonRespUser `json:"maker"`
	Memo      string       `json:"memo"`
	MakeDate  uint32       `json:"make_date"`
}

type jsonRespNotification struct {
	Message string `json:"message"`
}
