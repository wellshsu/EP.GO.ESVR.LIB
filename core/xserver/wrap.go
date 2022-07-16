//go:binary-only-package
//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//
package xserver

import (
	_ "math/rand"
	"time"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
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
