//go:binary-only-package
package xsession

import (
	_ "fmt"
	_ "sync"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
)

const (
	PIPE_CHAN_MAX = 100000 // 推送队列最大长度
)

// map[string]map[string]*GlobalObject
// map[string]bool
// map[string]map[string]int
// map[int]map[string]map[string]*SessionObject
// map[int]map[string]bool
// map[string]*DeleteRecord
// map[int64]int64
// map[int64]*ProfilerRecord
// map[int64]*PipeRecord
type SessionObject struct {
	Raw     xorm.ITable
	Ptr     xorm.ITable
	SWrite  bool
	SDelete bool
	SClear  bool
	Cond    *xorm.Condition
}
type GlobalObject struct {
	Ptr     xorm.ITable
	SDelete bool
}
type PipeObject struct {
	Raw     xorm.ITable
	SWrite  bool
	SDelete bool
	SClear  bool
	SDB     bool
	SRedis  bool
	Cond    *xorm.Condition
}
type PipeRecord struct {
	ID      int64
	Tag     interface{}
	Time    int
	Objects []*PipeObject
}
type DeleteRecord struct {
	Chan  chan bool
	Count int32
}
type DeferFun func()
type Session struct {
	ID       int
	TID      int
	Profiler *ProfilerRecord
}
type ProfilerRecord struct {
	Tag  interface{}
	Time int
	RW   bool
	Log  int
	TID  int64
	SID  int64
}

func (this *ProfilerRecord) Reset() {
	return
}
func (this *ProfilerRecord) LogPrefix() string {
	return ""
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogEmergency(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogAlert(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogCritical(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogError(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogWarn(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogNotice(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogInfo(f interface{}, v ...interface{}) {
	return
}

// Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogDebug(f interface{}, v ...interface{}) {
	return
}
