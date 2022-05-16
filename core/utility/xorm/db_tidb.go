//go:binary-only-package
// Copyright 2015 TiDB Author. All Rights Reserved.
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
)

// mysql dbBaser implementation.
type dbBaseTidb struct {
	dbBase
}

// get mysql operator.
func (d *dbBaseTidb) OperatorSQL(operator string) string {
	return ""
}

// get mysql table field types.
func (d *dbBaseTidb) DbTypes() map[string]string {
	return nil
}

// show table sql for mysql.
func (d *dbBaseTidb) ShowTablesQuery() string {
	return ""
}

// show columns sql of table for mysql.
func (d *dbBaseTidb) ShowColumnsQuery(table string) string {
	return ""
}

// execute sql to check index exist.
func (d *dbBaseTidb) IndexExists(db dbQuerier, table string, name string) bool {
	return false
}

// create new mysql dbBaser.
func newdbBaseTidb() dbBaser {
	return nil
}
