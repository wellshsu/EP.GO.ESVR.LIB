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
	_ "strconv"
)

// postgresql operators.
// postgresql column field types.
// postgresql dbBaser.
type dbBasePostgres struct {
	dbBase
}

// get postgresql operator.
func (d *dbBasePostgres) OperatorSQL(operator string) string {
	return ""
}

// generate functioned sql string, such as contains(text).
func (d *dbBasePostgres) GenerateOperatorLeftCol(fi *fieldInfo, operator string, leftCol *string) {
	return
}

// postgresql unsupports updating joined record.
func (d *dbBasePostgres) SupportUpdateJoin() bool {
	return false
}
func (d *dbBasePostgres) MaxLimit() uint64 {
	return 0
}

// postgresql quote is ".
func (d *dbBasePostgres) TableQuote() string {
	return ""
}

// postgresql value placeholder is $n.
// replace default ? to $n.
func (d *dbBasePostgres) ReplaceMarks(query *string) {
	return
}

// make returning sql support for postgresql.
func (d *dbBasePostgres) HasReturningID(mi *modelInfo, query *string) bool {
	return false
}

// sync auto key
func (d *dbBasePostgres) setval(db dbQuerier, mi *modelInfo, autoFields []string) error {
	return nil
}

// show table sql for postgresql.
func (d *dbBasePostgres) ShowTablesQuery() string {
	return ""
}

// show table columns sql for postgresql.
func (d *dbBasePostgres) ShowColumnsQuery(table string) string {
	return ""
}

// get column types of postgresql.
func (d *dbBasePostgres) DbTypes() map[string]string {
	return nil
}

// check index exist in postgresql.
func (d *dbBasePostgres) IndexExists(db dbQuerier, table string, name string) bool {
	return false
}

// create new postgresql dbBaser.
func newdbBasePostgres() dbBaser {
	return nil
}
