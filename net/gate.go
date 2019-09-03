package net

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

type Gate struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor
	AgentChanRPC    *chanrpc.Server

	// websocket
	WSAddr      string
	HTTPTimeout time.Duration
	CertFile    string
	KeyFile     string

	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool
}

func (gate *Gate) Run(closeSig chan bool) {
	log.Debug("gate server starting")
	var wsServer *network.WSServer
	if gate.WSAddr != "" {
		wsServer = new(network.WSServer)
		wsServer.Addr = gate.WSAddr
		wsServer.MaxConnNum = gate.MaxConnNum
		wsServer.PendingWriteNum = gate.PendingWriteNum
		wsServer.MaxMsgLen = gate.MaxMsgLen
		wsServer.HTTPTimeout = gate.HTTPTimeout
		wsServer.CertFile = gate.CertFile
		wsServer.KeyFile = gate.KeyFile
		wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	var kcpServer *KCPServer
	if gate.TCPAddr != "" {
		kcpServer = new(KCPServer)
		kcpServer.Addr = gate.TCPAddr
		kcpServer.MaxConnNum = gate.MaxConnNum
		kcpServer.PendingWriteNum = gate.PendingWriteNum
		kcpServer.LenMsgLen = gate.LenMsgLen
		kcpServer.MaxMsgLen = gate.MaxMsgLen
		kcpServer.LittleEndian = gate.LittleEndian
		kcpServer.NewAgent = func(conn *KCPConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	if wsServer != nil {
		log.Debug("starting ws server addr: %s", gate.WSAddr)
		wsServer.Start()
	}
	if kcpServer != nil {
		log.Debug("starting kcp server addr: %s", gate.TCPAddr)
		kcpServer.Start()
	}
	<-closeSig
	if wsServer != nil {
		wsServer.Close()
	}
	if kcpServer != nil {
		kcpServer.Close()
	}
}

func (gate *Gate) OnDestroy() {}

type agent struct {
	conn     network.Conn
	gate     *Gate
	userData interface{}
}

func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println(err)
			log.Debug("read message: %v", err)
			break
		}

		if a.gate.Processor != nil {
			msg, err := a.gate.Processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = a.gate.Processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *agent) OnClose() {
	if a.gate.AgentChanRPC != nil {
		err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
		if err != nil {
			log.Error("chanrpc error: %v", err)
		}
	}
}

func (a *agent) WriteMsg(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) Close() {
	a.conn.Close()
}

func (a *agent) Destroy() {
	a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}
