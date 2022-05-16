//go:binary-only-package
package xlog

import (
	_ "encoding/json"
	_ "esvr/core/utility/xjson"
	_ "fmt"
	_ "io/ioutil"
	_ "strings"
)

const (
	MAX_CHAN_LEN = 100000
)

type LogCfg struct {
	Filename string `json:"filename,omitempty"`
	MaxLines int    `json:"maxlines,omitempty"` // Rotate at line
	MaxSize  int    `json:"maxsize,omitempty"`  // Rotate at size
	Daily    bool   `json:"daily,omitempty"`    // Rotate daily
	MaxDays  int64  `json:"maxdays,omitempty"`
	Rotate   bool   `json:"rotate,omitempty"`
	Level    int    `json:"level,omitempty"`
	Perm     string `json:"perm,omitempty"`
}

func Init(cfgFile string) *BeeLogger {
	return nil
}
func Flush() {
	return
}
func Close() {
	return
}
func ChanSize() int {
	return 0
}

// 触发异常（Crash）
func Panic(v ...interface{}) {
	return
}
