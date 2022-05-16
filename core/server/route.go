//go:binary-only-package
package server

type Route struct {
	Name string
	GoL  int
	GoR  int
	RW   bool
	Log  bool
	Dst  []string
}

var ROUTEMAP map[int]*Route

func RegRoute(rmap map[int]*Route) {
	return
}
