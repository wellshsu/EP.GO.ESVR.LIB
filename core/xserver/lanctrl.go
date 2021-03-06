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

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xfs"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xproto"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xstring"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xtime"
)

func StartLan(lanCfg *LanCfg, handleMsg func(*xproto.MsgReq),
	handleRpc func(*xproto.RpcReq, *xproto.RpcResp),
	handleCgi func(*xproto.CgiReq, *xproto.CgiResp)) {
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
