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
	_ "os"
	_ "os/signal"
	_ "syscall"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xrun"
)

func doWatch(rch chan<- string) {
	return
}
func WatchSignal() <-chan string {
	return nil
}
