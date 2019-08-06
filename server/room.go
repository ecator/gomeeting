package server

import (
	"errors"

	"github.com/ecator/gomeeting/fun"
)

type room struct {
	ID   uint32
	Name string
}

func selectRoom(sql string) (*room, error) {
	r, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("not found " + sql)
	}
	ret := new(room)
	ret.ID = fun.Str2Uint32(r[0]["id"])
	ret.Name = r[0]["name"]
	return ret, nil
}
