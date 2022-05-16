//go:binary-only-package
// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package xorm

import (
	"reflect"
	"sync"
	_ "unsafe"
)

// model info collection
type _modelCache struct {
	sync.RWMutex // only used outsite for bootStrap
}

// get all model info
func (mc *_modelCache) all() map[string]*modelInfo {
	return nil
}

// get ordered model info
func (mc *_modelCache) allOrdered() []*modelInfo {
	return []*modelInfo{}
}

// get model info by table name
func (mc *_modelCache) get(table string) (mi *modelInfo, ok bool) {
	return mi, ok
}

// get model info by full name
func (mc *_modelCache) getByFullName(name string) (mi *modelInfo, ok bool) {
	return mi, ok
}

// set model info to collection
func (mc *_modelCache) set(table string, mi *modelInfo) *modelInfo {
	return nil
}

// clean all model info.
func (mc *_modelCache) clean() {
	return
}

// ResetModelCache Clean model cache. Then you can re-RegisterModel.
// Common use this api for test case.
func ResetModelCache() {
	return
}

//================== Custom20200910 ==================//
type ModelInfo struct {
	Pkg       string
	Name      string
	FullName  string
	Table     string
	Model     interface{}
	Fields    *Fields
	Manual    bool
	AddrField reflect.Value //store the original struct value
	Uniques   []string
	IsThrough bool
	LCache    bool
	LRedis    bool
	LDB       bool
	LRW       bool
	LSync     int
}
type Fields struct {
	Pk            *FieldInfo
	Columns       map[string]*FieldInfo
	Fields        map[string]*FieldInfo
	FieldsLow     map[string]*FieldInfo
	FieldsByType  map[int][]*FieldInfo
	FieldsRel     []*FieldInfo
	FieldsReverse []*FieldInfo
	FieldsDB      []*FieldInfo
	Rels          []*FieldInfo
	Orders        []string
	Dbcols        []string
	FieldsRedis   []*FieldRedis
}

// single field info
type FieldInfo struct {
	Mi                  *ModelInfo
	FieldIndex          []int
	FieldType           int
	Dbcol               bool // table column fk and onetoone
	InModel             bool
	Name                string
	FullName            string
	Column              string
	AddrValue           reflect.Value
	Sf                  reflect.StructField
	Auto                bool
	Pk                  bool
	Ok                  bool
	Null                bool
	Index               bool
	Unique              bool
	ColDefault          bool  // whether has default tag
	Initial             StrTo // store the default value
	Size                int
	ToText              bool
	AutoNow             bool
	AutoNowAdd          bool
	Rel                 bool // if type equal to RelForeignKey, RelOneToOne, RelManyToMany then true
	Reverse             bool
	ReverseField        string
	ReverseFieldInfo    *FieldInfo
	ReverseFieldInfoTwo *FieldInfo
	ReverseFieldInfoM2M *FieldInfo
	RelTable            string
	RelThrough          string
	RelThroughModelInfo *ModelInfo
	RelModelInfo        *ModelInfo
	Digits              int
	Decimals            int
	IsFielder           bool // implement Fielder interface
	OnDelete            string
	Description         string
}
type FieldRedis struct {
	Name   string
	Column string
	Kind   reflect.Kind
}

func GetModelInfo(name string) (mi *ModelInfo, ok bool) {
	return mi, ok
}

//================== Custom20200910 ==================//
