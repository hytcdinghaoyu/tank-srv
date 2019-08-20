package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Login{})
	Processor.Register(&LoginRsp{})

	Processor.Register(&JoinRoom{})
	Processor.Register(&JoinRoomRsp{})
}

//Login message
type Login struct {
	Name    string
	UID     int
	InOrOut uint8
}

//JoinRoom message
type JoinRoom struct {
	UID int
}

type LeaveRoom struct {
	RoomID string
	UID    int
}
