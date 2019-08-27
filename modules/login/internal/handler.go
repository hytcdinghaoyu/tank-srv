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
	handler(&msg.Login{}, handleLogin)
	handler(&msg.JoinRoom{}, handleJoinRoom)
	handler(&msg.LeaveRoom{}, handleJoinRoom)
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
	entity.OnlinePlayerMap[a] = &entity.Player{UID: m.UID, Name: m.Name}

	log.Debug("server player online num: %v", len(entity.OnlinePlayerMap))

	a.WriteMsg(&msg.LoginRsp{
		UID: m.UID,
	})
}

func handleJoinRoom(args []interface{}) {
	// m := args[0].(*msg.JoinRoom)
	a := args[1].(gate.Agent)

	var p *entity.Player
	var ret msg.JoinRoomRsp
	p, ok := entity.OnlinePlayerMap[a]
	if !ok {
		ret = msg.JoinRoomRsp{
			Name: p.Name,
			Ret:  constants.UserNotLogin,
		}
		a.WriteMsg(&ret)
	}

	var found = false
	var roomID = ""

	for _, room := range entity.RoomsMap {
		if len(room.Players) < entity.RoomMaxPlayerNum {
			roomID, err := room.JoinRoom(p)
			if err == nil {
				ret = msg.JoinRoomRsp{Ret: 0, RoomID: roomID, Name: p.Name}
				room.Broadcast(&ret)
			} else {
				ret = msg.JoinRoomRsp{Ret: -1, Msg: err.Error()}
				a.WriteMsg(&ret)
			}

			found = true
			break
		}
	}

	if !found {
		roomID, _ = entity.CreateRoom(p)
		ret = msg.JoinRoomRsp{Ret: 0, RoomID: roomID, Name: p.Name}
		a.WriteMsg(&ret)
	}

}

func handleLeaveRoom(args []interface{}) {
	m := args[0].(*msg.LeaveRoom)
	a := args[1].(gate.Agent)

	var p *entity.Player
	var ret msg.LeaveRoomRsp
	p, ok := entity.OnlinePlayerMap[a]
	if !ok {
		ret = msg.LeaveRoomRsp{
			Ret: constants.UserNotLogin,
		}
		a.WriteMsg(&ret)
	}

	roomID := m.RoomID
	room, ok := entity.RoomsMap[roomID]
	if !ok {
		a.WriteMsg(&msg.LeaveRoomRsp{Ret: -1, Msg: "room not found"})
	}

	if err := room.LeaveRoom(p); err != nil {
		a.WriteMsg(&msg.LeaveRoomRsp{Ret: -1, Msg: err.Error()})
	}

	a.WriteMsg(&msg.LeaveRoomRsp{Ret: 1})

}
