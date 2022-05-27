//go:binary-only-package
package xserver

type Route struct {
	ID   int
	Name string
	GoL  int
	GoR  int
	RW   bool
	Log  bool
	Dst  []string
}
type MsgRoute struct {
	Route
}

var MSGROUTEMAP map[int]*MsgRoute // msg路由
func RegMsgRoute(_map map[int]*MsgRoute) {
	return
}

type RpcRoute struct {
	Route
}

var RPCROUTEMAP map[int]*RpcRoute // rpc路由
func RegRpcRoute(_map map[int]*RpcRoute) {
	return
}

type CgiRoute struct {
	Route
	Path   string
	Method []string
}

var ICGIROUTEMAP map[int]*CgiRoute    // cgi路由(id)
var SCGIROUTEMAP map[string]*CgiRoute // cgi路由(path)
func RegCgiRoute(imap map[int]*CgiRoute, smap map[string]*CgiRoute) {
	return
}
