//go:binary-only-package
package xserver

import (
	_ "errors"
	_ "fmt"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmath"
	"nanomsg.org/go-mangos"
	_ "nanomsg.org/go-mangos/protocol/push"
	_ "nanomsg.org/go-mangos/transport/ipc"
	_ "nanomsg.org/go-mangos/transport/tcp"
)

type LanClt struct {
	*LanCfg
	Sockets []mangos.Socket
}

func NewLanClt(cfg *LanCfg) *LanClt {
	return nil
}
func (this *LanClt) Send(bytes []byte, idx int) error {
	return nil
}
func (this *LanClt) Close() {
	return
}
