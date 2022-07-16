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
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xmath"
)

const (
	NET_VERSION         = 1
	NET_HEAD_LENGTH     = 19
	NET_VERSION_OFFSET  = 2
	NET_LENGTH_OFFSET   = 3
	NET_ID_OFFSET       = 7
	NET_PLAYERID_OFFSET = 11
	NET_SERVERID_OFFSET = 15
	NET_MESSAGE_OFFSET  = 19
	NET_TAG_1           = 8
	NET_TAG_2           = 8
)

var (
	GMsgProto string = "pb" // msg消息协议类型，可选pb/json，默认pb
)

// 设置msg消息协议类型，可选pb/json，默认pb
func SetMsgProto(proto string) {
	return
}

// 解析消息ID（for msg）
func UnpackID(buf []byte) int {
	return 0
}

// 解析用户ID（for msg）
func UnpackUID(buf []byte) int {
	return 0
}

// 解析服务器ID（for msg）
func UnpackSID(buf []byte) int {
	return 0
}

// 封装msg消息
func PackMsg(id int, obj proto.Message) ([]byte, error) {
	return []byte{}, nil
}

// 解析msg消息
func UnpackMsg(buf []byte, obj proto.Message) error {
	return nil
}
