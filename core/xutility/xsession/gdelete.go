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

// 删除数据（会话内同步，远端异步）
// [注意：该操作会维护数据同步锁]
// [注意：条件删除请使用GClear]
func GDelete(model xorm.ITable) {
	return
}
