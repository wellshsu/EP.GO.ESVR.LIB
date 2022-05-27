//go:binary-only-package
package xsession

import (
	_ "sync"
)

func _GetProfiler() *ProfilerRecord {
	return nil
}
func _PutProfiler(obj *ProfilerRecord) {
	return
}
