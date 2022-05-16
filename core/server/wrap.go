//go:binary-only-package
package server

import (
	_ "esvr/core/utility/xlog"
	_ "math/rand"
	"time"
)

const (
	SERVER_SLEEP time.Duration = 10 * time.Millisecond
)

var (
	GWrap   *Wrap
	GServer IServer
)

type Wrap struct {
	Svr    IServer
	ChQuit chan bool // 阻塞chan
}

func NewWrap(s IServer) *Wrap {
	return nil
}
func (this *Wrap) Init() bool {
	return false
}
func (this *Wrap) Run() {
	return
}
func (this *Wrap) Destroy() {
	return
}
func (this *Wrap) Stop() {
	return
}
func Start(svr IServer) {
	return
}
func Stop() {
	return
}
