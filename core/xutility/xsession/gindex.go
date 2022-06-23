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
	_ "sync"

	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
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
