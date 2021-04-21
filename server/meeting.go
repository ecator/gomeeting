package server

import (
	"errors"

	"github.com/ecator/gomeeting/fun"
)

type meeting struct {
	MeetingID  string `json:"meeting_id"`
	RoomID     uint32 `json:"room_id"`
	StartTime  uint32 `json:"start_time"`
	EndTime    uint32 `json:"end_time"`
	Maker      uint32 `json:"maker"`
	Memo       string `json:"memo"`
	CreateTime uint32 `json:"create_time"`
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
		m.MeetingID = r[i]["id"]
		m.RoomID = fun.Str2Uint32(r[i]["room_id"])
		m.StartTime = fun.Str2Uint32(r[i]["start_time"])
		m.EndTime = fun.Str2Uint32(r[i]["end_time"])
		m.Maker = fun.Str2Uint32(r[i]["maker"])
		m.Memo = r[i]["memo"]
		m.CreateTime = fun.Str2Uint32(r[i]["create_time"])
		ms = append(ms, m)
	}

	return ms, nil

}
