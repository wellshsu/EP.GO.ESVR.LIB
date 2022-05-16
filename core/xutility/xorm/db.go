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
	"errors"
	_ "fmt"
	"reflect"
	_ "strings"
	"time"
)

var (
	// ErrMissPK missing pk error
	ErrMissPK = errors.New("missed pk value")
)

// "regex":       true,
// "iregex":      true,
// "year":        true,
// "month":       true,
// "day":         true,
// "week_day":    true,
// "search":      true,
// an instance of dbBaser interface/
type dbBase struct {
}

// check dbBase implements dbBaser interface.
// get struct columns values as interface slice.
func (d *dbBase) collectValues(mi *modelInfo, ind reflect.Value, cols []string, skipAuto bool, insert bool, names *[]string, tz *time.Location) (values []interface{}, autoFields []string, err error) {
	return values, autoFields, err
}

// get one field value in struct column as interface.
func (d *dbBase) collectFieldValue(mi *modelInfo, fi *fieldInfo, ind reflect.Value, insert bool, tz *time.Location) (interface{}, error) {
	return nil, nil
}

// create insert sql preparation statement object.
func (d *dbBase) PrepareInsert(q dbQuerier, mi *modelInfo) (stmtQuerier, string, error) {
	return nil, "", nil
}

// insert struct with prepared statement and given struct reflect value.
func (d *dbBase) InsertStmt(stmt stmtQuerier, mi *modelInfo, ind reflect.Value, tz *time.Location) (int64, error) {
	return 0, nil
}

// query sql ,read records and persist in dbBaser.
func (d *dbBase) Read(q dbQuerier, mi *modelInfo, ind reflect.Value, tz *time.Location, cols []string, isForUpdate bool) error {
	return nil
}

// execute insert sql dbQuerier with given struct reflect.Value.
func (d *dbBase) Insert(q dbQuerier, mi *modelInfo, ind reflect.Value, tz *time.Location) (int64, error) {
	return 0, nil
}

// multi-insert sql with given slice struct reflect.Value.
func (d *dbBase) InsertMulti(q dbQuerier, mi *modelInfo, sind reflect.Value, bulk int, tz *time.Location) (int64, error) {
	return 0, nil
}

// execute insert sql with given struct and given values.
// insert the given values, not the field values in struct.
func (d *dbBase) InsertValue(q dbQuerier, mi *modelInfo, isMulti bool, names []string, values []interface{}) (int64, error) {
	return 0, nil
}

// InsertOrUpdate a row
// If your primary key or unique column conflict will update
// If no will insert
func (d *dbBase) InsertOrUpdate(q dbQuerier, mi *modelInfo, ind reflect.Value, a *alias, args ...string) (int64, error) {
	return 0, nil
}

// execute update sql dbQuerier with given struct reflect.Value.
func (d *dbBase) Update(q dbQuerier, mi *modelInfo, ind reflect.Value, tz *time.Location, cols []string) (int64, error) {
	return 0, nil
}

// execute delete sql dbQuerier with given struct reflect.Value.
// delete index is pk.
func (d *dbBase) Delete(q dbQuerier, mi *modelInfo, ind reflect.Value, tz *time.Location, cols []string) (int64, error) {
	return 0, nil
}

// update table-related record by querySet.
// need querySet not struct reflect.Value to update related records.
func (d *dbBase) UpdateBatch(q dbQuerier, qs *querySet, mi *modelInfo, cond *Condition, params Params, tz *time.Location) (int64, error) {
	return 0, nil
}

// delete related records.
// do UpdateBanch or DeleteBanch by condition of tables' relationship.
func (d *dbBase) deleteRels(q dbQuerier, mi *modelInfo, args []interface{}, tz *time.Location) error {
	return nil
}

// delete table-related records.
func (d *dbBase) DeleteBatch(q dbQuerier, qs *querySet, mi *modelInfo, cond *Condition, tz *time.Location) (int64, error) {
	return 0, nil
}

// read related records.
func (d *dbBase) ReadBatch(q dbQuerier, qs *querySet, mi *modelInfo, cond *Condition, container interface{}, tz *time.Location, cols []string) (int64, error) {
	return 0, nil
}

// excute count sql and return count result int64.
func (d *dbBase) Count(q dbQuerier, qs *querySet, mi *modelInfo, cond *Condition, tz *time.Location) (cnt int64, err error) {
	return cnt, err
}

// generate sql with replacing operator string placeholders and replaced values.
func (d *dbBase) GenerateOperatorSQL(mi *modelInfo, fi *fieldInfo, operator string, args []interface{}, tz *time.Location) (string, []interface{}) {
	return "", []interface{}{}
}

// gernerate sql string with inner function, such as UPPER(text).
func (d *dbBase) GenerateOperatorLeftCol(*fieldInfo, string, *string) {
	return
}

// set values to struct column.
func (d *dbBase) setColsValues(mi *modelInfo, ind *reflect.Value, cols []string, values []interface{}, tz *time.Location) {
	return
}

// convert value from database result to value following in field type.
func (d *dbBase) convertValueFromDB(fi *fieldInfo, val interface{}, tz *time.Location) (interface{}, error) {
	return nil, nil
}

// set one value to struct column field.
func (d *dbBase) setFieldValue(fi *fieldInfo, value interface{}, field reflect.Value) (interface{}, error) {
	return nil, nil
}

// query sql, read values , save to *[]ParamList.
func (d *dbBase) ReadValues(q dbQuerier, qs *querySet, mi *modelInfo, cond *Condition, exprs []string, container interface{}, tz *time.Location) (int64, error) {
	return 0, nil
}
func (d *dbBase) RowsTo(dbQuerier, *querySet, *modelInfo, *Condition, interface{}, string, string, *time.Location) (int64, error) {
	return 0, nil
}

// flag of update joined record.
func (d *dbBase) SupportUpdateJoin() bool {
	return false
}
func (d *dbBase) MaxLimit() uint64 {
	return 0
}

// return quote.
func (d *dbBase) TableQuote() string {
	return ""
}

// replace value placeholder in parametered sql string.
func (d *dbBase) ReplaceMarks(query *string) {
	return
}

// flag of RETURNING sql.
func (d *dbBase) HasReturningID(*modelInfo, *string) bool {
	return false
}

// sync auto key
func (d *dbBase) setval(db dbQuerier, mi *modelInfo, autoFields []string) error {
	return nil
}

// convert time from db.
func (d *dbBase) TimeFromDB(t *time.Time, tz *time.Location) {
	return
}

// convert time to db.
func (d *dbBase) TimeToDB(t *time.Time, tz *time.Location) {
	return
}

// get database types.
func (d *dbBase) DbTypes() map[string]string {
	return nil
}

// gt all tables.
func (d *dbBase) GetTables(db dbQuerier) (map[string]bool, error) {
	return nil, nil
}

// get all cloumns in table.
func (d *dbBase) GetColumns(db dbQuerier, table string) (map[string][3]string, error) {
	return nil, nil
}

// not implement.
func (d *dbBase) OperatorSQL(operator string) string {
	return ""
}

// not implement.
func (d *dbBase) ShowTablesQuery() string {
	return ""
}

// not implement.
func (d *dbBase) ShowColumnsQuery(table string) string {
	return ""
}

// not implement.
func (d *dbBase) IndexExists(dbQuerier, string, string) bool {
	return false
}
