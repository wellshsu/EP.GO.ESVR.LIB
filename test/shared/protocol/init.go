package protocol

import (
	"esvr/core/server"
	"reflect"
	"strconv"
	"strings"
)

func init() {
	rmap := make(map[int]*server.Route)
	ttpe := reflect.TypeOf(EID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &server.Route{}
			route.Name = tfld.Name
			route.Dst = strings.Split(_dst, ",")
			if _gid != "" {
				if strings.Contains(_gid, ":") {
					strs := strings.Split(_gid, ":")
					gid, _ := strconv.Atoi(strs[0])
					route.GoL = gid
					if len(strs) == 1 {
						route.GoR = 0
					} else {
						gid2, _ := strconv.Atoi(strs[1])
						route.GoR = gid2
					}
				} else {
					gid, _ := strconv.Atoi(_gid)
					route.GoL = gid
					route.GoR = 0
				}
			} else {
				route.GoL = 0
				route.GoR = 0
			}
			if _rw != "" {
				route.RW = _rw == "1"
			} else {
				route.RW = true
			}
			if _log != "" {
				route.Log = _log == "1"
			} else {
				route.Log = true
			}
			_id := tfld.Tag.Get("id")
			id, _ := strconv.Atoi(_id)
			rmap[id] = route
		}
	}
	server.RegRoute(rmap)
}
