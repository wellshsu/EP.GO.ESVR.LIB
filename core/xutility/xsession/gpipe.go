//go:binary-only-package
package xsession

import (
	_ "fmt"
	_ "sync/atomic"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

// 分配一个推送管道（多线程）
func allocPipe(tid int64, log int, prefix string) chan *PipeRecord {
	return nil
}
