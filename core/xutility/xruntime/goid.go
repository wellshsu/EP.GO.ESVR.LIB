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
package xruntime

// 获取当前的GoroutineID
// TODO: 应当尽早使用context概念替换GOID逻辑
// 共有三个版本获取GOID
// 1.stack: 通过堆栈信息获取
// 2.go-tls: go 1.17.*版本无法使用go-tls，参考https://github.com/huandu/go-tls，[20211125待修复]
// 3.getg: 通过汇编代码获取g结构体，再确定goid的offset，用offset获取goid，参考https://github.com/v2pro/plz
// 以上性能比为: 90:3:1
func GoID(bystack ...bool) int64 {
	return 0
}
