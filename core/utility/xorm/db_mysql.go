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
	_ "fmt"
	"reflect"
	_ "strings"
)

// mysql operators.
// "regex":       "REGEXP BINARY ?",
// "iregex":      "REGEXP ?",
// mysql column field types.
// mysql dbBaser implementation.
type dbBaseMysql struct {
	dbBase
}

// get mysql operator.
func (d *dbBaseMysql) OperatorSQL(operator string) string {
	return ""
}

// get mysql table field types.
func (d *dbBaseMysql) DbTypes() map[string]string {
	return nil
}

// show table sql for mysql.
func (d *dbBaseMysql) ShowTablesQuery() string {
	return ""
}

// show columns sql of table for mysql.
func (d *dbBaseMysql) ShowColumnsQuery(table string) string {
	return ""
}

// execute sql to check index exist.
func (d *dbBaseMysql) IndexExists(db dbQuerier, table string, name string) bool {
	return false
}

// InsertOrUpdate a row
// If your primary key or unique column conflict will update
// If no will insert
// Add "`" for mysql sql building
func (d *dbBaseMysql) InsertOrUpdate(q dbQuerier, mi *modelInfo, ind reflect.Value, a *alias, args ...string) (int64, error) {
	return 0, nil
}

// create new mysql dbBaser.
func newdbBaseMysql() dbBaser {
	return nil
}
