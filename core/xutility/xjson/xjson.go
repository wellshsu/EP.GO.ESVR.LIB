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
package xjson

import (
	_ "bytes"
	_ "encoding/json"
	"errors"
	_ "fmt"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xstring"
)

var (
	TO_OBJ_INVALID_PARAM  = errors.New("xjson.ToObj invalid param")
	TO_OBJ_PARSE_ERROR    = errors.New("xjson.ToObj parse content fail")
	TO_OBJ_UNSUPPORT_DATA = errors.New("xjson.ToObj unsupport type")
)

func encodeNoEscape(v interface{}) ([]byte, error) {
	return []byte{}, nil
}

// 对象转字符串输出
func ToPrint(v interface{}) string {
	return ""
}

// 对象转字符串
func ToString(v interface{}) (string, error) {
	return "", nil
}

// 对象转字节数组
func ToByte(v interface{}) ([]byte, error) {
	return []byte{}, nil
}

// 字符串/字节数组转对象（指针）
func ToObj(data interface{}, obj interface{}) error {
	return nil
}
