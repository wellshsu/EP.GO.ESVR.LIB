//go:binary-only-package
package xmsg

import (
	_ "fmt"
	_ "math"

	_ "github.com/golang/protobuf/proto"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmath"
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
	GUrl string = ""
)

// 设置全局url
func SetGUrl(url string) {
	return
}

// 解析消息ID
func ParseID(buf []byte) int {
	return 0
}

// 解析用户ID
func ParseUID(buf []byte) int {
	return 0
}

// 解析网络帧
func ParseFrame(frame *Frame, out interface{}) error {
	return nil
}

// 解析消息
func ParseMsg(buf []byte, out interface{}) error {
	return nil
}

// 封装消息
func PackMsg(id int, in interface{}) ([]byte, error) {
	return []byte{}, nil
}

// 转换网络帧（src和dest互换）
func ShiftFrame(frame *Frame) *Frame {
	return nil
}

// 使用dest和uid构建一个简易的网络帧
func SimpleFrame(dest string, uid int) *Frame {
	return nil
}

// 封装GoID（l-左区间，r-右区间，皆为闭区间）
func PackGo(l int, r int) int {
	return 0
}

// 解析GoID（l-左区间，r-右区间，皆为闭区间）
func ParseGo(g int) (l, r int) {
	return l, r
}
