package msg

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
