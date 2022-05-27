//-- Auto generated by gproto --//
//--        DO NOT EDIT       --//
package protocol

import (
	"reflect"
	"strconv"
)

var CID _CID
var CIDIS map[int]string = make(map[int]string)
var CIDSI map[string]int = make(map[string]int)

func init() {
	enums := make([][]interface{}, 0)
	enums = append(enums, []interface{}{
		reflect.TypeOf(CID),
		reflect.ValueOf(&CID).Elem(),
		CIDIS,
		CIDSI,
	})
	for _, enum := range enums {
		ttpe := enum[0].(reflect.Type)
		vtpe := enum[1].(reflect.Value)
		len := ttpe.NumField()
		for i := 0; i < len; i++ {
			tfld := ttpe.Field(i)
			vfld := vtpe.Field(i)
			sid := tfld.Tag.Get("id")
			id, _ := strconv.Atoi(sid)
			vfld.SetInt(int64(id))
			name := tfld.Tag.Get("name")
			enum[2].(map[int]string)[id] = name
			enum[3].(map[string]int)[name] = id
		}
	}
}

type CIDEnum int

type _CID struct {
	HELLO_WORLD CIDEnum `id:"0" name:"HELLO_WORLD" path:"/hello/world" to:"hall" go:"0" rw:"1" method:"GET,POST" timeout:"-1"` //
}
