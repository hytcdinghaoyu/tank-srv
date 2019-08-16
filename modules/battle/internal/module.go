package internal

import (
	"fmt"
	"log"
	"tank-srv/base"
	"time"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	ticker   = time.NewTicker(1000 * time.Millisecond)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) Run(closeSig chan bool) {
	defer ticker.Stop()
	preTime := time.Now().UnixNano()
	for {
		select {
		case <-ticker.C:
			{
				nowTime := time.Now().UnixNano()
				diff := float32(nowTime - preTime)
				fmt.Println("diff =", diff)
			}
		case close := <-closeSig:
			if close {
				fmt.Println("Ticker Stopped!")
				return
			}

		}
	}
}

func (m *Module) OnDestroy() {
	log.Println("destroy")

}
