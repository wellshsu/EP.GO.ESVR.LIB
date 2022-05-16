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
	"io"
	_ "runtime"
	"sync"
	"time"
)

type logWriter struct {
	sync.Mutex
}

func newLogWriter(wr io.Writer) *logWriter {
	return nil
}
func (lg *logWriter) writeln(when time.Time, msg string) {
	return
}
func formatTimeHeader(when time.Time) ([]byte, int, int) {
	return []byte{}, 0, 0
}
func initColor() {
	return
}

// ColorByStatus return color by http code
// 2xx return Green
// 3xx return White
// 4xx return Yellow
// 5xx return Red
func ColorByStatus(code int) string {
	return ""
}

// ColorByMethod return color by http code
func ColorByMethod(method string) string {
	return ""
}

// ResetColor return reset color
func ResetColor() string {
	return ""
}
