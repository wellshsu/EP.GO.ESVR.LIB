//go:binary-only-package
package xos

import (
	_ "syscall"
	_ "time"
	_ "unsafe"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
	_ "github.com/lxn/win"
	_ "golang.org/x/sys/windows"
)

func init() {
	return
}
func setConsoleTitleWFunc(name *uint16) {
	return
}
func SetCmdTitle(title string) {
	return
}
func DisableClose() {
	return
}