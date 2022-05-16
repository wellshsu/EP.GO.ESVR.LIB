//go:binary-only-package
package xconn

import (
	"errors"
	_ "io"
	"net"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmath"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
)

const (
	RECEIVER_CHAN_LENGTH = 512
)

var (
	ERR_RECEIVER_CHAN_FULL          = errors.New("receiver chan is full")
	ERR_RECEIVER_INVALID_HEAD       = errors.New("receiver head is invalid")
	ERR_RECEIVER_LARGE_BODY_SIZE    = errors.New("receiver body is too large")
	ERR_RECEIVER_NEGATIVE_BODY_SIZE = errors.New("receiver body is negative")
)

type Receiver struct {
	ChMsg   chan []byte
	MaxBody int
}

func NewReceiver(maxBody int) *Receiver {
	return nil
}
func (this *Receiver) Recv(c net.Conn) (int64, error) {
	return 0, nil
}
func (this *Receiver) Clear() {
	return
}
