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
package xruntime

import (
	_ "runtime"
)

func get_stack(offset, amount int) (stack []uintptr, next_offset int) {
	return stack, next_offset
}
func read_stack() (tag uint, ok bool) {
	return tag, ok
}
func goid_stack() int64 {
	return 0
}
