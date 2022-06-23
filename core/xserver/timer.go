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
	"sync"
	_ "sync/atomic"
	"time"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xruntime"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

const (
	UPDATE_SLEEP time.Duration = 10 * time.Millisecond
)

var (
	MainTID int64 = -1
)

type TimerRecord struct {
	Timers   sync.Map
	TimerID  int64
	LastTime int
}

func (this *TimerRecord) MaxID() int64 {
	return 0
}

type TimerEntity struct {
	ID      int
	Func    func()
	Time    int
	RawTime int
	Repeat  bool
	Crash   bool
	RW      bool
	Tag     interface{}
	Log     int
}

// 设置会话的可读性（默认为可读可写）
func (this *TimerEntity) SetRW(sig bool) *TimerEntity {
	return nil
}

// 设置会话的标签
func (this *TimerEntity) SetTag(tag interface{}) *TimerEntity {
	return nil
}

// 设置会话的日志层级
func (this *TimerEntity) SetLog(log int) *TimerEntity {
	return nil
}
func procStart(proc *Proc) {
	return
}
func procUpdate() {
	return
}
func procStop(proc *Proc) {
	return
}

// 设置超时调用（时间以秒为单位）（务必在逻辑线程中调用或指定线程ID[tid]）
func SetTimeout(fun func(), timeout float32, tid ...int64) *TimerEntity {
	return nil
}

// 取消超时调用（务必在逻辑线程中调用或指定线程ID[tid]）
func ClearTimeout(id int, tid ...int64) {
	return
}

// 设置间歇调用（周期以秒为单位）（务必在逻辑线程中调用或指定线程ID[tid]）
func SetInterval(fun func(), interval float32, tid ...int64) *TimerEntity {
	return nil
}

// 取消间歇调用（务必在逻辑线程中调用或指定线程ID[tid]）
func ClearInterval(id int, tid ...int64) {
	return
}

// 在逻辑主线程中调用（返回的chan可用于阻塞当前线程）
func RunInMain(fun func()) chan bool {
	return nil
}

// 在指定逻辑线程中调用（返回的chan可用于阻塞当前线程）
func RunIn(tid int64, fun func()) chan bool {
	return nil
}

// 在当前逻辑线程中的下一帧调用
func RunInNext(fun func()) *TimerEntity {
	return nil
}
