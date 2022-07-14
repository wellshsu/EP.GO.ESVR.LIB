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

import _ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"

type Route struct {
	ID   int      // 路由ID
	Name string   // 路由名称
	GoL  int      // 协程ID（左）
	GoR  int      //协程ID（右）
	RW   bool     // 可读可写（默认true）
	Log  int      // 日志层级（参考xlog的LogLevel）
	Dst  []string // 目标
}

// 获取日志层级，若未指定则使用全局日志层级
func (this *Route) GetLog() int {
	return 0
}

type MsgRoute struct {
	Route
}

var MSGROUTEMAP map[int]*MsgRoute // msg路由
func RegMsgRoute(_map map[int]*MsgRoute) {
	return
}

type RpcRoute struct {
	Route
}

var RPCROUTEMAP map[int]*RpcRoute // rpc路由
func RegRpcRoute(_map map[int]*RpcRoute) {
	return
}

type CgiRoute struct {
	Route
	Method  []string
	Timeout int
}

var CGIROUTEMAP map[int]*CgiRoute // cgi路由
func RegCgiRoute(_map map[int]*CgiRoute) {
	return
}
