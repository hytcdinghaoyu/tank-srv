package gate

import (
	"tank-srv/modules/battle"
	"tank-srv/modules/login"
	"tank-srv/msg"
)

func init() {
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.JoinRoom{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.LeaveRoom{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.Battle{}, battle.ChanRPC)

}
