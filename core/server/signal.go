//go:binary-only-package
package server

import (
	_ "esvr/core/utility/xrun"
	_ "os"
	_ "os/signal"
	_ "syscall"
)

func doWatch(rch chan<- string) {
	return
}
func WatchSignal() <-chan string {
	return nil
}
