package entity

import (
	"github.com/name5566/leaf/gate"
)

var (
	OnlinePlayerMap = make(map[gate.Agent]Player)
)

// MaxPlayerNum defines single server max user num
const MaxPlayerNum = 1000
