package server

import (
	"errors"

	"github.com/ecator/gomeeting/fun"
)

type org struct {
	ID   uint32
	Name string
}

func selectOrg(sql string) (*org, error) {
	r, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("not found " + sql)
	}
	ret := new(org)
	ret.ID = fun.Str2Uint32(r[0]["id"])
	ret.Name = r[0]["name"]
	return ret, nil

}
