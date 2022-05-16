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
	_ "bytes"
	_ "encoding/json"
	_ "fmt"
	_ "strings"
	"time"
)

// AccessLogRecord struct for holding access log data.
type AccessLogRecord struct {
	RemoteAddr     string        `json:"remote_addr"`
	RequestTime    time.Time     `json:"request_time"`
	RequestMethod  string        `json:"request_method"`
	Request        string        `json:"request"`
	ServerProtocol string        `json:"server_protocol"`
	Host           string        `json:"host"`
	Status         int           `json:"status"`
	BodyBytesSent  int64         `json:"body_bytes_sent"`
	ElapsedTime    time.Duration `json:"elapsed_time"`
	HTTPReferrer   string        `json:"http_referrer"`
	HTTPUserAgent  string        `json:"http_user_agent"`
	RemoteUser     string        `json:"remote_user"`
}

func (r *AccessLogRecord) json() ([]byte, error) {
	return []byte{}, nil
}
func disableEscapeHTML(i interface{}) {
	return
}

// AccessLog - Format and print access log.
func AccessLog(r *AccessLogRecord, format string) {
	return
}
