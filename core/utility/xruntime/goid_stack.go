//go:binary-only-package
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
