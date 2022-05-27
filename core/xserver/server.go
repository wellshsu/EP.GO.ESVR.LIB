//go:binary-only-package
package xserver

import (
	_ "fmt"
	_ "sort"
	_ "time"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xobj"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xos"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

const (
	EVT_SERVER_STARTED = -1 // 服务就绪（配置就绪 & 日志就绪 & DB就绪 & Redis就绪 & Lan就绪）
	EVT_SERVER_CHANGED = -2 // 服务变更（参数类型：[]interface{}{added map[string][]string, removed map[string][]string}）
)

type IServer interface {
	Init() bool                                     // 初始化
	Start()                                         // 服务启动
	Update(delta float32)                           // 服务循环
	Destroy()                                       // 服务结束
	PreQuit()                                       // 服务即将退出
	Name() string                                   // 服务名称
	InitConfig() bool                               // 读取配置
	GetConfig() *SvrCfg                             // 获取配置
	UpdateTitle() string                            // 更新标题
	RecvMsg(mreq *xmsg.MsgReq)                      // 接收Msg消息
	RecvRPC(rreq *xmsg.RpcReq, rresp *xmsg.RpcResp) // 接收Rpc消息
	RecvCgi(creq *xmsg.CgiReq, cresp *xmsg.CgiResp) // 接收Cgi消息
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
func (this *Server) RecvMsg(mreq *xmsg.MsgReq) {
	return
}

// 接收Rpc消息
func (this *Server) RecvRPC(rreq *xmsg.RpcReq, rresp *xmsg.RpcResp) {
	return
}

// 接收Cgi消息
func (this *Server) RecvCgi(rreq *xmsg.CgiReq, rresp *xmsg.CgiResp) {
	return
}
