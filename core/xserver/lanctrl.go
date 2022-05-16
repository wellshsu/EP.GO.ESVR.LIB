//go:binary-only-package
package xserver

import (
	_ "fmt"
	_ "net/http"
	_ "net/url"
	_ "time"

	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xevt"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xfs"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
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
