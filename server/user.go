package server

import (
	"errors"

	"github.com/ecator/gomeeting/fun"
)

type user struct {
	ID       uint32
	Username string
	Pw       string `json:"password"`
	Level    uint32
	OrgID    uint32 `json:"org_id"`
	Name     string
	Email    string
}

func selectUser(sql string) (*user, error) {
	r, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("not found " + sql)
	}

	ret := new(user)
	ret.ID = fun.Str2Uint32(r[0]["id"])
	ret.Username = r[0]["username"]
	ret.Level = fun.Str2Uint32(r[0]["level"])
	ret.OrgID = fun.Str2Uint32(r[0]["org_id"])
	ret.Pw = r[0]["pw"]
	ret.Name = r[0]["name"]
	ret.Email = r[0]["email"]
	return ret, nil
}
