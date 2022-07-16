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
	_ "fmt"
	_ "sync/atomic"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xrun"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xtime"
)

// 分配一个推送管道（多线程）
func allocPipe(tid int64) chan *PipeRecord {
	return nil
}
