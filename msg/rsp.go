package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&LoginRsp{})
}

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
