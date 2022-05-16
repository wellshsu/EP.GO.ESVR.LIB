// 线程分布（共10个线程）
// [0 - 主线程，分配任务等，读写]
// [1:9 - 数据统计等，只读]
package app

import (
	"esvr/core/server"
	_ "esvr/test/hall/ctx"
	"esvr/test/shared/models/mmn"
)

func init() {
	mmn.RegPlayer().LCache(true).LDB(true).LRedis(true).LRW(true)
}

type HallServer struct {
	server.Server
}

func NewHallServer() *HallServer {
	this := &HallServer{}
	this.CTOR(this)
	return this
}
