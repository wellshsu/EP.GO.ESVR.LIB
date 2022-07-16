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
package xsession

import (
	_ "fmt"
	_ "reflect"
	_ "strings"
	_ "sync"
	_ "sync/atomic"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xfs"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xruntime"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xtime"
)

// 开始会话（返回会话id，使用该id结束会话）
// [tag-会话标签]
// [log-日志层级（参考xlog的LogLevel）]
// [rw-是否读写]
func GStart(tag interface{}, log int, rw bool) int64 {
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
