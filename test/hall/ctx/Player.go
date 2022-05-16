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

func (this *Player) Send(id protocol.EIDEnum, data proto.Message, frame ...*xmsg.Frame) {
	if len(frame) == 1 {
		xserver.SendMsg(int(id), data, frame[0])
	} else {
		xserver.SendMsg(int(id), data, xmsg.SimpleFrame(this.GateUrl, int(this.ID)))
	}
}

func (this *Player) OnLogin(req *protocol.GM_LoginReq, frame *xmsg.Frame) {
	this.Online = 1
	this.GateUrl = frame.GetSrcUrl()
	if this.GateUrl != frame.GetSrcUrl() { // 网关
		if this.GateUrl != "" {
			xserver.SendMsg(int(protocol.EID.GM_KICK_OFF), &protocol.GM_Common{Result: proto.Int(0)}, xmsg.SimpleFrame(this.GateUrl, int(this.ID)))
		}
	}
	this.GateUrl = frame.GetSrcUrl()
	frame.UID = proto.Int(this.ID)

	resp := &protocol.GM_LoginResp{}
	resp.ID = proto.Int(this.ID)
	resp.Status = proto.Int(this.Online)
	xserver.SendMsg(int(protocol.EID.GM_LOGIN_RESPONSE), resp, xmsg.ShiftFrame(frame))
}

func (this *Player) OnLogout(req *protocol.RPC_GateNotifyOfflineReq, frame *xmsg.Frame) {
	playerUrl := req.GetUrl()
	if this.GateUrl != playerUrl {
		xlog.Warn("ctx.Player.OnLogout: player-%v's oldurl-%v doesn't equals to newurl-%v", this.ID, playerUrl, this.GateUrl)
	} else {
		xlog.Info("ctx.Player.OnLogout: player-%v offline, url-%v", this.ID, this.GateUrl)
		this.Online = 0
	}
}
