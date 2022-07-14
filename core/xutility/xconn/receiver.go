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
package xconn

import (
	"errors"
	_ "io"
	"net"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xmath"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xproto"
)

const (
	RECEIVER_CHAN_LENGTH = 512
)

var (
	ERR_RECEIVER_CHAN_FULL          = errors.New("receiver chan is full")
	ERR_RECEIVER_INVALID_HEAD       = errors.New("receiver head is invalid")
	ERR_RECEIVER_LARGE_BODY_SIZE    = errors.New("receiver body is too large")
	ERR_RECEIVER_NEGATIVE_BODY_SIZE = errors.New("receiver body is negative")
)

type Receiver struct {
	ChMsg   chan []byte
	MaxBody int
}

func NewReceiver(maxBody int) *Receiver {
	return nil
}
func (this *Receiver) Recv(c net.Conn) (int64, error) {
	return 0, nil
}
func (this *Receiver) Clear() {
	return
}
