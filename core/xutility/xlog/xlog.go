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
package xlog

import (
	_ "encoding/json"
	_ "fmt"
	_ "io/ioutil"
	_ "strings"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xjson"
)

const (
	MAX_CHAN_LEN = 300000
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
