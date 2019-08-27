package internal

import (
	"reflect"
	"tank-srv/constants"
	"tank-srv/entity"
	"tank-srv/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	handler(&msg.Battle{}, handleBattle)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleBattle(args []interface{}) {
	m := args[0].(*msg.Battle)
	a := args[1].(gate.Agent)

	player, ok := entity.OnlinePlayerMap[a]
	if !ok || player.RoomID == "" {
		ret := msg.BattleRsp{
			Ret: constants.UserNotJoin,
		}
		a.WriteMsg(&ret)
	}

	room, ok := entity.RoomsMap[player.RoomID]
	if !ok {
		ret := msg.BattleRsp{
			Ret: constants.BattleIdNotExist,
		}
		a.WriteMsg(&ret)
	}

	userItem := room.Players[a]
	userItem.AccX = m.MoveX * userItem.AccPower / userItem.Mass
	userItem.AccY = m.MoveY * userItem.AccPower / userItem.Mass
	userItem.Rotation = m.Rotation
	userItem.SkillID = m.SkillID

}
