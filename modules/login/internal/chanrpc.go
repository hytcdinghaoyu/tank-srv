package internal

import (
	"fmt"
	"tank-srv/entity"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	fmt.Println(a)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	var player = entity.OnlinePlayerMap[a]

	//如果正在房间，先离开房间
	if player.RoomID != "" {
		entity.RoomsMap[player.RoomID].LeaveRoom(player)
	}

	delete(entity.OnlinePlayerMap, a)

	log.Debug("agent disconneted, server player online num: %v", len(entity.OnlinePlayerMap))

}
