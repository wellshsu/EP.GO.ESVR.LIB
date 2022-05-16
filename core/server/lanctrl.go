//go:binary-only-package
package server

import (
	"esvr/core/utility/xevt"
	_ "esvr/core/utility/xfs"
	_ "esvr/core/utility/xjson"
	_ "esvr/core/utility/xlog"
	"esvr/core/utility/xmsg"
	_ "esvr/core/utility/xrun"
	_ "esvr/core/utility/xsession"
	_ "esvr/core/utility/xstring"
	_ "esvr/core/utility/xtime"
	_ "fmt"
	_ "net/http"
	_ "net/url"
	_ "time"
)

func StartLan(lanCfg *LanCfg, handleRpc func(*xevt.EvtReply, *xmsg.Frame), handleMsg func(*xmsg.Frame)) {
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
