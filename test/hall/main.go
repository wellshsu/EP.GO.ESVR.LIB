package main

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/protocol"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xrun"

	"github.com/hsu2017/EP.GO.ESVR.LIB/test/hall/app"
)

func main() {
	defer xrun.Caught(true)

	xserver.Start(app.NewHallServer())
}
