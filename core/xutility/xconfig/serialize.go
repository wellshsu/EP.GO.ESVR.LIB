//go:binary-only-package
//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//
package xconfig

import (
	_ "encoding/json"
	_ "encoding/xml"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xfs"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
)

type parseFunc func(data []byte, v interface{}) error

func doParse(parse parseFunc, file string, out interface{}) error {
	return nil
}
func JsonToObj(file string, out interface{}) error {
	return nil
}
func XmlToObj(file string, out interface{}) error {
	return nil
}
