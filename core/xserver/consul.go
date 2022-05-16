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
	CONSUL_RESP_OK    = "ok"
	CONSUL_CHECK_PATH = "/check"
)

var (
	CslClt *consulapi.Client
	CslCfg *consulapi.Config
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
func PostKV(key string, value string, version string, block ...bool) bool {
	return false
}
func PullKV(key string) []byte {
	return []byte{}
}
func UpdateKV(key string, interval int, onUpdate func(data []byte)) {
	return
}
