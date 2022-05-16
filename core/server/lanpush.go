//go:binary-only-package
package server

import (
	"errors"
	_ "esvr/core/utility/xlog"
	_ "esvr/core/utility/xmath"
	"esvr/core/utility/xmsg"
	_ "esvr/core/utility/xrun"
	_ "esvr/core/utility/xruntime"
	_ "esvr/core/utility/xsession"
	_ "esvr/core/utility/xstring"
	_ "esvr/core/utility/xtime"
	_ "fmt"
	_ "time"

	"github.com/golang/protobuf/proto"
)

var (
	ERR_SEND_CHAN_FULL = errors.New("send chan is full")
	ERR_DO_RPC_TIMEOUT = errors.New("do rpc timeout")
	ERR_NO_ROUTE_FOUND = errors.New("no route found")
)

func pushMsg(goNum int) {
	return
}
func doRpc(id int, uid int, req interface{}, addr string, offsetAndTimeout ...int) (*xmsg.Frame, error) {
	return nil, nil
}

// [发送消息（同步），用于各服务间通讯]
// [id-消息ID]
// [uid-用户ID]
// [req-请求结构体]
// [resp-返回结构体]
// [addr-目标服务器]
// [offset-目标协程偏移（基于protocol中定义）]
// [timeout-超时时长]
func SendSync(id int, uid int, req proto.Message, resp proto.Message, addr string, offsetAndTimeout ...int) error {
	return nil
}

// [发送消息（异步），用于各服务间通讯]
// [id-消息ID]
// [uid-用户ID]
// [req-请求结构体]
// [addr-目标服务器]
// [callback-回调函数]
// [offset-目标协程ID偏移（基于protocol中定义）]
// [timeout-超时时长]
func SendAsync(id int, uid int, req proto.Message, addr string, callback func(frame *xmsg.Frame, err error), offsetAndTimeout ...int) {
	return
}

// [发送消息，用于各服务间通讯]
// [id-消息ID]
// [msg-结构体]
// [frame-网络帧]
func SendMsg(id int, msg proto.Message, frame *xmsg.Frame) bool {
	return false
}

// [发送网络帧，用于各服务间通讯]
func SendFrame(frame *xmsg.Frame) bool {
	return false
}
