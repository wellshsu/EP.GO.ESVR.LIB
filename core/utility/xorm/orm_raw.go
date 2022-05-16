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
	"database/sql"
	_ "fmt"
	"reflect"
	_ "time"
)

// raw sql string prepared statement
type rawPrepare struct {
}

func (o *rawPrepare) Exec(args ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (o *rawPrepare) Close() error {
	return nil
}
func newRawPreparer(rs *rawSet) (RawPreparer, error) {
	return nil, nil
}

// raw query seter
type rawSet struct {
}

// set args for every query
func (o rawSet) SetArgs(args ...interface{}) RawSeter {
	return nil
}

// execute raw sql and return sql.Result
func (o *rawSet) Exec() (sql.Result, error) {
	return nil, nil
}

// set field value to row container
func (o *rawSet) setFieldValue(ind reflect.Value, value interface{}) {
	return
}

// set field value in loop for slice container
func (o *rawSet) loopSetRefs(refs []interface{}, sInds []reflect.Value, nIndsPtr *[]reflect.Value, eTyps []reflect.Type, init bool) {
	return
}

// query data and map to container
func (o *rawSet) QueryRow(containers ...interface{}) error {
	return nil
}

// query data rows and map to container
func (o *rawSet) QueryRows(containers ...interface{}) (int64, error) {
	return 0, nil
}
func (o *rawSet) readValues(container interface{}, needCols []string) (int64, error) {
	return 0, nil
}
func (o *rawSet) queryRowsTo(container interface{}, keyCol, valueCol string) (int64, error) {
	return 0, nil
}

// query data to []map[string]interface
func (o *rawSet) Values(container *[]Params, cols ...string) (int64, error) {
	return 0, nil
}

// query data to [][]interface
func (o *rawSet) ValuesList(container *[]ParamsList, cols ...string) (int64, error) {
	return 0, nil
}

// query data to []interface
func (o *rawSet) ValuesFlat(container *ParamsList, cols ...string) (int64, error) {
	return 0, nil
}

// query all rows into map[string]interface with specify key and value column name.
// keyCol = "name", valueCol = "value"
// table data
// name  | value
// total | 100
// found | 200
// to map[string]interface{}{
// 	"total": 100,
// 	"found": 200,
// }
func (o *rawSet) RowsToMap(result *Params, keyCol, valueCol string) (int64, error) {
	return 0, nil
}

// query all rows into struct with specify key and value column name.
// keyCol = "name", valueCol = "value"
// table data
// name  | value
// total | 100
// found | 200
// to struct {
// 	Total int
// 	Found int
// }
func (o *rawSet) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	return 0, nil
}

// return prepared raw statement for used in times.
func (o *rawSet) Prepare() (RawPreparer, error) {
	return nil, nil
}
func newRawSet(orm *orm, query string, args []interface{}) RawSeter {
	return nil
}
