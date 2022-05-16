//go:binary-only-package
package xevt

import (
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xcollect"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
)

type IHandler interface {
	Handle(reply *EvtReply, receiver interface{}, param interface{})
}
type EvtMgr struct {
	HID   int
	Mutil bool
	Evts  map[int]*EvtEntity
}
type EvtEntity struct {
	ID   int
	Hnds []*HndEntity
}
type HndEntity struct {
	ID   int
	Func IHandler
}
type EvtReply struct {
	Result interface{}
	Pend   chan int
}

func NewEvtReply(pend ...chan int) *EvtReply {
	return nil
}
func (this *EvtReply) Done(ret ...interface{}) {
	return
}
func NewEvtMgr(mutli bool) *EvtMgr {
	return nil
}
func (this *EvtMgr) Clear() {
	return
}
func (this *EvtMgr) Get(id int) *EvtEntity {
	return nil
}
func (this *EvtMgr) Reg(id int, handler IHandler) int {
	return 0
}
func (this *EvtMgr) Unreg(id int, hid int) bool {
	return false
}
func (this *EvtMgr) Notify(id int, reply *EvtReply, receiver interface{}, param interface{}) bool {
	return false
}
