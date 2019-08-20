package entity

import (
	"github.com/name5566/leaf/gate"
)

type Player struct {
	UID       int
	LoginTime int
	RoomID    string
	Name      string
	gate.Agent
}
