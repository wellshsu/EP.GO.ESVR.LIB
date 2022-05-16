//go:binary-only-package
package xfs

import (
	_ "esvr/core/utility/xlog"
	_ "esvr/core/utility/xstring"
	_ "io/ioutil"
	"os"
	_ "path/filepath"
)

// 删除文件
func DeleteFile(path string) bool {
	return false
}

// 判断文件是否存在
func FileExist(path string) bool {
	return false
}

// 判断文件夹是否存在
func PathExist(path string, createIsNotExist ...bool) bool {
	return false
}

// 根据文件名称获取该文件的文件夹名称
func Directory(file string) string {
	return ""
}

// 读取文件（以字符串形式返回）
func ReadString(file string) string {
	return ""
}

// 读取文件（以字节数组形式返回）
func ReadBytes(file string) []byte {
	return []byte{}
}

// 写入文件（字符串）
func WriteString(file string, value string, _mode ...os.FileMode) {
	return
}

// 写入文件（字节数组）
func WriteBytes(file string, value []byte, _mode ...os.FileMode) {
	return
}
