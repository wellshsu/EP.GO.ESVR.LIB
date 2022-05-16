//go:binary-only-package
package xsession

import (
	_ "esvr/core/utility/xfs"
	_ "esvr/core/utility/xlog"
	_ "esvr/core/utility/xrun"
	_ "esvr/core/utility/xruntime"
	_ "esvr/core/utility/xtime"
	_ "fmt"
	_ "reflect"
	_ "strings"
	_ "sync"
	_ "sync/atomic"
)

// 开始会话（返回会话id，使用该id结束会话）
// [tag-会话标签]
// [log-完整日志]
// [rw-是否读写]
func GStart(tag interface{}, log bool, rw bool) int64 {
	return 0
}

// 完成会话（写入/更新/删除内存，性能分析等）
func GFinish(id int64) {
	return
}

// Dump内存
// [reset-是否清除]
func GDump(reset bool) string {
	return ""
}
