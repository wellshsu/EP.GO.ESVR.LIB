//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//

package main

import (
	"github.com/wellshsu/EP.GO.ESVR.LIB/core/xserver"
	_ "github.com/wellshsu/EP.GO.ESVR.LIB/test/shared/proto"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xrun"

	"github.com/wellshsu/EP.GO.ESVR.LIB/test/public/cgi/app"
)

func main() {
	defer xrun.Caught(true)

	xserver.Start(app.NewCgiServer())
}
