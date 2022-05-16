//go:binary-only-package
package server

import (
	_ "errors"
	_ "esvr/core/utility/xlog"
	_ "esvr/core/utility/xmath"
	_ "fmt"

	"nanomsg.org/go-mangos"
	_ "nanomsg.org/go-mangos/protocol/push"
	_ "nanomsg.org/go-mangos/transport/ipc"
	_ "nanomsg.org/go-mangos/transport/tcp"
)

type LanClt struct {
	Clear bool
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
