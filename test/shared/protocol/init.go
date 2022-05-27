package protocol

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xserver"
)

func init() {
	initmid()
	initrid()
	initcid()
}

func initmid() {
	_map := make(map[int]*xserver.MsgRoute)
	ttpe := reflect.TypeOf(MID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &xserver.MsgRoute{}
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
			route.ID = id
			_map[id] = route
		}
	}
	xserver.RegMsgRoute(_map)
}

func initrid() {
	_map := make(map[int]*xserver.RpcRoute)
	ttpe := reflect.TypeOf(RID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &xserver.RpcRoute{}
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
			route.ID = id
			_map[id] = route
		}
	}
	xserver.RegRpcRoute(_map)
}

func initcid() {
	imap := make(map[int]*xserver.CgiRoute)
	smap := make(map[string]*xserver.CgiRoute)
	ttpe := reflect.TypeOf(CID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		path := tfld.Tag.Get("path")
		if path == "" {
			fmt.Println("initcgi: empty path")
			continue
		}
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &xserver.CgiRoute{}
			route.Name = tfld.Name
			route.Path = path
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
			route.Method = strings.Split(tfld.Tag.Get("method"), ",")
			_id := tfld.Tag.Get("id")
			id, _ := strconv.Atoi(_id)
			route.ID = id
			imap[id] = route
			if smap[route.Path] != nil {
				fmt.Println(fmt.Sprintf("initcgi: dumplicated path %v", route.Path))
			}
			smap[route.Path] = route
		}
	}
	xserver.RegCgiRoute(imap, smap)
}
