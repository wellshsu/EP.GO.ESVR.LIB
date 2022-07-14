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
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/models/mmn"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protos/cpb"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protos/mpb"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protos/rpb"
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
	req := &mpb.GM_Common{}
	xproto.UnpackCgi(creq.Body, req)
	resp := &mpb.GM_Common{}
	defer func() {
		cresp.Status = proto.Int(200)
		cresp.Body, _ = xproto.PackCgi(resp)
	}()
	resp.Result = proto.Int(10086)
	resp.Params = append(resp.Params, "Hi, this is cgi resp.")
	resp.Params = append(resp.Params, req.GetParams()...)
	if creq != nil {
	}
}
