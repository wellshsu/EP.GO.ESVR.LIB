//-- Auto generated by gopb --//
//--       DO NOT EDIT      --//
package mpb

import (
	"reflect"
	"strconv"
)

var MID _MID
var MIDIS map[int]string = make(map[int]string)
var MIDSI map[string]int = make(map[string]int)

func init() {
	enums := make([][]interface{}, 0)
	enums = append(enums, []interface{}{
		reflect.TypeOf(MID),
		reflect.ValueOf(&MID).Elem(),
		MIDIS,
		MIDSI,
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

type MIDEnum int

type _MID struct {
	GM_HEART_BEAT     MIDEnum `id:"0" name:"GM_HEART_BEAT" to:"gate"`                  // 心跳包
	GM_KICK_OFF       MIDEnum `id:"1" name:"GM_KICK_OFF" to:"client"`                  // 下线
	GM_LOGIN_REQUEST  MIDEnum `id:"2" name:"GM_LOGIN_REQUEST" to:"hall" go:"0" rw:"1"` // 请求设备登录
	GM_LOGIN_RESPONSE MIDEnum `id:"3" name:"GM_LOGIN_RESPONSE" to:"client"`            // 返回设备登录
}