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
package xsession

import (
	_ "errors"
	_ "fmt"
	_ "reflect"
	_ "strconv"
	_ "strings"
	"sync"

	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xruntime"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

// 获取当前线程的ID
func GID() int {
	return 0
}

// 获取会话的profiler
func MProfiler(id int64) *ProfilerRecord {
	return nil
}

// 获取全局内存map
func MGMem(prefix string, model xorm.ITable) *sync.Map {
	return nil
}

// 获取会话内存map
func MSMem(prefix string, model xorm.ITable) *sync.Map {
	return nil
}

// 获取会话的pipe
func MPipe() chan *PipeRecord {
	return nil
}

// 获取全局的最大表ID
func MGMax(model xorm.ITable) *sync.Map {
	return nil
}

// Pipe队列等待个数
func MPCount() int64 {
	return 0
}

// 将数据保存至全局内存中（该操作会覆盖既有的对象）
func _UGSet(prefix string, model xorm.ITable) *GlobalObject {
	return nil
}

// 将数据保存至会话内存中（该操作会覆盖既有的对象）
func _USSet(prefix string, model xorm.ITable) *SessionObject {
	return nil
}

// 内存拷贝（init-调用CTOR和Decode初始化）
func _UClone(model xorm.ITable, init bool) xorm.ITable {
	return nil
}

// 标记是否读写（若该数据已经被标记为读写，则无法修改为只读）
func _URW(model xorm.ITable, rw bool) {
	return
}

// 根据传入的cond（条件）判断（条件唯一且为主键）该数据（RedisPrefix）是否被列举过（全局内存）
// [record-标记]
// [clear-清除]
func _UGListed(prefix string, model xorm.ITable, record bool, clear bool, _cond ...*xorm.Condition) bool {
	return false
}

// 根据传入的cond（条件）判断（条件唯一且为主键）该数据（RedisPrefix）是否被列举过（会话内存）
// [record-标记]
// [clear-清除]
func _USListed(prefix string, model xorm.ITable, record bool, clear bool, _cond ...*xorm.Condition) bool {
	return false
}

// 根据指定的条件对对象进行过滤
func _UFilter(model xorm.ITable, conds []xorm.CondValue) bool {
	return false
}

// 根据指定操作符对字段进行比较
func _UComp(model xorm.ITable, cond xorm.CondValue) bool {
	return false
}

// 数据同步等待（一般用于GRead和GList的远端读取）
func _UGWait(source string, prefix string, model xorm.ITable) bool {
	return false
}

// 获取会话ID和线程ID
func _USTID() (int64, int64) {
	return 0, 0
}
