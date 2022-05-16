//go:binary-only-package
package xconn

import (
	"errors"
	"net"
)

const (
	SENDER_CHAN_LENGTH = 512
)

var (
	ERR_SENDER_FULL = errors.New("sender chan is full")
)

type Sender struct {
	ChMsg  chan []byte
	ChSend chan bool
}

func NewSender() *Sender {
	return nil
}
func (this *Sender) Send(conn net.Conn) error {
	return nil
}
func (this *Sender) Write(data []byte) error {
	return nil
}
func (this *Sender) Clear() {
	return
}
