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
	_ "errors"
	_ "fmt"
	"reflect"
	_ "strings"
)

// field info collection
type fields struct {
}

// add field info
func (f *fields) Add(fi *fieldInfo) (added bool) {
	return added
}

// get field info by name
func (f *fields) GetByName(name string) *fieldInfo {
	return nil
}

// get field info by column name
func (f *fields) GetByColumn(column string) *fieldInfo {
	return nil
}

// get field info by string, name is prior
func (f *fields) GetByAny(name string) (*fieldInfo, bool) {
	return nil, false
}

// create new field info collection
func newFields() *fields {
	return nil
}

// single field info
type fieldInfo struct {
	// table column fk and onetoone
	// whether has default tag
	// store the default value
	// if type equal to RelForeignKey, RelOneToOne, RelManyToMany then true
	// implement Fielder interface
}

// new field info
func newFieldInfo(mi *modelInfo, field reflect.Value, sf reflect.StructField, mName string) (fi *fieldInfo, err error) {
	return fi, err
}

type fieldRedis struct {
}

func newRedisInfo(mi *modelInfo, field reflect.Value, sf reflect.StructField, mName string) *fieldRedis {
	return nil
}
