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
	ID         string       `json:"id"`
	Room       jsonRespRoom `json:"room"`
	StartTime  uint32       `json:"start_time"`
	EndTime    uint32       `json:"end_time"`
	Maker      jsonRespUser `json:"maker"`
	Memo       string       `json:"memo"`
	CreateTime uint32       `json:"create_time"`
}

type jsonRespNotification struct {
	Message string `json:"message"`
}
