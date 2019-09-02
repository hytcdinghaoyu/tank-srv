package internal

import (
	"tank-srv/conf"
	"tank-srv/modules/login"
	"tank-srv/msg"

	"tank-srv/net"
)

// var ticker = time.NewTicker(1000 * time.Millisecond)

type Module struct {
	*net.Gate
}

func (m *Module) OnInit() {
	m.Gate = &net.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    login.ChanRPC,
	}
}

// func (m *Module) Run(closeSig chan bool) {
// 	defer ticker.Stop()
// 	preTime := time.Now().UnixNano()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			{
// 				nowTime := time.Now().UnixNano()
// 				diff := float32(nowTime - preTime)
// 				fmt.Println("diff =", diff)
// 			}
// 		case close := <-closeSig:
// 			if close {
// 				fmt.Println("Ticker Stopped!")
// 				return
// 			}

// 		}
// 	}
// }
