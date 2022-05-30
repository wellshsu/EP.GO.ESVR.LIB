//go:binary-only-package
package xserver

import (
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io/ioutil"
	_ "net/http"
	_ "strings"
	_ "time"

	consulapi "github.com/hashicorp/consul/api"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
)

const (
	CONSUL_RESP_OK    = "ok"     // Consul响应200
	CONSUL_CHECK_PATH = "/check" // Consul心跳检测
)

var (
	CslClt *consulapi.Client // Consul连接
	CslCfg *consulapi.Config // Consul配置
)

func startConsul(onUpdate func(map[string][]string)) {
	return
}
func httpConsul(addr string) {
	return
}
func regConsul(name string, nodeID string, addr string, port int, regAddr string, timeout string, interval string, unregAfter string) error {
	return nil
}
func unregConsul(nodeID string) error {
	return nil
}
func pullConsul(url string, servers string, onUpdate func(map[string][]string)) {
	return
}

// 推送KV（键值对）至Consul Storage（version-版本号，以'VERSION_'为前缀，block-是否阻塞）
func PostKV(key string, value string, version string, block ...bool) bool {
	return false
}

// 从Consul Storage中拉取KV（键值对）（阻塞）
func PullKV(key string) []byte {
	return []byte{}
}

// 订阅Consul Storage中的KV（键值对）（阻塞）（interval-间歇时间）（注意订阅的Key需要设置版本，以'VERSION_'为前缀）
func SubKV(key string, interval int, onUpdate func(data []byte)) {
	return
}
