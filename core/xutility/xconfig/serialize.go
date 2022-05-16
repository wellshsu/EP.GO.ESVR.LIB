//go:binary-only-package
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
