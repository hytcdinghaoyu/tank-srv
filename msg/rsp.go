package msg

import "tank-srv/entity"

// LoginRsp Login Response
type LoginRsp struct {
	Ret int
	UID int
}

type JoinRoomRsp struct {
	Ret    int
	Name   string
	RoomID string
	Msg    string
}

type LeaveRoomRsp struct {
	Ret int
	Msg string
}

type BattleRsp struct {
	Ret     int
	RunTime int
	Users   []entity.BattleUser
	Balls   []entity.Ball
}
