//go:binary-only-package
package xmsg

import (
	_ "errors"
	_ "fmt"
	_ "math"

	"github.com/golang/protobuf/proto"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmath"
	_ "google.golang.org/protobuf/proto"
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

// 解析消息ID（for msg）
func UnpackID(buf []byte) int {
	return 0
}

// 解析用户ID（for msg）
func UnpackUID(buf []byte) int {
	return 0
}

// 解析消息（for msg）
func UnpackMsg(buf []byte, out interface{}) error {
	return nil
}

// 封装消息（for msg）
func PackMsg(id int, in interface{}) ([]byte, error) {
	return []byte{}, nil
}

// 封装protobuf
func PackProto(obj proto.Message) []byte {
	return []byte{}
}

// 解析protobuf
func UnpackProto(data []byte, obj proto.Message) error {
	return nil
}
func NewLoopReq() *LoopReq {
	return nil
}
func NewMsgReq() *MsgReq {
	return nil
}
func NewRpcReq() *RpcReq {
	return nil
}
func NewRpcResp() *RpcResp {
	return nil
}
func NewCgiReq() *CgiReq {
	return nil
}
func NewCgiResp() *CgiResp {
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
func ShiftFrame(frame IFrame) IFrame {
	return nil
}
