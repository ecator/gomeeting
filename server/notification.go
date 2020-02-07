package server

import (
	"errors"
)

type notification struct {
	Message string
}

func selectNotification(sql string) (*notification, error) {
	r, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("not found " + sql)
	}
	ret := new(notification)
	ret.Message = r[0]["message"]
	return ret, nil
}
