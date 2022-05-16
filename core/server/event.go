//go:binary-only-package
package server

import (
	"esvr/core/utility/xevt"
	"esvr/core/utility/xmsg"
)

var (
	GRpc = xevt.NewEvtMgr(false)
	GMsg = xevt.NewEvtMgr(true)
	GEvt = xevt.NewEvtMgr(true)
)

type RpcFunc func(reply *xevt.EvtReply, frame *xmsg.Frame)

func (this RpcFunc) Handle(reply *xevt.EvtReply, receiver interface{}, param interface{}) {
	return
}

// 注册Rpc消息（用于服务器之间交互）（全局）
func RegRpc(id int, fun func(reply *xevt.EvtReply, frame *xmsg.Frame)) int {
	return 0
}

// 注销Rpc消息（用于服务器之间交互）（全局）
func UnregRpc(id int, hid int) bool {
	return false
}

// 广播Rpc消息（用于服务器之间交互）（全局）
func NotifyRpc(id int, reply *xevt.EvtReply, param interface{}) bool {
	return false
}

type MsgFunc func(*xmsg.Frame)

func (this MsgFunc) Handle(reply *xevt.EvtReply, receiver interface{}, param interface{}) {
	return
}

// 注册Msg消息（用于客户端和服务器或者服务器之间交互）（全局）
func RegMsg(id int, fun func(*xmsg.Frame)) int {
	return 0
}

// 注销Msg消息（用于客户端和服务器或者服务器之间交互）（全局）
func UnregMsg(id int, hid int) bool {
	return false
}

// 广播Msg消息（用于客户端和服务器或者服务器之间交互）（全局）
func NotifyMsg(id int, frame *xmsg.Frame) bool {
	return false
}

type EvtFunc func(interface{})

func (this EvtFunc) Handle(reply *xevt.EvtReply, receiver interface{}, param interface{}) {
	return
}

// 注册Evt消息（用于服务器内部）（全局）
func RegEvt(id int, fun func(interface{})) int {
	return 0
}

// 注销Evt消息（用于服务器内部）（全局）
func UnregEvt(id int, hid int) bool {
	return false
}

// 广播Evt消息（用于服务器内部）（全局）
func NotifyEvt(id int, param interface{}) bool {
	return false
}
