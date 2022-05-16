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
	_ "strings"
)

// oracle operators.
// oracle column field types.
// oracle dbBaser
type dbBaseOracle struct {
	dbBase
}

// create oracle dbBaser.
func newdbBaseOracle() dbBaser {
	return nil
}

// OperatorSQL get oracle operator.
func (d *dbBaseOracle) OperatorSQL(operator string) string {
	return ""
}

// DbTypes get oracle table field types.
func (d *dbBaseOracle) DbTypes() map[string]string {
	return nil
}

//ShowTablesQuery show all the tables in database
func (d *dbBaseOracle) ShowTablesQuery() string {
	return ""
}

// Oracle
func (d *dbBaseOracle) ShowColumnsQuery(table string) string {
	return ""
}

// check index is exist
func (d *dbBaseOracle) IndexExists(db dbQuerier, table string, name string) bool {
	return false
}

// execute insert sql with given struct and given values.
// insert the given values, not the field values in struct.
func (d *dbBaseOracle) InsertValue(q dbQuerier, mi *modelInfo, isMulti bool, names []string, values []interface{}) (int64, error) {
	return 0, nil
}
