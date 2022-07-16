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
package xrun

import (
	_ "bytes"
	_ "fmt"
	_ "os"
	_ "reflect"
	_ "runtime"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xfs"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xtime"
)

const (
	UNKONWN_SOURCE = "[?]"
)

// 获取函数调用堆栈信息，stack-堆栈层级，fullpath-全路径
func Caller(stack int, fullpath bool) string {
	return ""
}
func StackTrace(stack int, err interface{}) (string, int) {
	return "", 0
}

// 计算函数执行消耗的时间，在起始处调用defer Elapse(stack)()，stack-堆栈层级（0为当前层）
func Elapse(stack int, callback ...func()) func() {
	return nil
}
func doCaught(exit bool, stack int, handler ...func(string, int)) {
	return
}

// 捕捉goroutine的异常，exit-是否结束进程，handler-回调
func Caught(exit bool, handler ...func(string, int)) {
	return
}

// goroutine的wrapper（panic不会引起crash，recover后该goroutine结束）
func Exec(fun interface{}, params ...interface{}) {
	return
}
func doRun(fun interface{}, stack int, params ...interface{}) {
	return
}

// goroutine的wrapper（panic不会引起crash，recover后重启该goroutine）
func Run(fun interface{}, params ...interface{}) {
	return
}
