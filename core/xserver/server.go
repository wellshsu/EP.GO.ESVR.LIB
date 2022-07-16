//go:binary-only-package
//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//
package xserver

import (
	_ "fmt"
	_ "sort"
	_ "time"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xobj"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xorm"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xos"
	"github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xproto"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xtime"
)

const (
	EVT_SERVER_STARTED = -1 // 服务就绪（配置就绪 & 日志就绪 & DB就绪 & Redis就绪 & Lan就绪）
	EVT_SERVER_CHANGED = -2 // 服务变更（参数类型：[]interface{}{added map[string][]string, removed map[string][]string}）
	EVT_SERVER_PREQUIT = -3 // 服务即将退出
)

type IServer interface {
	Init() bool                                         // 初始化
	Start()                                             // 服务启动
	Update(delta float32)                               // 服务循环
	Destroy()                                           // 服务结束
	PreQuit()                                           // 服务即将退出
	Name() string                                       // 服务名称
	InitConfig() bool                                   // 读取配置
	GetConfig() *SvrCfg                                 // 获取配置
	UpdateTitle() string                                // 更新标题
	RecvMsg(mreq *xproto.MsgReq)                        // 接收Msg消息
	RecvRpc(rreq *xproto.RpcReq, rresp *xproto.RpcResp) // 接收Rpc消息
	RecvCgi(creq *xproto.CgiReq, cresp *xproto.CgiResp) // 接收Cgi消息
}
type Server struct {
	xobj.OBJECT
	REAL   IServer
	Config *SvrCfg
	FPS    int
}

func (this *Server) CTOR(CHILD interface{}) {
	return
}

// 初始化
func (_this *Server) Init() bool {
	return false
}

// 服务启动
func (this *Server) Start() { return }
func (this *Server) Update(delta float32) {
	return
}

// 服务结束
func (this *Server) Destroy() { return }
func (_this *Server) PreQuit() {
	return
}

// 服务名称
func (this *Server) Name() string {
	return ""
}

// 读取配置
func (this *Server) InitConfig() bool {
	return false
}

// 获取配置
func (this *Server) GetConfig() *SvrCfg {
	return nil
}

// 更新标题
func (this *Server) UpdateTitle() string {
	return ""
}

// 接收Msg消息
func (this *Server) RecvMsg(mreq *xproto.MsgReq) {
	return
}

// 接收Rpc消息
func (this *Server) RecvRpc(rreq *xproto.RpcReq, rresp *xproto.RpcResp) {
	return
}

// 接收Cgi消息
func (this *Server) RecvCgi(rreq *xproto.CgiReq, rresp *xproto.CgiResp) {
	return
}
