//go:binary-only-package
package xsession

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
)

// 读取数据
// [优先级: 会话内存 > 全局缓存(LCache == true) > Redis(LRedis == true) > DB(LDB == true)]
// [注意：远端读取（LCache == false、该会话首次读取、条件查找），会触发数据同步锁，可能会严重影响效率（等待GDelete操作完成），可预加载（GList）该模型优化之]
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
func GRead(model xorm.ITable, _cond ...*xorm.Condition) xorm.ITable {
	return nil
}
