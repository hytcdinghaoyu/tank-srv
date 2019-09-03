package net

import (
	"net"
	"sync"

	"github.com/name5566/leaf/log"
)

type ConnSet map[net.Conn]struct{}

type KCPConn struct {
	sync.Mutex
	conn      net.Conn
	writeChan chan []byte
	closeFlag bool
	msgParser *MsgParser
}

func newKCPConn(conn net.Conn, pendingWriteNum int, msgParser *MsgParser) *KCPConn {
	kcpConn := new(KCPConn)
	kcpConn.conn = conn
	kcpConn.writeChan = make(chan []byte, pendingWriteNum)
	kcpConn.msgParser = msgParser

	go func() {
		for b := range kcpConn.writeChan {
			if b == nil {
				break
			}

			_, err := conn.Write(b)
			if err != nil {
				break
			}
		}

		conn.Close()
		kcpConn.Lock()
		kcpConn.closeFlag = true
		kcpConn.Unlock()
	}()

	return kcpConn
}

func (kcpConn *KCPConn) doDestroy() {
	kcpConn.conn.(*net.TCPConn).SetLinger(0)
	kcpConn.conn.Close()

	if !kcpConn.closeFlag {
		close(kcpConn.writeChan)
		kcpConn.closeFlag = true
	}
}

func (kcpConn *KCPConn) Destroy() {
	kcpConn.Lock()
	defer kcpConn.Unlock()

	kcpConn.doDestroy()
}

func (kcpConn *KCPConn) Close() {
	kcpConn.Lock()
	defer kcpConn.Unlock()
	if kcpConn.closeFlag {
		return
	}

	kcpConn.doWrite(nil)
	kcpConn.closeFlag = true
}

func (kcpConn *KCPConn) doWrite(b []byte) {
	if len(kcpConn.writeChan) == cap(kcpConn.writeChan) {
		log.Debug("close conn: channel full")
		kcpConn.doDestroy()
		return
	}

	kcpConn.writeChan <- b
}

// b must not be modified by the others goroutines
func (kcpConn *KCPConn) Write(b []byte) {
	kcpConn.Lock()
	defer kcpConn.Unlock()
	if kcpConn.closeFlag || b == nil {
		return
	}

	kcpConn.doWrite(b)
}

func (kcpConn *KCPConn) Read(b []byte) (int, error) {
	return kcpConn.conn.Read(b)
}

func (kcpConn *KCPConn) LocalAddr() net.Addr {
	return kcpConn.conn.LocalAddr()
}

func (kcpConn *KCPConn) RemoteAddr() net.Addr {
	return kcpConn.conn.RemoteAddr()
}

func (kcpConn *KCPConn) ReadMsg() ([]byte, error) {
	return kcpConn.msgParser.Read(kcpConn)
}

func (kcpConn *KCPConn) WriteMsg(args ...[]byte) error {
	return kcpConn.msgParser.Write(kcpConn, args...)
}
