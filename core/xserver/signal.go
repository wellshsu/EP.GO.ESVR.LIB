//go:binary-only-package
package xserver

import (
	_ "os"
	_ "os/signal"
	_ "syscall"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
)

func doWatch(rch chan<- string) {
	return
}
func WatchSignal() <-chan string {
	return nil
}
