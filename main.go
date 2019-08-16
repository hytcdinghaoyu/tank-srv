package main

import (
	"tank-srv/conf"
	"tank-srv/modules/battle"
	"tank-srv/modules/gate"
	"tank-srv/modules/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		battle.Module,
		gate.Module,
		login.Module,
	)
}
