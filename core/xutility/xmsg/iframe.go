//go:binary-only-package
package xmsg

type IFrame interface {
	GetRoute() *Route
	Reset()
	String() string
	ProtoMessage()
}
