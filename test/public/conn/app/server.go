package app

import (
	"fmt"

	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xconn"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protocol"

	"github.com/golang/protobuf/proto"
)

type ConnClient struct {
	xconn.Client
	UID int
	Url string
}

func NewConnClient(server *xconn.Server) *ConnClient {
	this := new(ConnClient)
	this.UID = -1
	this.Server = server
	this.CTOR(this)
	return this
}

type ConnServer struct {
	xserver.Server
	Svr *xconn.Server

	IsWS    bool   // 是否面向WebSocket连接（默认true）
	MaxConn int    // 最大连接数（超过该值将不再接收新的连接）（默认5000）
	MaxLoad int    // 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
	MaxBody int    // 单连接最大的消息体（不包含头部）（默认1024 * 1024 = 1048576）
	Address string // 前端连接地址
	Timeout int    // 连接超时时间（超过该值则主动断开连接）（默认15秒）

	BeatMsg []byte
}

func NewConnServer() *ConnServer {
	this := &ConnServer{}
	this.CTOR(this)
	this.BeatMsg, _ = xmsg.PackMsg(int(protocol.MID.GM_HEART_BEAT), &protocol.GM_Common{Result: proto.Int(0)})
	xserver.RegEvt(xserver.EVT_SERVER_STARTED, func(param interface{}) {
		svrID := this.Config.SvrID()
		this.Svr = xconn.NewServer().
			SetAddr(this.Address).SetIsWS(this.IsWS).
			SetMaxConn(this.MaxConn).SetMaxLoad(this.MaxLoad).
			SetMaxBody(this.MaxBody).SetTimeout(this.Timeout).
			SetOnAccept(func(client xconn.IClient) {
				gclient := client.(*ConnClient)
				gclient.Url = svrID
			}).
			SetOnRemove(func(client xconn.IClient) {
				gclient := client.(*ConnClient)
				this.RemoveClient(gclient)
			}).
			SetOnReceive(func(client xconn.IClient, bytes []byte) {
				gclient := client.(*ConnClient)
				this.FromClient(gclient, bytes)
			}).
			SetNewClient(func() interface{} {
				return NewConnClient(this.Svr)
			}).
			Start()
	})
	xserver.RegRpc(int(protocol.RID.RPC_GET_ONLINE_FROM_CONN), func(rreq *xmsg.RpcReq, rresp *xmsg.RpcResp) {
		resp := &protocol.RPC_GetOnlineFromConnResp{}
		defer func() {
			b, _ := proto.Marshal(resp)
			rresp.Data = b
		}()
		this.Svr.Clients.Range(func(k, v interface{}) bool {
			client := v.(*ConnClient)
			if client.UID != -1 {
				resp.ID = append(resp.ID, int32(client.UID))
				resp.Url = append(resp.Url, client.Url)
				resp.CID = append(resp.CID, int64(client.ID))
			}
			return true
		})
	})
	return this
}

func (this *ConnServer) InitConfig() bool {
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

func (this *ConnServer) UpdateTitle() string {
	title := this.Server.UpdateTitle()
	if this.Svr != nil {
		title += fmt.Sprintf("[CON-%v]", this.Svr.OnlineNum)
	}
	return title
}

func (this *ConnServer) RecvMsg(frame *xmsg.MsgReq) {
	this.ToClient(frame)
}

func (this *ConnServer) ToClient(frame *xmsg.MsgReq) {
	url := frame.Route.GetDst()
	id := int(frame.Route.GetMID1())
	client := this.Svr.GetClient(id)
	if client == nil {
		xlog.Warn("ctx.ConnServer.ToClient: no client-%v found, uid=%v, msgId=%v, dst=%v, src=%v",
			id, frame.Route.GetUID(), xmsg.UnpackID(frame.Data), url, frame.Route.GetSrc())
		return
	}
	this.FromServer(client.(*ConnClient), frame)
}

func (this *ConnServer) ToServer(client *ConnClient, rid int, dst string, bytes []byte, gol int, gor int) bool {
	frame := xmsg.GetMsgReq()
	frame.Route.Src = proto.String(client.Url)
	frame.Route.Dst = proto.String(dst)
	frame.Route.RID = proto.Int(rid)
	frame.Route.UID = proto.Int(client.UID)
	frame.Route.GID = proto.Int(xmsg.PackGID(gol, gor))
	frame.Route.MID1 = proto.Int64(int64(client.ID))
	frame.Data = bytes
	return xserver.SendFrame(frame)
}

func (this *ConnServer) RemoveClient(client *ConnClient) {
	if client.UID != -1 {
		hall := xserver.GLan.SelectRand("hall")
		if hall != nil {
			req := &protocol.RPC_ConnNotifyOfflineReq{}
			req.UID = proto.Int(client.UID)
			req.Url = proto.String(client.Url)
			req.CID = proto.Int64(int64(client.ID))
			xserver.SendAsync(int(protocol.RID.RPC_CONN_NOTIFY_OFFLINE), client.UID, req, hall.ServerID(), nil)
		}
		client.UID = -1
	}
}

func (this *ConnServer) FromClient(client *ConnClient, bytes []byte) {
	rid := xmsg.UnpackID(bytes)
	eid := protocol.MIDEnum(rid)
	if eid == protocol.MID.GM_HEART_BEAT {
		client.Write(this.BeatMsg)
	} else {
		route := xserver.MSGROUTEMAP[rid]
		if route == nil {
			xlog.Warn("ctx.ConnServer.FromClient: no route for id-%v found", rid)
		} else {
			for _, dst := range route.Dst {
				if dst == "client" {
					continue
				}
				conn := xserver.GLan.SelectRand(dst)
				if conn == nil {
					xlog.Warn("ctx.ConnServer.FromClient: select conn failed, uid=%v, id=%v, tag=%v", client.UID, rid, dst)
				} else {
					this.ToServer(client, rid, conn.ServerID(), bytes, route.GoL, route.GoR)
				}
			}
		}
	}
}

func (this *ConnServer) FromServer(client *ConnClient, frame *xmsg.MsgReq) {
	id := xmsg.UnpackID(frame.Data)
	if id == -1 {
		xlog.Warn("ctx.ConnServer.FromServer: parse msg id failed, uid=%v", client.UID)
		return
	}
	eid := protocol.MIDEnum(id)
	xlog.Info("ctx.ConnServer.FromServer: client-%v, msgid-%v, size-%v", client.UID, id, len(frame.Data))
	if eid == protocol.MID.GM_LOGIN_RESPONSE && client.UID == -1 {
		client.UID = int(frame.Route.GetUID())
		xlog.Notice("ctx.ConnServer.FromServer: client-%v was binded to uid-%v", client.ID, client.UID)
	}
	if eid == protocol.MID.GM_KICK_OFF && client.UID != -1 {
		xlog.Notice("ctx.ConnServer.FromServer: client-%v was unbinded from uid-%v", client.ID, client.UID)
		client.UID = -1
	}
	client.Write(frame.Data)
}
