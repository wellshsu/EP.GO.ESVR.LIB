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
)

// sqlite operators.
// sqlite column types.
// sqlite dbBaser.
type dbBaseSqlite struct {
	dbBase
}

// get sqlite operator.
func (d *dbBaseSqlite) OperatorSQL(operator string) string {
	return ""
}

// generate functioned sql for sqlite.
// only support DATE(text).
func (d *dbBaseSqlite) GenerateOperatorLeftCol(fi *fieldInfo, operator string, leftCol *string) {
	return
}

// unable updating joined record in sqlite.
func (d *dbBaseSqlite) SupportUpdateJoin() bool {
	return false
}

// max int in sqlite.
func (d *dbBaseSqlite) MaxLimit() uint64 {
	return 0
}

// get column types in sqlite.
func (d *dbBaseSqlite) DbTypes() map[string]string {
	return nil
}

// get show tables sql in sqlite.
func (d *dbBaseSqlite) ShowTablesQuery() string {
	return ""
}

// get columns in sqlite.
func (d *dbBaseSqlite) GetColumns(db dbQuerier, table string) (map[string][3]string, error) {
	return nil, nil
}

// get show columns sql in sqlite.
func (d *dbBaseSqlite) ShowColumnsQuery(table string) string {
	return ""
}

// check index exist in sqlite.
func (d *dbBaseSqlite) IndexExists(db dbQuerier, table string, name string) bool {
	return false
}

// create new sqlite dbBaser.
func newdbBaseSqlite() dbBaser {
	return nil
}
