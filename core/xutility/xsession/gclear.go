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
	"github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xorm"
)

// 清除数据（会话内同步，远端异步）
// [条件（表达式形式）（首选）：orm.Parse("a > {0} && b == {1} && c contains {2}", 1, 2, "hello")，对象运算符：contains/startswith/endswith/isnull]
// [条件（对象形式）：orm.NewCondition().And("a__gt",1).And("b",2).And("c__contains","hello")]
// [ __gt-大于 ]
// [ __gte-大于等于 ]
// [ __lt-小于 ]
// [ __lte-小于等于 ]
// [ __in-特定（最多65534个KEY）]
// [ __exact-等于 ]
// [ __ne-不等于]
// [ __contains-包含 ]
// [ __startswith-以...起始 ]
// [ __endswith-以...结束 ]
// [ __isnull-是否为空 ]
func GClear(model xorm.ITable, _cond ...*xorm.Condition) {
	return
}
