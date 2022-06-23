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
package xserver

import (
	"errors"
	_ "fmt"
	"net/http"
	_ "time"

	"github.com/golang/protobuf/proto"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmath"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmsg"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xruntime"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xsession"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

var (
	ERR_SEND_CHAN_FULL  = errors.New("send chan is full")
	ERR_NO_ROUTE_FOUND  = errors.New("no route found")
	ERR_RPC_TIMEOUT     = errors.New("rpc call timeout")
	ERR_CGI_TIMEOUT     = errors.New("cgi call timeout")
	ERR_RPC_INTERRUPTED = errors.New("rpc call has been interrupted, see log context for more details.")
	ERR_CGI_INTERRUPTED = errors.New("cgi call has been interrupted, see log context for more details.")
)

func pushFrame(goNum int) {
	return
}

// [发送网络帧（根据UID负载均衡）]
func SendFrame(frame xmsg.IFrame) bool {
	return false
}
func doRpc(rid int, uid int, req interface{}, addr string, offsetAndTimeout ...int) (*xmsg.RpcResp, error) {
	return nil, nil
}

// [发送Rpc消息（同步）]
// [id-消息ID]
// [uid-用户ID（负载均衡）]
// [req-请求结构体]
// [resp-返回结构体]
// [addr-目标服务器]
// [offset-目标协程偏移（基于protocol中定义）]
// [timeout-超时时长]
func SendSync(id int, uid int, req proto.Message, resp proto.Message, addr string, offsetAndTimeout ...int) error {
	return nil
}

// [发送Rpc消息（异步）]
// [id-消息ID]
// [uid-用户ID（负载均衡）]
// [req-请求结构体]
// [addr-目标服务器]
// [callback-回调函数]
// [offset-目标协程ID偏移（基于protocol中定义）]
// [timeout-超时时长]
func SendAsync(id int, uid int, req proto.Message, addr string, callback func(frame *xmsg.RpcResp, err error), offsetAndTimeout ...int) {
	return
}

// [发送Msg消息]
// [id-消息ID]
// [msg-结构体]
// [mreq-msg帧]
func SendMsg(id int, msg proto.Message, mreq *xmsg.MsgReq) bool {
	return false
}

// [发送Cgi消息（同步，否则ResponseWriter无法输出）]
// [id-消息ID]
// [uid-用户ID（负载均衡）]
// [req-请求结构体]
// [addr-目标服务器]
// [timeout-超时时长]
func SendCgi(id int, uid int, req *http.Request, addr string, timeout ...int) (cresp *xmsg.CgiResp, err error) {
	return cresp, err
}
