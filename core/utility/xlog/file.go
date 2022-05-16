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
	_ "errors"
	_ "fmt"
	_ "io"
	"os"
	_ "path"
	_ "path/filepath"
	_ "strconv"
	_ "strings"
	"sync"
	"time"
)

// fileLogWriter implements LoggerInterface.
// It writes messages by lines limit, file size limit, or time frequency.
type fileLogWriter struct {
	sync.RWMutex // write log order by order and  atomic incr maxLinesCurLines and maxSizeCurSize
	// The opened file
	Filename string `json:"filename"`
	// Rotate at line
	MaxLines         int `json:"maxlines"`
	MaxFiles         int `json:"maxfiles"`
	MaxFilesCurFiles int
	// Rotate at size
	MaxSize int `json:"maxsize"`
	// Rotate daily
	Daily   bool  `json:"daily"`
	MaxDays int64 `json:"maxdays"`
	// Rotate hourly
	Hourly     bool   `json:"hourly"`
	MaxHours   int64  `json:"maxhours"`
	Rotate     bool   `json:"rotate"`
	Level      int    `json:"level"`
	Perm       string `json:"perm"`
	RotatePerm string `json:"rotateperm"`
	// like "project.log", project is fileNameOnly and .log is suffix
}

// newFileWriter create a FileLogWriter returning as LoggerInterface.
func newFileWriter() Logger {
	return nil
}

// Init file logger with json config.
// jsonConfig like:
//  {
//  "filename":"logs/beego.log",
//  "maxLines":10000,
//  "maxsize":1024,
//  "daily":true,
//  "maxDays":15,
//  "rotate":true,
//      "perm":"0600"
//  }
func (w *fileLogWriter) Init(jsonConfig string) error {
	return nil
}

// start file logger. create log file and set to locker-inside file writer.
func (w *fileLogWriter) startLogger() error {
	return nil
}
func (w *fileLogWriter) needRotateDaily(size int, day int) bool {
	return false
}
func (w *fileLogWriter) needRotateHourly(size int, hour int) bool {
	return false
}

// WriteMsg write logger message into file.
func (w *fileLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}
func (w *fileLogWriter) createLogFile() (*os.File, error) {
	return nil,
		nil
}
func (w *fileLogWriter) initFd() error {
	return nil
}
func (w *fileLogWriter) dailyRotate(openTime time.Time) {
	return
}
func (w *fileLogWriter) hourlyRotate(openTime time.Time) {
	return
}
func (w *fileLogWriter) lines() (int, error) {
	return 0, nil
}

// DoRotate means it need to write file in new file.
// new file name like xx.2013-01-01.log (daily) or xx.001.log (by line or size)
func (w *fileLogWriter) doRotate(logTime time.Time) error {
	return nil
}
func (w *fileLogWriter) deleteOldLog() {
	return
}

// Destroy close the file description, close file writer.
func (w *fileLogWriter) Destroy() {
	return
}

// Flush flush file logger.
// there are no buffering messages in file logger in memory.
// flush file means sync file from disk.
func (w *fileLogWriter) Flush() {
	return
}
func init() {
	return
}
