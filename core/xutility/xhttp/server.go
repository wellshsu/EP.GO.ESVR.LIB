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
package xhttp

import (
	"net/http"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
)

type Server struct {
	http.Server
	Key  string
	Cert string
}

func NewServer() *Server {
	return nil
}

// 设置连接地址和端口（:PORT）
func (this *Server) SetAddr(addr string) *Server {
	return nil
}

// 设置Https密钥和证书（若未设置则启动http监听）
// 生成密钥：openssl genrsa -out server.key 2048
// 生成证书：openssl req -new -x509 -key server.key -out server.crt -days 365
// 参数说明：Common Name = localhost/domain/ip
func (this *Server) SetHttps(key string, cert string) *Server {
	return nil
}

// 设置请求回调
func (this *Server) SetHandler(handler http.HandlerFunc) *Server {
	return nil
}

// 启动服务
func (this *Server) Start() *Server {
	return nil
}

// 停止服务
func (this *Server) Stop() {
	return
}
