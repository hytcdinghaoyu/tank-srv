package entity

import (
	"errors"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/segmentio/ksuid"
)

var (
	RoomsMap          = make(map[string]*Room)
	errRoomNotFound   = errors.New("room id not found")
	errRoomIsFull     = errors.New("room is full")
	errAllReadyInRoom = errors.New("player in room")
)

const (
	RoomMaxPlayerNum int = 2
)

type Room struct {
	RoomID  string
	State   int
	Players map[gate.Agent]Player
}

func CreateRoom(a gate.Agent) (string, error) {
	var roomID = ksuid.New().String()

	RoomsMap[roomID] = &Room{
		RoomID:  roomID,
		Players: make(map[gate.Agent]Player),
	}
	RoomsMap[roomID].Players[a] = Player{RoomID: roomID}

	log.Debug("created a new room: %s, total room num: %v", roomID, len(RoomsMap))

	return roomID, nil
}

func (r *Room) JoinRoom(a gate.Agent) (string, error) {
	//检查房间是否存在
	var _, ok = RoomsMap[r.RoomID]
	if !ok {
		return "", errRoomNotFound
	}

	//房间人数是否满
	if len(r.Players) == RoomMaxPlayerNum {
		return "", errRoomIsFull
	}

	//是否已经加入，不能重复加入
	if _, ok := r.Players[a]; ok {
		return "", errAllReadyInRoom
	}

	r.Players[a] = Player{RoomID: r.RoomID}
	//设置大厅玩家的房间号
	OnlinePlayerMap[a].RoomID = r.RoomID

	return r.RoomID, nil
}

func (r *Room) LeaveRoom(a gate.Agent) error {
	var _, ok = RoomsMap[r.RoomID]
	if !ok {
		return errRoomNotFound
	}

	delete(r.Players, a)
	OnlinePlayerMap[a].RoomID = ""

	//全部人离开后，删除房间
	if len(r.Players) == 0 {
		delete(RoomsMap, r.RoomID)
		log.Debug("room destroyed: %s", r.RoomID)
	}

	return nil
}

func (r *Room) Broadcast(msg interface{}) {
	for a := range r.Players {
		a.WriteMsg(msg)
	}
}
