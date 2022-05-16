//go:binary-only-package
// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package xlog

import (
	_ "encoding/json"
	_ "io"
	_ "net"
	"time"
)

// connWriter implements LoggerInterface.
// it writes messages in keep-live tcp connection.
type connWriter struct {
	ReconnectOnMsg bool   `json:"reconnectOnMsg"`
	Reconnect      bool   `json:"reconnect"`
	Net            string `json:"net"`
	Addr           string `json:"addr"`
	Level          int    `json:"level"`
}

// NewConn create new ConnWrite returning as LoggerInterface.
func NewConn() Logger {
	return nil
}

// Init init connection writer with json config.
// json config only need key "level".
func (c *connWriter) Init(jsonConfig string) error {
	return nil
}

// WriteMsg write message in connection.
// if connection is down, try to re-connect.
func (c *connWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}

// Flush implementing method. empty.
func (c *connWriter) Flush() {
	return
}

// Destroy destroy connection writer and close tcp listener.
func (c *connWriter) Destroy() {
	return
}
func (c *connWriter) connect() error {
	return nil
}
func (c *connWriter) needToConnectOnMsg() bool {
	return false
}
func init() {
	return
}
