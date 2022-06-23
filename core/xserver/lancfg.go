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
	_ "fmt"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
)

type LanCfg struct {
	Name  string // 名称
	Addr  string // tcp://$ip:$port
	Raw   string // $ip:$port
	IP    string // IP
	Port  int    // 端口
	GO    int    // 逻辑线程数
	MaxRx int    // 最大接收字节数（KB）
}

func NewLanCfg(name, addr string) *LanCfg {
	return nil
}

// 服务器ID（$name@tcp://$ip:$port）
func (this *LanCfg) ServerID() string {
	return ""
}
