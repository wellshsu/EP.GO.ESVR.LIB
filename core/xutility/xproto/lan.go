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
package xproto

import (
	_ "errors"
	_ "fmt"
	_ "math"

	"github.com/golang/protobuf/proto"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "google.golang.org/protobuf/proto"
)

var (
	GUrl  string  = ""  // 当前线路地址
	PGUrl *string = nil // 当前线路地址（指针）
)

// 设置线路全局Url
func SetGUrl(url string) {
	return
}

// 封装protobuf
func PackProto(obj proto.Message) []byte {
	return []byte{}
}

// 解析protobuf
func UnpackProto(data []byte, obj proto.Message) error {
	return nil
}

// 封装GID（l-左区间，r-右区间，皆为闭区间）
func PackGID(l int, r int) int {
	return 0
}

// 解析GID（l-左区间，r-右区间，皆为闭区间）
func UnpackGID(g int) (l, r int) {
	return l, r
}

// 封装网络帧
func PackFrame(frame IFrame) ([]byte, error) {
	return []byte{}, nil
}

// 解析网络帧
func UnpackFrame(bytes []byte) (byte, IFrame, error) {
	return 0, nil, nil
}

// 转换网络帧（src和dst互换）
func ShiftFrame(src IFrame) IFrame {
	return nil
}

// 拷贝网络帧
func CloneFrame(frame IFrame) IFrame {
	return nil
}

// 缓存网络帧
func PoolFrame(frame IFrame) {
	return
}
