//go:binary-only-package
package server

import (
	_ "esvr/core/utility/xstring"
	_ "fmt"
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
func (this *LanCfg) ServerID() string {
	return ""
}
func (this *LanCfg) ServerUrl() string {
	return ""
}
