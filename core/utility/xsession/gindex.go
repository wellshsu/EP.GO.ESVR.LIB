//go:binary-only-package
package xsession

import (
	_ "esvr/core/utility/xlog"
	"esvr/core/utility/xorm"
	_ "sync"
)

// 获取DB指定Column（列）的最大值，并且自增（指定delta值或+1），若未指定Column则自增主键（Column需为int类型）
// [ argcount1:column/delta]
// [ argcount2:column&delta]
func GIncrease(model xorm.ITable, columnAndDelta ...interface{}) int {
	return 0
}

// 获取DB指定Column（列）的最大值，若未指定Column则获取主键的最大值（Column需为int类型）
// 注意：异步写入可能会导致该值不同步
func GMaxIndex(model xorm.ITable, column ...string) int {
	return 0
}

// 获取DB指定Column（列）的最小值，若未指定Column则获取主键的最小值（Column需为int类型）
// 注意：异步写入可能会导致该值不同步
func GMinIndex(model xorm.ITable, column ...string) int {
	return 0
}
