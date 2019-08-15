package main

import (
	"fmt"
	"tank-srv/conf"
	"tank-srv/modules/battle"
	"tank-srv/modules/gate"
	"tank-srv/modules/login"
	"time"

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

	preTime := time.Now().UnixNano()
	t := time.NewTicker(33 * time.Millisecond)
	for {
		select {
		case <-t.C:
			{
				nowTime := time.Now().UnixNano()
				diff := float32(nowTime - preTime)
				fmt.Println("diff =", diff)
			}
		}
	}
}
