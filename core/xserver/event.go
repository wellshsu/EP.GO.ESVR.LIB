//go:binary-only-package
package xserver

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xevt"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
)

var (
	GMsg = xevt.NewEvtMgr(true)
	GRpc = xevt.NewEvtMgr(false)
	GCgi = xevt.NewEvtMgr(false)
	GEvt = xevt.NewEvtMgr(true)
)

type MsgFunc func(*xmsg.MsgReq)

func (this MsgFunc) Handle(reply *xevt.EvtReply, param1 interface{}, param2 interface{}) {
	return
}

// 注册Msg消息（用于客户端和服务器之间交互）（全局）
func RegMsg(id int, fun func(*xmsg.MsgReq)) int {
	return 0
}

// 注销Msg消息（用于客户端和服务器之间交互）（全局）
func UnregMsg(id int, hid int) bool {
	return false
}

// 广播Msg消息（用于客户端和服务器之间交互）（全局）
func NotifyMsg(id int, mreq *xmsg.MsgReq) bool {
	return false
}

type RpcFunc func(rreq *xmsg.RpcReq, rresp *xmsg.RpcResp)

func (this RpcFunc) Handle(reply *xevt.EvtReply, param1 interface{}, param2 interface{}) {
	return
}

// 注册Rpc消息（用于服务器之间交互）（全局）
func RegRpc(id int, fun func(rreq *xmsg.RpcReq, rresp *xmsg.RpcResp)) int {
	return 0
}

// 注销Rpc消息（用于服务器之间交互）（全局）
func UnregRpc(id int, hid int) bool {
	return false
}

// 广播Rpc消息（用于服务器之间交互）（全局）
func NotifyRpc(id int, req *xmsg.RpcReq, resp *xmsg.RpcResp) bool {
	return false
}

type CgiFunc func(req *xmsg.CgiReq, resp *xmsg.CgiResp)

func (this CgiFunc) Handle(reply *xevt.EvtReply, req interface{}, resp interface{}) {
	return
}

// 注册Cgi消息（用于客户端和服务器之间交互）（全局）
func RegCgi(id int, fun func(req *xmsg.CgiReq, resp *xmsg.CgiResp)) int {
	return 0
}

// 注销Cgi消息（用于客户端和服务器之间交互）（全局）
func UnregCgi(id int, hid int) bool {
	return false
}

// 广播Cgi消息（用于客户端和服务器之间交互）（全局）
func NotifyCgi(id int, req *xmsg.CgiReq, resp *xmsg.CgiResp) bool {
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
