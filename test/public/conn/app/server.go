//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//

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
	Key     string // Https密钥
	Cert    string // Https证书
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
		this.Svr = xconn.NewServer().
			SetAddr(this.Address).SetIsWS(this.IsWS, this.Key, this.Cert).
			SetMaxConn(this.MaxConn).SetMaxLoad(this.MaxLoad).
			SetMaxBody(this.MaxBody).SetTimeout(this.Timeout).
			SetOnAccept(func(client xconn.IClient) {
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
				resp.Url = append(resp.Url, xmsg.GUrl)
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
	this.Key = config.Raw.DefaultString("client::key", "")
	this.Cert = config.Raw.DefaultString("client::cert", "")
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

func (this *ConnServer) RecvMsg(mreq *xmsg.MsgReq) {
	this.ToClient(mreq)
}

func (this *ConnServer) ToClient(mreq *xmsg.MsgReq) {
	url := mreq.Route.GetDst()
	id := int(mreq.Route.GetMID1())
	client := this.Svr.GetClient(id)
	if client == nil {
		xlog.Warn("ctx.ConnServer.ToClient: no client-%v found, uid=%v, msgId=%v, dst=%v, src=%v",
			id, mreq.Route.GetUID(), xmsg.UnpackID(mreq.Data), url, mreq.Route.GetSrc())
		return
	}
	this.FromServer(client.(*ConnClient), mreq)
}

func (this *ConnServer) ToServer(client *ConnClient, rid int, dst string, bytes []byte, gol int, gor int) bool {
	mreq := xmsg.GetMsgReq()
	mreq.Route.Src = xmsg.PGUrl
	mreq.Route.Dst = proto.String(dst)
	mreq.Route.RID = proto.Int(rid)
	mreq.Route.UID = proto.Int(client.UID)
	mreq.Route.GID = proto.Int(xmsg.PackGID(gol, gor))
	mreq.Route.MID1 = proto.Int64(int64(client.ID))
	mreq.Data = bytes
	return xserver.SendFrame(mreq)
}

func (this *ConnServer) RemoveClient(client *ConnClient) {
	if client.UID != -1 {
		hall := xserver.GLan.SelectRand("hall")
		if hall != nil {
			req := &protocol.RPC_ConnNotifyOfflineReq{}
			req.UID = proto.Int(client.UID)
			req.Url = xmsg.PGUrl
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

func (this *ConnServer) FromServer(client *ConnClient, mreq *xmsg.MsgReq) {
	id := xmsg.UnpackID(mreq.Data)
	if id == -1 {
		xlog.Warn("ctx.ConnServer.FromServer: parse msg id failed, uid=%v", client.UID)
		return
	}
	eid := protocol.MIDEnum(id)
	xlog.Info("ctx.ConnServer.FromServer: client-%v, msgid-%v, size-%v", client.UID, id, len(mreq.Data))
	if eid == protocol.MID.GM_LOGIN_RESPONSE && client.UID == -1 {
		client.UID = int(mreq.Route.GetUID())
		xlog.Notice("ctx.ConnServer.FromServer: client-%v was binded to uid-%v", client.ID, client.UID)
	}
	if eid == protocol.MID.GM_KICK_OFF && client.UID != -1 {
		xlog.Notice("ctx.ConnServer.FromServer: client-%v was unbinded from uid-%v", client.ID, client.UID)
		client.UID = -1
	}
	client.Write(mreq.Data)
}
