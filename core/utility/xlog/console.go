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
	_ "os"
	_ "strings"
	"time"

	_ "github.com/shiena/ansicolor"
)

// brush is a color join function
type brush func(string) string

// newBrush return a fix color Brush
func newBrush(color string) brush {
	return nil
}

// Emergency          white
// Alert              cyan
// Critical           magenta
// Error              red
// Warning            yellow
// Notice             green
// Informational      blue
// Debug              Background blue
// consoleWriter implements LoggerInterface and writes messages to terminal.
type consoleWriter struct {
	Level    int  `json:"level"`
	Colorful bool `json:"color"` //this filed is useful only when system's terminal supports color
}

// NewConsole create ConsoleWriter returning as LoggerInterface.
func NewConsole() Logger {
	return nil
}

// Init init console logger.
// jsonConfig like '{"level":LevelTrace}'.
func (c *consoleWriter) Init(jsonConfig string) error {
	return nil
}

// WriteMsg write message in console.
func (c *consoleWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}

// Destroy implementing method. empty.
func (c *consoleWriter) Destroy() {
	return
}

// Flush implementing method. empty.
func (c *consoleWriter) Flush() {
	return
}
func init() {
	return
}
