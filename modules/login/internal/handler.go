package internal

import (
	"reflect"
	"tank-srv/msg"

	"tank-srv/constants"
	"tank-srv/entity"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Login{}, handleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {

	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)

	if len(entity.OnlinePlayerMap) > entity.MaxPlayerNum {
		a.WriteMsg(&msg.LoginRsp{
			UID: m.UID,
			Ret: constants.ServerIsFull,
		})
	}

	//加入游戏大厅
	entity.OnlinePlayerMap[a] = entity.Player{UID: m.UID}
	log.Debug("server player online num: %v", len(entity.OnlinePlayerMap))

	a.WriteMsg(&msg.LoginRsp{
		UID: m.UID,
	})
}

func handleJoinRoom(args []interface{}) {
	m := args[0].(*msg.JoinRoom)
	a := args[1].(gate.Agent)

	_, ok := entity.OnlinePlayerMap[a]
	if !ok {
		a.WriteMsg(&msg.JoinRoomRsp{
			UID: m.UID,
			Ret: constants.UserNotLogin,
		})
	}

	// var found = false
	// var roomID = ""

	for roomID, room := range entity.RoomsMap {
		if len(room.Players) < entity.RoomMaxPlayerNum {
			room.JoinRoom(roomID, a)
		}
	}

}
