package ctx

import (
	"esvr/core/server"
	"esvr/core/utility/xlog"
	"esvr/core/utility/xmsg"
	"esvr/test/shared/models/mmn"
	"esvr/test/shared/protocol"

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
		server.SendMsg(int(id), data, frame[0])
	} else {
		server.SendMsg(int(id), data, xmsg.SimpleFrame(this.GateUrl, int(this.ID)))
	}
}

func (this *Player) OnLogin(req *protocol.GM_LoginReq, frame *xmsg.Frame) {
	this.Online = 1
	this.GateUrl = frame.GetSrcUrl()
	if this.GateUrl != frame.GetSrcUrl() { // 网关
		if this.GateUrl != "" {
			server.SendMsg(int(protocol.EID.GM_KICK_OFF), &protocol.GM_Common{Result: proto.Int(0)}, xmsg.SimpleFrame(this.GateUrl, int(this.ID)))
		}
	}
	this.GateUrl = frame.GetSrcUrl()
	frame.UID = proto.Int(this.ID)

	resp := &protocol.GM_LoginResp{}
	resp.ID = proto.Int(this.ID)
	resp.Status = proto.Int(this.Online)
	server.SendMsg(int(protocol.EID.GM_LOGIN_RESPONSE), resp, xmsg.ShiftFrame(frame))
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
