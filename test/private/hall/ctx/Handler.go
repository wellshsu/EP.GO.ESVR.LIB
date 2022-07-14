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
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xproto"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/model/mmn"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/proto/cpb"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/proto/mpb"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/proto/rpb"
)

func init() {
	xserver.RegMsg(int(mpb.MID.GM_LOGIN_REQUEST), HandleLogin)
	xserver.RegRpc(int(rpb.RID.RPC_CONN_NOTIFY_OFFLINE), HandleLogout)
	xserver.RegCgi(int(cpb.CID.HELLO_WORLD), HandleCgi)
}

func HandleLogin(mreq *xproto.MsgReq) {
	req := &mpb.GM_LoginReq{}
	if xproto.UnpackMsg(mreq.Data, req) != nil {
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

func HandleLogout(rreq *xproto.RpcReq, rresp *xproto.RpcResp) {
	req := &rpb.RPC_ConnNotifyOfflineReq{}
	if xproto.UnpackMsg(rreq.Data, req) != nil {
		return
	}
	uid := int(req.GetUID())
	player := ReadPlayer(uid)
	if player != nil {
		player.OnLogout(req, rreq)
	}
}

func HandleCgi(creq *xproto.CgiReq, cresp *xproto.CgiResp) {
	req := &cpb.CGI_Hello{}
	xproto.UnpackCgi(creq.Body, req)
	resp := &cpb.CGI_Hello{}
	defer func() {
		cresp.Status = proto.Int(200)
		cresp.Body, _ = xproto.PackCgi(resp)
	}()
	resp.ID = proto.Int(10086)
	resp.Desc = proto.String(req.GetDesc() + " : " + "Hi, this is cgi resp.")
	if creq != nil {
	}
}
