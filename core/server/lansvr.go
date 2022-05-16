//go:binary-only-package
package server

import (
	_ "errors"
	_ "esvr/core/utility/xlog"
	_ "fmt"
	_ "math/rand"
	"sync"

	"nanomsg.org/go-mangos"
	_ "nanomsg.org/go-mangos/protocol/pull"
	_ "nanomsg.org/go-mangos/transport/ipc"
	_ "nanomsg.org/go-mangos/transport/tcp"
)

type LanSvr struct {
	*LanCfg
	mangos.Socket
	Clients  sync.Map // map[string][]*LanClt
	ClientID sync.Map // map[string]*LanClt
	SClosed  bool
}

func NewLanSvr(cfg *LanCfg) *LanSvr {
	return nil
}
func (this *LanSvr) Recv() ([]byte, error) {
	return []byte{}, nil
}
func (this *LanSvr) Close() {
	return
}
func (this *LanSvr) Update(smap map[string][]string) {
	return
}
func (this *LanSvr) SelectAll(svr string) []*LanClt {
	return []*LanClt{}
}
func (this *LanSvr) SelectRand(svr string) *LanClt {
	return nil
}
func (this *LanSvr) SendMsg(svr string, bytes []byte, idx int) error {
	return nil
}
