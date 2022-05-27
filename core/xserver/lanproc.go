//go:binary-only-package
package xserver

import (
	"sync"
	_ "sync/atomic"
	_ "time"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
)

const (
	LAN_CIN_MAX_FRAME  = 50000
	LAN_COUT_MAX_FRAME = 50000
)

var (
	GLan  *LanSvr
	GProc []*Proc
)

type Proc struct {
	TID   int64            // 线路的GoID
	Num   int              // 线路线程总数
	CIN   chan xmsg.IFrame // 输入队列
	COUT  chan xmsg.IFrame // 输出队列
	Loop  bool
	Pause bool
	Resp  sync.Map // map[int64]chan *xmsg.RpcReq/*xmsg.CgiFrame
	// 自增ID
}

func NewProc() *Proc {
	return nil
}
func (this *Proc) PopCIN() (xmsg.IFrame, bool) {
	return nil, false
}
func (this *Proc) PushCIN(frame xmsg.IFrame) bool {
	return false
}
func (this *Proc) MaxID() int64 {
	return 0
}
func procFrame(goNum int, handleMsg func(*xmsg.MsgReq),
	handleRpc func(*xmsg.RpcReq, *xmsg.RpcResp),
	handleCgi func(*xmsg.CgiReq, *xmsg.CgiResp)) {
	return
}
