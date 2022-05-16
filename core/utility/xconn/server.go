//go:binary-only-package
package xconn

import (
	_ "esvr/core/utility/xlog"
	_ "esvr/core/utility/xrun"
	"net"
	_ "net/http"
	"sync"
	_ "sync/atomic"
	_ "time"

	_ "golang.org/x/net/websocket"
)

const (
	MAXID = 999999999
)

type Server struct {
	Addr      string // 前端连接地址
	IsWS      bool   // 是否面向WebSocket连接（默认true）
	MaxConn   int    // 最大连接数（超过该值将不再接收新的连接）（默认5000）
	MaxLoad   int    // 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
	MaxBody   int    // 单连接最大的消息体（不包含头部）（默认1024 * 1024 = 1048576）
	Timeout   int    // 连接超时时间（超过该值则主动断开连接）（默认15秒）
	ChClose   chan bool
	Clients   sync.Map // map[int]IClient
	Pool      sync.Pool
	OnlineNum int64
	NewClient func() interface{}    // 连接对象构造
	OnAccept  func(IClient)         // 接收到新连接
	OnRemove  func(IClient)         // 移除一个连接
	OnReceive func(IClient, []byte) // 接收连接数据
}

func NewServer() *Server {
	return nil
}

// 前端连接地址
func (this *Server) SetAddr(addr string) *Server {
	return nil
}

// 是否面向WebSocket连接（默认true）
func (this *Server) SetIsWS(ws bool) *Server {
	return nil
}

// 最大连接数（超过该值将不再接收新的连接）（默认5000）
func (this *Server) SetMaxConn(maxConn int) *Server {
	return nil
}

// 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
func (this *Server) SetMaxLoad(maxLoad int) *Server {
	return nil
}

// 单连接最大的消息体（不包含头部）（默认1024 * 1024 = 1048576）
func (this *Server) SetMaxBody(maxBody int) *Server {
	return nil
}

// 连接超时时间（超过该值则主动断开连接）（默认15秒）
func (this *Server) SetTimeout(timeout int) *Server {
	return nil
}

// 连接对象构造
func (this *Server) SetNewClient(newClient func() interface{}) *Server {
	return nil
}

// 接收到新连接
func (this *Server) SetOnAccept(onAccept func(IClient)) *Server {
	return nil
}

// 移除一个连接
func (this *Server) SetOnRemove(onRemove func(IClient)) *Server {
	return nil
}

// 接收连接数据
func (this *Server) SetOnReceive(onReceive func(IClient, []byte)) *Server {
	return nil
}

// 启动服务
func (this *Server) Start() *Server {
	return nil
}

// 停止服务
func (this *Server) Stop() {
	return
}
func (this *Server) MaxID() int {
	return 0
}
func (this *Server) AddClient(conn net.Conn) {
	return
}
func (this *Server) DelClient(client IClient) {
	return
}
func (this *Server) GetClient(id int) IClient {
	return nil
}
