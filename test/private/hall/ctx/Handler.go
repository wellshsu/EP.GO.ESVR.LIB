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
	xserver.RegCgi(int(protocol.CID.LOGIN_REQUEST), HandleCgi)
}

func HandleLogin(mreq *xmsg.MsgReq) {
	req := &protocol.GM_LoginReq{}
	if xmsg.ParseMsg(mreq.Data, req) != nil {
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
	if xmsg.ParseMsg(rreq.Data, req) != nil {
		return
	}
	playerID := int(req.GetUID())
	player := ReadPlayer(playerID)
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
