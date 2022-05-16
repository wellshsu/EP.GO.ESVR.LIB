//go:binary-only-package
package xsession

import (
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
)

// 删除数据（会话内同步，远端异步）
// [注意：该操作会维护数据同步锁]
// [注意：条件删除请使用GClear]
func GDelete(model xorm.ITable) {
	return
}
