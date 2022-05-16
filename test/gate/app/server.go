// 线程分布（共30个线程）
package app

import (
	"esvr/core/server"
	"esvr/core/utility/xconn"
	"esvr/core/utility/xevt"
	"esvr/core/utility/xlog"
	"esvr/core/utility/xmsg"
	"esvr/core/utility/xstring"
	"esvr/test/shared/protocol"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type GateClient struct {
	xconn.Client
	UID     int
	GateUrl string
}

func NewGateClient(server *xconn.Server) *GateClient {
	this := new(GateClient)
	this.UID = -1
	this.Server = server
	this.CTOR(this)
	return this
}

type GateServer struct {
	server.Server
	Svr *xconn.Server

	IsWS    bool   // 是否面向WebSocket连接（默认true）
	MaxConn int    // 最大连接数（超过该值将不再接收新的连接）（默认5000）
	MaxLoad int    // 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
	MaxBody int    // 单连接最大的消息体（不包含头部）（默认1024 * 1024 = 1048576）
	Address string // 前端连接地址
	Timeout int    // 连接超时时间（超过该值则主动断开连接）（默认15秒）

	BeatMsg []byte
}

func NewGateServer() *GateServer {
	this := &GateServer{}
	this.CTOR(this)
	this.BeatMsg, _ = xmsg.PackMsg(int(protocol.EID.GM_HEART_BEAT), &protocol.GM_Common{Result: proto.Int(0)})
	server.RegEvt(server.EVT_SERVER_STARTED, func(param interface{}) {
		svrID := this.Config.SvrID()
		this.Svr = xconn.NewServer().
			SetAddr(this.Address).SetIsWS(this.IsWS).
			SetMaxConn(this.MaxConn).SetMaxLoad(this.MaxLoad).
			SetMaxBody(this.MaxBody).SetTimeout(this.Timeout).
			SetOnAccept(func(client xconn.IClient) {
				gclient := client.(*GateClient)
				gclient.GateUrl = fmt.Sprintf("%v#%v", svrID, gclient.GetID())
			}).
			SetOnRemove(func(client xconn.IClient) {
				gclient := client.(*GateClient)
				this.RemoveClient(gclient)
			}).
			SetOnReceive(func(client xconn.IClient, bytes []byte) {
				gclient := client.(*GateClient)
				this.FromClient(gclient, bytes)
			}).
			SetNewClient(func() interface{} {
				return NewGateClient(this.Svr)
			}).
			Start()
	})
	server.RegRpc(int(protocol.EID.RPC_GET_ONLINE_FROM_GATE), func(reply *xevt.EvtReply, frame *xmsg.Frame) {
		resp := &protocol.RPC_GetOnlineFromGateResp{}
		reply.Done(resp)
		this.Svr.Clients.Range(func(k, v interface{}) bool {
			client := v.(*GateClient)
			if client.UID != -1 {
				resp.ID = append(resp.ID, int32(client.UID))
				resp.Url = append(resp.Url, client.GateUrl)
			}
			return true
		})
	})
	return this
}

func (this *GateServer) InitConfig() bool {
	if this.Server.InitConfig() == false {
		return false
	}
	config := this.GetConfig()
	this.IsWS = config.Raw.DefaultBool("client::ws", true)
	this.MaxConn = config.Raw.DefaultInt("client::maxConn", 5000)
	this.MaxLoad = config.Raw.DefaultInt("client::maxLoad", 100)
	this.MaxBody = config.Raw.DefaultInt("client::maxBody", 1024*1024)
	this.Address = config.Raw.String("client::addr")
	this.Timeout = config.Raw.DefaultInt("client::timeout", 15)
	return true
}

func (this *GateServer) UpdateTitle() string {
	title := this.Server.UpdateTitle()
	if this.Svr != nil {
		title += fmt.Sprintf("[CON-%v]", this.Svr.OnlineNum)
	}
	return title
}

func (this *GateServer) RecvMsg(frame *xmsg.Frame) {
	this.ToClient(frame)
}

func (this *GateServer) ToClient(frame *xmsg.Frame) {
	url := frame.GetDstUrl()
	strs := xstring.Split(url, "#")
	if len(strs) != 2 {
		xlog.Warn("ctx.GateServer.ToClient: invalid dst url: %v", url)
		return
	}
	id := xstring.ToInt(strs[1])
	client := this.Svr.GetClient(id)
	if client == nil {
		xlog.Warn("ctx.GateServer.ToClient: no client-%v found, uid=%v, msgId=%v, dst=%v, src=%v",
			id, frame.GetUID(), xmsg.ParseID(frame.RawData), url, frame.GetSrcUrl())
		return
	}
	this.FromServer(client.(*GateClient), frame)
}

func (this *GateServer) ToServer(client *GateClient, dst string, bytes []byte, gol int, gor int, rw bool) bool {
	frame := &xmsg.Frame{
		SrcUrl:  proto.String(client.GateUrl),
		DstUrl:  proto.String(dst),
		UID:     proto.Int(client.UID),
		RawData: bytes,
		GoID:    proto.Int(xmsg.PackGo(gol, gor)),
		RW:      proto.Bool(rw),
	}
	return server.SendFrame(frame)
}

func (this *GateServer) RemoveClient(client *GateClient) {
	if client.UID != -1 {
		hall := server.GLan.SelectRand("hall")
		if hall != nil {
			req := &protocol.RPC_GateNotifyOfflineReq{}
			req.UID = proto.Int(client.UID)
			req.Url = proto.String(client.GateUrl)
			server.SendAsync(int(protocol.EID.RPC_GATE_NOTIFY_OFFLINE), client.UID, req, hall.ServerUrl(), nil)
		}
		client.UID = -1
	}
}

func (this *GateServer) FromClient(client *GateClient, bytes []byte) {
	id := xmsg.ParseID(bytes)
	eid := protocol.EIDEnum(id)
	if eid == protocol.EID.GM_HEART_BEAT {
		client.Write(this.BeatMsg)
	} else {
		route := server.ROUTEMAP[id]
		if route == nil {
			xlog.Warn("ctx.GateServer.FromClient: no route for id-%v found", id)
		} else {
			for _, dst := range route.Dst {
				if dst == "client" {
					continue
				}
				conn := server.GLan.SelectRand(dst)
				if conn == nil {
					xlog.Warn("ctx.GateServer.FromClient: select conn failed, uid=%v, id=%v, tag=%v", client.UID, id, dst)
				} else {
					this.ToServer(client, conn.ServerUrl(), bytes, route.GoL, route.GoR, route.RW)
				}
			}
		}
	}
}

func (this *GateServer) FromServer(client *GateClient, frame *xmsg.Frame) {
	id := xmsg.ParseID(frame.RawData)
	if id == -1 {
		xlog.Warn("ctx.GateServer.FromServer: parse msg id failed, uid=%v", client.UID)
		return
	}
	eid := protocol.EIDEnum(id)
	xlog.Info("ctx.GateServer.FromServer: client-%v, msgid-%v, size-%v", client.UID, id, len(frame.RawData))
	if eid == protocol.EID.GM_LOGIN_RESPONSE && client.UID == -1 {
		client.UID = int(frame.GetUID())
		xlog.Notice("ctx.GateServer.FromServer: client-%v was binded to uid-%v", client.ID, client.UID)
	}
	if eid == protocol.EID.GM_KICK_OFF && client.UID != -1 {
		xlog.Notice("ctx.GateServer.FromServer: client-%v was unbinded from uid-%v", client.ID, client.UID)
		client.UID = -1
	}
	client.Write(frame.RawData)
}
