package main

import (
	_ "esvr/test/shared/protocol"

	_ "esvr/core/utility/xlog"
	"esvr/core/utility/xrun"

	"esvr/core/server"
	"esvr/test/hall/app"
)

func main() {
	defer xrun.Caught(true)

	server.Start(app.NewHallServer())
}
