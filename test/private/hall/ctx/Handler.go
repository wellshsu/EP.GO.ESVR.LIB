//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//

package ctx

import (
	"github.com/golang/protobuf/proto"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/models/mmn"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protocol"
)

func init() {
	xserver.RegMsg(int(protocol.MID.GM_LOGIN_REQUEST), HandleLogin)
	xserver.RegRpc(int(protocol.RID.RPC_CONN_NOTIFY_OFFLINE), HandleLogout)
	xserver.RegCgi(int(protocol.CID.HELLO_WORLD), HandleCgi)
}

func HandleLogin(mreq *xmsg.MsgReq) {
	req := &protocol.GM_LoginReq{}
	if xmsg.UnpackMsg(mreq.Data, req) != nil {
		return
	}
	account := req.GetAccount()
	password := req.GetPassword()
	player := ReadPlayer(xorm.Parse("account == {0} && password == {1}", account, password))
	if player != nil {
		player.OnLogin(req, mreq)
	} else {
		p := mmn.NewPlayer()
		p.Account = account
		p.Password = password
		xsession.GWrite(p)
		player = NewPlayer(p)
		player.OnLogin(req, mreq)
	}
}

func HandleLogout(rreq *xmsg.RpcReq, rresp *xmsg.RpcResp) {
	req := &protocol.RPC_ConnNotifyOfflineReq{}
	if xmsg.UnpackMsg(rreq.Data, req) != nil {
		return
	}
	uid := int(req.GetUID())
	player := ReadPlayer(uid)
	if player != nil {
		player.OnLogout(req, rreq)
	}
}

func HandleCgi(creq *xmsg.CgiReq, cresp *xmsg.CgiResp) {
	defer func() {
		cresp.Status = proto.Int(200)
		cresp.Body = xstring.StrToBytes("hello world!")
	}()
	if creq != nil {
	}
}
