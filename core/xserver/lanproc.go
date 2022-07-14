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
	"sync"
	_ "sync/atomic"
	_ "time"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xproto"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
)

const (
	LAN_CIN_MAX_FRAME  = 50000
	LAN_COUT_MAX_FRAME = 50000
)

var (
	GLan  *LanSvr
	GProc []*Proc
)

type Proc struct {
	TID   int64              // 线路的GoID
	Num   int                // 线路线程总数
	CIN   chan xproto.IFrame // 输入队列
	COUT  chan xproto.IFrame // 输出队列
	Loop  bool
	Pause bool
	Resp  sync.Map // map[int64]chan *xproto.RpcReq/*xproto.CgiFrame
	// 自增ID
}

func NewProc() *Proc {
	return nil
}
func (this *Proc) PopCIN() (xproto.IFrame, bool) {
	return nil, false
}
func (this *Proc) PushCIN(frame xproto.IFrame) bool {
	return false
}
func (this *Proc) MaxID() int64 {
	return 0
}
func procFrame(goNum int, handleMsg func(*xproto.MsgReq),
	handleRpc func(*xproto.RpcReq, *xproto.RpcResp),
	handleCgi func(*xproto.CgiReq, *xproto.CgiResp)) {
	return
}
