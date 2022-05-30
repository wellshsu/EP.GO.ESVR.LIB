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

func (this *Player) Send(id protocol.MIDEnum, data proto.Message, mreq ...*xmsg.MsgReq) {
	var nmreq *xmsg.MsgReq
	if len(mreq) == 1 {
		nmreq = mreq[0]
	} else {
		nmreq = xmsg.GetMsgReq()
		nmreq.Route.UID = proto.Int(this.ID)
		nmreq.Route.RID = proto.Int(int(id))
		nmreq.Route.Src = proto.String(xmsg.GUrl)
		nmreq.Route.Dst = proto.String(this.ConnUrl)
		nmreq.Route.MID1 = proto.Int64(this.ConnID)
	}
	xserver.SendMsg(int(id), data, nmreq)
}

func (this *Player) OnLogin(req *protocol.GM_LoginReq, mreq *xmsg.MsgReq) {
	this.Online = 1
	if this.ConnUrl != mreq.Route.GetSrc() && this.ConnID != mreq.Route.GetMID1() { // 网关
		if this.ConnUrl != "" && this.ConnID != -1 {
			this.Send(protocol.MID.GM_KICK_OFF, &protocol.GM_Common{Result: proto.Int(0)})
		}
	}
	this.ConnUrl = mreq.Route.GetSrc()
	this.ConnID = mreq.Route.GetMID1()
	mreq.Route.UID = proto.Int(this.ID)

	resp := &protocol.GM_LoginResp{}
	resp.ID = proto.Int(this.ID)
	resp.Status = proto.Int(this.Online)
	this.Send(protocol.MID.GM_LOGIN_RESPONSE, resp)
}

func (this *Player) OnLogout(req *protocol.RPC_ConnNotifyOfflineReq, rreq *xmsg.RpcReq) {
	if this.ConnUrl != req.GetUrl() && this.ConnID != req.GetCID() {
		xlog.Warn("ctx.Player.OnLogout: player-%v's oldurl-%v#%v doesn't equals to newurl-%v#%v", this.ID, this.ConnUrl, this.ConnID, req.GetUrl(), req.GetCID())
	} else {
		xlog.Info("ctx.Player.OnLogout: player-%v offline, url-%v#%v", this.ID, this.ConnUrl, this.ConnID)
		this.Online = 0
	}
}
