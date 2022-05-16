package ctx

import (
	"esvr/core/server"
	"esvr/core/utility/xevt"
	"esvr/core/utility/xmsg"
	"esvr/core/utility/xorm"
	"esvr/core/utility/xsession"
	"esvr/test/shared/models/mmn"
	"esvr/test/shared/protocol"
)

func init() {
	server.RegMsg(int(protocol.EID.GM_LOGIN_REQUEST), HandleLogin)
	server.RegRpc(int(protocol.EID.RPC_GATE_NOTIFY_OFFLINE), HandleLogout)
}

func HandleLogin(frame *xmsg.Frame) {
	req := &protocol.GM_LoginReq{}
	if xmsg.ParseMsg(frame.RawData, req) != nil {
		return
	}
	account := req.GetAccount()
	password := req.GetPassword()
	player := ReadPlayer(xorm.Parse("account == {0} && password == {1}", account, password))
	if player != nil {
		player.OnLogin(req, frame)
	} else {
		p := mmn.NewPlayer()
		p.Account = account
		p.Password = password
		xsession.GWrite(p)
		player = NewPlayer(p)
		player.OnLogin(req, frame)
	}
}

func HandleLogout(reply *xevt.EvtReply, frame *xmsg.Frame) {
	req := &protocol.RPC_GateNotifyOfflineReq{}
	if xmsg.ParseMsg(frame.RawData, req) != nil {
		return
	}
	playerID := int(req.GetUID())
	player := ReadPlayer(playerID)
	if player != nil {
		player.OnLogout(req, frame)
	}
}

func ReadPlayer(idOrCond ...interface{}) *Player {
	player := mmn.NewPlayer()
	if len(idOrCond) == 1 {
		switch idOrCond[0].(type) {
		case int:
			player.ID = idOrCond[0].(int)
			player = xsession.GRead(player).(*mmn.Player)
			break
		case *xorm.Condition:
			player = xsession.GRead(player, idOrCond[0].(*xorm.Condition)).(*mmn.Player)
			break
		}
	}
	if player.IsValid() {
		return NewPlayer(player)
	} else {
		return nil
	}
}

func ListPlayer(rw bool, cond ...*xorm.Condition) []*Player {
	player := mmn.NewPlayer()
	player.RW(rw)
	players := *xsession.GList(player, cond...).(*[]*mmn.Player)
	fplayers := []*Player{}
	for _, temp := range players {
		fplayers = append(fplayers, NewPlayer(temp))
	}
	return fplayers
}
