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
	_ "os"
	_ "strings"
)

type dbIndex struct {
	Table string
	Name  string
	SQL   string
}

// create database drop sql.
func getDbDropSQL(al *alias) (sqls []string) {
	return sqls
}

// get database column type string.
func getColumnTyp(al *alias, fi *fieldInfo) (col string) {
	return col
}

// create alter sql string.
func getColumnAddQuery(al *alias, fi *fieldInfo) string {
	return ""
}

// create database creation string.
func getDbCreateSQL(al *alias) (sqls []string, tableIndexes map[string][]dbIndex) {
	return sqls, tableIndexes
}

// Get string value for the attribute "DEFAULT" for the CREATE, ALTER commands
func getColumnDefault(fi *fieldInfo) string {
	return ""
}
