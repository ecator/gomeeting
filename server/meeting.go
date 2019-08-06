package server

import (
	"errors"

	"github.com/ecator/gomeeting/fun"
)

type meeting struct {
	RoomID    uint32 `json:"room_id"`
	StartTime uint32 `json:"start_time"`
	EndTime   uint32 `json:"end_time"`
	Maker     uint32 `json:"maker"`
	Memo      string `json:"memo"`
	MakeDate  uint32 `json:"make_date"`
}

func selectMeetings(sql string) ([]meeting, error) {
	var (
		ms []meeting
		m  meeting
	)
	r, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("not found " + sql)
	}
	m = meeting{}
	for i := 0; i < len(r); i++ {
		m.RoomID = fun.Str2Uint32(r[i]["room_id"])
		m.StartTime = fun.Str2Uint32(r[i]["start_time"])
		m.EndTime = fun.Str2Uint32(r[i]["end_time"])
		m.Maker = fun.Str2Uint32(r[i]["maker"])
		m.Memo = r[i]["memo"]
		m.MakeDate = fun.Str2Uint32(r[i]["make_date"])
		ms = append(ms, m)
	}

	return ms, nil

}
