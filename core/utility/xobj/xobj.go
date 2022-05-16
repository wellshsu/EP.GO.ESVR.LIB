//go:binary-only-package
// 完善面向对象逻辑
package xobj

// 基类
type OBJECT struct {
	CHILD interface{} `orm:"-" json:"-"` // 实例指针对象
}

// 构造函数
func (this *OBJECT) CTOR(CHILD interface{}) {
	return
}
