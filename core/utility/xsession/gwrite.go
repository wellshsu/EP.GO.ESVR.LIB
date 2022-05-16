//go:binary-only-package
package xsession

import (
	_ "esvr/core/utility/xlog"
	"esvr/core/utility/xorm"
)

// 写入数据（会话内同步，远端异步）
func GWrite(model xorm.ITable) {
	return
}
