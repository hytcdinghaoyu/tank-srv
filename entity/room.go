package entity

import (
	"github.com/name5566/leaf/gate"
)

var (
	RoomsMap = make(map[string]Room)
)

const (
	RoomMaxPlayerNum = 2
)

type Room struct {
	RoomID  string
	State   int
	Players map[gate.Agent]Player
}

func (r *Room) CreateRoom(a gate.Agent) {

}

func (r *Room) JoinRoom(roomID string, a gate.Agent) error {
	var _, ok = RoomsMap[roomID]
	if !ok {
	}
	return nil
}
