package msg

// LoginRsp Login Response
type LoginRsp struct {
	Ret int
	UID int
}

type JoinRoomRsp struct {
	Ret      int
	UID      int
	BattleID int
}
