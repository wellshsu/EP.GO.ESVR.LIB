//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//

package app

import (
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/test/private/hall/ctx"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/model/mmn"
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
