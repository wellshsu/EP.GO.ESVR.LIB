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
package xstring

import (
	_ "fmt"
	_ "strconv"
	_ "strings"
	_ "unsafe"
)

// 字符串转int
func ToInt(str string) int {
	return 0
}

// int转字符串
func ToString(itr int) string {
	return ""
}

// int64转字符串
func Int64ToString(itr int64) string {
	return ""
}

// 保留指定位数的小数（参数为float32或float64）（默认保留两位小数）
func ToFixed(float interface{}, fixed ...int) string {
	return ""
}

// 字符串分割
func Split(str string, sep string) []string {
	return []string{}
}

// 找到指定字符的索引
func IndexOf(str string, of string) int {
	return 0
}

// 找到指定字符的索引（后）
func LastIndexOf(str string, of string) int {
	return 0
}

// 是否以指定字符起始
func StartWith(str string, of string) bool {
	return false
}

// 是否以指定字符结束
func EndWith(str string, of string) bool {
	return false
}

// 是否包含指定字符
func Contains(str string, of string) bool {
	return false
}

// 是否为空
func IsEmpty(str string) bool {
	return false
}

// 截取
func Sub(str string, from int, to int) string {
	return ""
}

// 替换所有指定字符
func Replace(str string, from string, to string) string {
	return ""
}

// 剔除多余的空格
func Trim(str string) string {
	return ""
}

// 字符串转字节数组
func StrToBytes(s string) []byte {
	return []byte{}
}

// 字节数组转字符串
func BytesToStr(b []byte) string {
	return ""
}

// 字符串格式化
func Format(format string, args ...interface{}) string {
	return ""
}
