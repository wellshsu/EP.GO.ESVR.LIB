//go:binary-only-package
package xjson

import (
	_ "bytes"
	_ "encoding/json"
	"errors"
	_ "esvr/core/utility/xstring"
)

var (
	TO_OBJ_INVALID_PARAM  = errors.New("xjson.ToObj invalid param")
	TO_OBJ_PARSE_ERROR    = errors.New("xjson.ToObj parse content fail")
	TO_OBJ_UNSUPPORT_DATA = errors.New("xjson.ToObj unsupport type")
)

func encodeNoEscape(v interface{}) ([]byte, error) {
	return []byte{}, nil
}

// 对象转字符串
func ToString(v interface{}) string {
	return ""
}

// 对象转字节数组
func ToByte(v interface{}) []byte {
	return []byte{}
}

// 字符串/字节数组转对象（指针）
func ToObj(data interface{}, obj interface{}) error {
	return nil
}
