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
// TONOTE: sync.Pool的对象生命周期位于两次GC之间；
// 实现的逻辑是绑定P，故而跨协程存取会消耗一些性能(偷缓存)；
// 测试：CgiResp, New-53, Get-4630, 复用率>90%；
package xproto

import (
	_ "sync"
)

// 重置路由
func ResetRoute(route *Route) {
	return
}

// 克隆路由（浅拷贝）
func CloneRoute(src *Route, dst *Route) {
	return
}

// 从缓存中获取*LoopReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func GetLoopReq() *LoopReq {
	return nil
}

// 缓存*LoopReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func PutLoopReq(obj *LoopReq) {
	return
}

// 从缓存中获取*MsgReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func GetMsgReq() *MsgReq {
	return nil
}

// 缓存*MsgReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func PutMsgReq(obj *MsgReq) {
	return
}

// 从缓存中获取*RpcReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func GetRpcReq() *RpcReq {
	return nil
}

// 缓存*RpcReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func PutRpcReq(obj *RpcReq) {
	return
}

// 从缓存中获取*RpcResp对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func GetRpcResp() *RpcResp {
	return nil
}

// 缓存*RpcResp对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func PutRpcResp(obj *RpcResp) {
	return
}

// 从缓存中获取*CgiReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func GetCgiReq() *CgiReq {
	return nil
}

// 缓存*CgiReq对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func PutCgiReq(obj *CgiReq) {
	return
}

// 从缓存中获取*CgiResp对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func GetCgiResp() *CgiResp {
	return nil
}

// 缓存*CgiResp对象（任何地方都不应对其保持引用，若需引用请使用xmsg.CloneFrame）
func PutCgiResp(obj *CgiResp) {
	return
}
