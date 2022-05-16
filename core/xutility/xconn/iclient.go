//go:binary-only-package
package xconn

import (
	"net"
)

const (
	LOGON_WAIT_TIME     = 10
	READ_WRITE_DEADLINE = 1e10
)

type IClient interface {
	GetServer() *Server
	SetServer(server *Server)
	GetConn() net.Conn
	SetConn(conn net.Conn)
	GetID() int
	SetID(id int)
	GetExpired() int
	SetExpired(expired int)
	GetTick() int
	SetTick(tick int)
	GetLoad() int
	SetLoad(load int)
	GetChClose() chan bool
	SetChClose(ch chan bool)
	GetSender() *Sender
	SetSender(sender *Sender)
	GetReceiver() *Receiver
	SetReceiver(receiver *Receiver)
	Reset(id int, conn net.Conn)
	Clear()
	Write(bytes []byte) error
	Kick()
	Update() bool
}
