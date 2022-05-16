// 线程分布（共10个线程）
// [0 - 主线程，分配任务等，读写]
// [1:9 - 数据统计等，只读]
package app

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/test/hall/ctx"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/models/mmn"
)

func init() {
	mmn.RegPlayer().LCache(true).LDB(true).LRedis(true).LRW(true)
}

type HallServer struct {
	xserver.Server
}

func NewHallServer() *HallServer {
	this := &HallServer{}
	this.CTOR(this)
	return this
}
