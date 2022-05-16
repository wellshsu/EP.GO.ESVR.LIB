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
	"time"
)

// A filesLogWriter manages several fileLogWriter
// filesLogWriter will write logs to the file in json configuration  and write the same level log to correspond file
// means if the file name in configuration is project.log filesLogWriter will create project.error.log/project.debug.log
// and write the error-level logs to project.error.log and write the debug-level logs to project.debug.log
// the rotate attribute also  acts like fileLogWriter
type multiFileLogWriter struct {
	// the last one for fullLogWriter
	Separate []string `json:"separate"`
}

// Init file logger with json config.
// jsonConfig like:
//	{
//	"filename":"logs/beego.log",
//	"maxLines":0,
//	"maxsize":0,
//	"daily":true,
//	"maxDays":15,
//	"rotate":true,
//  	"perm":0600,
//	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
//	}
func (f *multiFileLogWriter) Init(config string) error {
	return nil
}
func (f *multiFileLogWriter) Destroy() {
	return
}
func (f *multiFileLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}
func (f *multiFileLogWriter) Flush() {
	return
}

// newFilesWriter create a FileLogWriter returning as LoggerInterface.
func newFilesWriter() Logger {
	return nil
}
func init() {
	return
}
