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
	_ "fmt"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
)

var (
	GCgiProto string = "json" // cgi消息协议类型，可选pb/json，默认json
)

// 设置cgi消息协议类型，可选pb/json，默认json
func SetCgiProto(proto string) {
	return
}

// 封装cgi消息
func PackCgi(obj proto.Message) ([]byte, error) {
	return []byte{}, nil
}

// 解析cgi消息
func UnpackCgi(data []byte, obj proto.Message) error {
	return nil
}
