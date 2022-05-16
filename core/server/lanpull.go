//go:binary-only-package
package server

import (
	_ "esvr/core/utility/xlog"
	_ "esvr/core/utility/xmsg"
	_ "esvr/core/utility/xrun"

	_ "github.com/golang/protobuf/proto"
)

func pullMsg(goNum int) {
	return
}
