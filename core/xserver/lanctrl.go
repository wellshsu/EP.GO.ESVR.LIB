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
	_ "fmt"
	_ "net/http"
	_ "net/url"
	_ "time"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xfs"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

func StartLan(lanCfg *LanCfg, handleMsg func(*xmsg.MsgReq),
	handleRpc func(*xmsg.RpcReq, *xmsg.RpcResp),
	handleCgi func(*xmsg.CgiReq, *xmsg.CgiResp)) {
	return
}
func RecvLan() {
	return
}
func PauseLan() {
	return
}
func ResumeLan() {
	return
}
func CloseLan() {
	return
}
func MonitorLan() {
	return
}
func BackupLan() int {
	return 0
}
func RestoreLan() {
	return
}
