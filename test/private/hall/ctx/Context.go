package ctx

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/models/mmn"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protocol"
)

func init() {
	xserver.RegEvt(xserver.EVT_SERVER_STARTED, func(param interface{}) {
		xsession.GList(mmn.NewPlayer())

		players := ListPlayer(true)
		for _, player := range players {
			player.Online = 0
			player.ConnUrl = ""
		}
		conns := xserver.GLan.SelectAll("conn")
		if conns != nil {
			for _, conn := range conns {
				resp := &protocol.RPC_GetOnlineFromConnResp{}
				xserver.SendSync(int(protocol.RID.RPC_GET_ONLINE_FROM_CONN), 0, nil, resp, conn.ServerID())
				for idx, id := range resp.ID {
					playerID := int(id)
					playerUrl := resp.Url[idx]
					player := ReadPlayer(playerID)
					if player != nil {
						player.Online = 1
						player.ConnUrl = playerUrl
					}
				}
			}
		}
	})
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
