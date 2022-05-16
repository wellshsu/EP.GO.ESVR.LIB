//go:binary-only-package
package xos

import (
	_ "esvr/core/utility/xstring"
	_ "esvr/core/utility/xtime"
	_ "syscall"
	_ "time"
	_ "unsafe"

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
