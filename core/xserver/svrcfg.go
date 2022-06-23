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
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xconfig"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
)

type SvrCfg struct {
	Raw              xconfig.Configer
	Env              string // 环境标识: 测试，内测，生产
	LanCfg           *LanCfg
	LinkServer       string // 需要连接的内部服务器
	ConsulAddr       string // Consul中心地址
	ConsulHttp       string // Consul检测地址
	ConsulTimeout    string // Consul超时时间
	ConsulInterval   string // Consul访问间隔
	ConsulDeregister string // Consul延迟注销
}

func (this *SvrCfg) Init(config string) bool {
	return false
}
func (this *SvrCfg) SvrID() string {
	return ""
}
func (this *SvrCfg) SvrName() string {
	return ""
}
func (this *SvrCfg) IsDebugEnv() bool {
	return false
}
