package ctx

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/models/mmn"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protocol"

	"github.com/golang/protobuf/proto"
)

type Player struct {
	*mmn.Player
}

func NewPlayer(md ...*mmn.Player) *Player {
	this := new(Player)
	if len(md) == 1 && md[0] != nil {
		this.Player = md[0]
	} else {
		this.Player = mmn.NewPlayer()
	}
	return this
}

func (this *Player) Send(id protocol.MIDEnum, data proto.Message, mreq *xmsg.MsgReq) {
	nmreq := xmsg.NewMsgReq()
	nmreq.Route.UID = proto.Int(this.ID)
	nmreq.Route.Src = proto.String(xmsg.GUrl)
	nmreq.Route.Dst = proto.String(this.ConnUrl)
	nmreq.Route.MID1 = mreq.Route.MID1
	xserver.SendMsg(int(id), data, nmreq)
}

func (this *Player) OnLogin(req *protocol.GM_LoginReq, mreq *xmsg.MsgReq) {
	this.Online = 1
	if this.ConnUrl != mreq.Route.GetSrc() { // 网关
		if this.ConnUrl != "" {
			this.Send(protocol.MID.GM_LOGIN_RESPONSE, &protocol.GM_Common{Result: proto.Int(0)}, mreq)
		}
	}
	this.ConnUrl = mreq.Route.GetSrc()
	mreq.Route.UID = proto.Int(this.ID)

	resp := &protocol.GM_LoginResp{}
	resp.ID = proto.Int(this.ID)
	resp.Status = proto.Int(this.Online)
	this.Send(protocol.MID.GM_LOGIN_RESPONSE, resp, mreq)
}

func (this *Player) OnLogout(req *protocol.RPC_ConnNotifyOfflineReq, rreq *xmsg.RpcReq) {
	playerUrl := req.GetUrl()
	if this.ConnUrl != playerUrl {
		xlog.Warn("ctx.Player.OnLogout: player-%v's oldurl-%v doesn't equals to newurl-%v", this.ID, playerUrl, this.ConnUrl)
	} else {
		xlog.Info("ctx.Player.OnLogout: player-%v offline, url-%v", this.ID, this.ConnUrl)
		this.Online = 0
	}
}
