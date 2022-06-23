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
