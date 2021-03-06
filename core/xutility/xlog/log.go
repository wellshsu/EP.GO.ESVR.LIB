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
// Package logs provide a general log interface
// Usage:
//
// import "github.com/astaxie/beego/logs"
//
//	log := NewLogger(10000)
//	log.SetLogger("console", "")
//
//	> the first params stand for how many channel
//
// Use it like this:
//
//	log.Trace("trace")
//	log.Info("info")
//	log.Warn("warning")
//	log.Debug("debug")
//	log.Critical("critical")
//
//  more docs http://beego.me/docs/module/logs.md
package xlog

import (
	_ "fmt"
	"log"
	_ "os"
	_ "strings"
	_ "sync"
	"time"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

// levelLogLogger is defined to implement log.Logger
// the real log level will be LevelEmergency
// Name for adapter with beego official support
const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
	AdapterMail      = "smtp"
	AdapterConn      = "conn"
	AdapterEs        = "es"
	AdapterJianLiao  = "jianliao"
	AdapterSlack     = "slack"
	AdapterAliLS     = "alils"
)

// Legacy log level constants to ensure backwards compatibility.
const (
	LevelInfo  = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
)

type newLoggerFunc func() Logger

// Logger defines the behavior of a log provider.
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

// Register makes a log provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, log newLoggerFunc) {
	return
}

// BeeLogger is default logger in beego application.
// it can contain several providers and log message into all providers.
type BeeLogger struct {
}
type nameLogger struct {
	Logger
}
type logMsg struct {
}

func (this *logMsg) toString(prefix string) string {
	return ""
}

// NewLogger returns a new BeeLogger.
// channelLen means the number of messages in chan(used where asynchronous is true).
// if the buffering chan is full, logger adapters write to file or other way.
func NewLogger(channelLens ...int64) *BeeLogger {
	return nil
}

// Async set the log to asynchronous and start the goroutine
func (bl *BeeLogger) Async(msgLen ...int64) *BeeLogger {
	return nil
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func (bl *BeeLogger) setLogger(adapterName string, configs ...string) error {
	return nil
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func (bl *BeeLogger) SetLogger(adapterName string, configs ...string) error {
	return nil
}

// DelLogger remove a logger adapter in BeeLogger.
func (bl *BeeLogger) DelLogger(adapterName string) error {
	return nil
}
func (bl *BeeLogger) writeToLoggers(when time.Time, msg string, level int) {
	return
}
func (bl *BeeLogger) Write(p []byte) (n int, err error) {
	return n, err
}
func (bl *BeeLogger) WriteMsg(logLevel int, msg string, v ...interface{}) error {
	return nil
}

// SetLevel Set log message level.
// If message level (such as LevelDebug) is higher than logger level (such as LevelWarning),
// log providers will not even be sent the message.
func (bl *BeeLogger) SetLevel(l int) {
	return
}

// GetLevel Get Current log message level.
func (bl *BeeLogger) GetLevel() int {
	return 0
}

// SetLogFuncCallDepth set log funcCallDepth
func (bl *BeeLogger) SetLogFuncCallDepth(d int) {
	return
}

// GetLogFuncCallDepth return log funcCallDepth for wrapper
func (bl *BeeLogger) GetLogFuncCallDepth() int {
	return 0
}

// EnableFuncCallDepth enable log funcCallDepth
func (bl *BeeLogger) EnableFuncCallDepth(b bool) {
	return
}

// set prefix
func (bl *BeeLogger) SetPrefix(s string) {
	return
}

// start logger chan reading.
// when chan is not empty, write logs.
func (bl *BeeLogger) startLogger() {
	return
}

// Emergency Log EMERGENCY level message.
func (bl *BeeLogger) Emergency(format string, v ...interface{}) {
	return
}

// Alert Log ALERT level message.
func (bl *BeeLogger) Alert(format string, v ...interface{}) {
	return
}

// Critical Log CRITICAL level message.
func (bl *BeeLogger) Critical(format string, v ...interface{}) {
	return
}

// Error Log ERROR level message.
func (bl *BeeLogger) Error(format string, v ...interface{}) {
	return
}

// Warning Log WARNING level message.
func (bl *BeeLogger) Warning(format string, v ...interface{}) {
	return
}

// Notice Log NOTICE level message.
func (bl *BeeLogger) Notice(format string, v ...interface{}) {
	return
}

// Informational Log INFORMATIONAL level message.
func (bl *BeeLogger) Informational(format string, v ...interface{}) {
	return
}

// Debug Log DEBUG level message.
func (bl *BeeLogger) Debug(format string, v ...interface{}) {
	return
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func (bl *BeeLogger) Warn(format string, v ...interface{}) {
	return
}

// Info Log INFO level message.
// compatibility alias for Informational()
func (bl *BeeLogger) Info(format string, v ...interface{}) {
	return
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func (bl *BeeLogger) Trace(format string, v ...interface{}) {
	return
}

// Flush flush all chan data.
func (bl *BeeLogger) Flush() {
	return
}

// Close close logger, flush all chan data and destroy all adapters in BeeLogger.
func (bl *BeeLogger) Close() {
	return
}

// Reset close all outputs, and set bl.outputs to nil
func (bl *BeeLogger) Reset() {
	return
}
func (bl *BeeLogger) flush() {
	return
}

// beeLogger references the used application logger.
// GetBeeLogger returns the default BeeLogger
func GetBeeLogger() *BeeLogger {
	return nil
}

// GetLogger returns the default BeeLogger
func GetLogger(prefixes ...string) *log.Logger {
	return nil
}

// Reset will remove all the adapter
func Reset() {
	return
}

// Async set the beelogger with Async mode and hold msglen messages
func Async(msgLen ...int64) *BeeLogger {
	return nil
}

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	return
}

// SetPrefix sets the prefix
func SetPrefix(s string) {
	return
}

// EnableFuncCallDepth enable log funcCallDepth
func EnableFuncCallDepth(b bool) {
	return
}

// SetLogFuncCall set the CallDepth, default is 4
func SetLogFuncCall(b bool) {
	return
}

// SetLogFuncCallDepth set log funcCallDepth
func SetLogFuncCallDepth(d int) {
	return
}

// SetLogger sets a new logger.
func SetLogger(adapter string, config ...string) error {
	return nil
}

// Emergency logs a message at emergency level.
func Emergency(f interface{}, v ...interface{}) {
	return
}

// Alert logs a message at alert level.
func Alert(f interface{}, v ...interface{}) {
	return
}

// Critical logs a message at critical level.
func Critical(f interface{}, v ...interface{}) {
	return
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	return
}

// Warning logs a message at warning level.
func Warning(f interface{}, v ...interface{}) {
	return
}

// Warn compatibility alias for Warning()
func Warn(f interface{}, v ...interface{}) {
	return
}

// Notice logs a message at notice level.
func Notice(f interface{}, v ...interface{}) {
	return
}

// Informational logs a message at info level.
func Informational(f interface{}, v ...interface{}) {
	return
}

// Info compatibility alias for Warning()
func Info(f interface{}, v ...interface{}) {
	return
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	return
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(f interface{}, v ...interface{}) {
	return
}
func Format(f interface{}, v ...interface{}) string {
	return ""
}
