//go:binary-only-package
package server

import (
	"esvr/core/utility/xevt"
	_ "esvr/core/utility/xlog"
	"esvr/core/utility/xmsg"
	_ "esvr/core/utility/xrun"
	"sync"
	_ "sync/atomic"
	_ "time"

	_ "github.com/golang/protobuf/proto"
)

const (
	CHANMSG_CACHENUM = 50000 // max cache frame count
)

var (
	GLan  *LanSvr
	GProc []*Proc
)

type Proc struct {
	TID   int64
	Num   int
	Msg   chan *xmsg.Frame
	Send  chan *xmsg.Frame
	Loop  bool
	Pause bool
	Resp  sync.Map // map[int64]chan *xmsg.Frame
	// 自增ID
}

func NewProc() *Proc {
	return nil
}
func (this *Proc) Pop() (*xmsg.Frame, bool) {
	return nil, false
}
func (this *Proc) Push(frame *xmsg.Frame) bool {
	return false
}
func (this *Proc) MaxID() int64 {
	return 0
}
func procMsg(goNum int, handleRpc func(*xevt.EvtReply, *xmsg.Frame), handleMsg func(*xmsg.Frame)) {
	return
}
