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
	"net"
)

const (
	SENDER_CHAN_LENGTH = 512
)

var (
	ERR_SENDER_FULL = errors.New("sender chan is full")
)

type Sender struct {
	ChMsg  chan []byte
	ChSend chan bool
}

func NewSender() *Sender {
	return nil
}
func (this *Sender) Send(conn net.Conn) error {
	return nil
}
func (this *Sender) Write(data []byte) error {
	return nil
}
func (this *Sender) Clear() {
	return
}
