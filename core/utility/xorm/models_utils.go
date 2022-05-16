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
	_ "database/sql"
	_ "fmt"
	"reflect"
	_ "strings"
	_ "time"
)

// 1 is attr
// 2 is tag
// get reflect.Type name with package path.
func getFullName(typ reflect.Type) string {
	return ""
}

// getTableName get struct table name.
// If the struct implement the TableName, then get the result as tablename
// else use the struct name which will apply snakeString.
func getTableName(val reflect.Value) string {
	return ""
}

// get table engine, mysiam or innodb.
func getTableEngine(val reflect.Value) string {
	return ""
}

// get table index from method.
func getTableIndex(val reflect.Value) [][]string {
	return [][]string{}
}

// get table unique from method
func getTableUnique(val reflect.Value) [][]string {
	return [][]string{}
}

// get snaked column name
func getColumnName(ft int, addrField reflect.Value, sf reflect.StructField, col string) string {
	return ""
}

// return field type as type constant from reflect.Value
func getFieldType(val reflect.Value) (ft int, err error) {
	return ft, err
}

// parse struct tag string
func parseStructTag(data string) (attrs map[string]bool, tags map[string]string) {
	return attrs, tags
}
