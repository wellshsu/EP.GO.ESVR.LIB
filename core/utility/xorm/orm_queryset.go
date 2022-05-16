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
	"context"
	_ "fmt"
)

type colValue struct {
}
type operator int

// define Col operations
const (
	ColAdd operator = iota
	ColMinus
	ColMultiply
	ColExcept
)

// ColValue do the field raw changes. e.g Nums = Nums + 10. usage:
// 	Params{
// 		"Nums": ColValue(Col_Add, 10),
// 	}
func ColValue(opt operator, value interface{}) interface{} {
	return nil
}

// real query struct
type querySet struct {
}

// add condition expression to QuerySeter.
func (o querySet) Filter(expr string, args ...interface{}) QuerySeter {
	return nil
}

// add raw sql to querySeter.
func (o querySet) FilterRaw(expr string, sql string) QuerySeter {
	return nil
}

// add NOT condition to querySeter.
func (o querySet) Exclude(expr string, args ...interface{}) QuerySeter {
	return nil
}

// set offset number
func (o *querySet) setOffset(num interface{}) {
	return
}

// add LIMIT value.
// args[0] means offset, e.g. LIMIT num,offset.
func (o querySet) Limit(limit interface{}, args ...interface{}) QuerySeter {
	return nil
}

// add OFFSET value
func (o querySet) Offset(offset interface{}) QuerySeter {
	return nil
}

// add GROUP expression
func (o querySet) GroupBy(exprs ...string) QuerySeter {
	return nil
}

// add ORDER expression.
// "column" means ASC, "-column" means DESC.
func (o querySet) OrderBy(exprs ...string) QuerySeter {
	return nil
}

// add DISTINCT to SELECT
func (o querySet) Distinct() QuerySeter {
	return nil
}

// add FOR UPDATE to SELECT
func (o querySet) ForUpdate() QuerySeter {
	return nil
}

// set relation model to query together.
// it will query relation models and assign to parent model.
func (o querySet) RelatedSel(params ...interface{}) QuerySeter {
	return nil
}

// set condition to QuerySeter.
func (o querySet) SetCond(cond *Condition) QuerySeter {
	return nil
}

// get condition from QuerySeter
func (o querySet) GetCond() *Condition {
	return nil
}

// return QuerySeter execution result number
func (o *querySet) Count() (int64, error) {
	return 0, nil
}

// check result empty or not after QuerySeter executed
func (o *querySet) Exist() bool {
	return false
}

// execute update with parameters
func (o *querySet) Update(values Params) (int64, error) {
	return 0, nil
}

// execute delete
func (o *querySet) Delete() (int64, error) {
	return 0, nil
}

// return a insert queryer.
// it can be used in times.
// example:
// 	i,err := sq.PrepareInsert()
// 	i.Add(&user1{},&user2{})
func (o *querySet) PrepareInsert() (Inserter, error) {
	return nil, nil
}

// query all data and map to containers.
// cols means the columns when querying.
func (o *querySet) All(container interface{}, cols ...string) (int64, error) {
	return 0, nil
}

// query one row data and map to containers.
// cols means the columns when querying.
func (o *querySet) One(container interface{}, cols ...string) error {
	return nil
}

// query all data and map to []map[string]interface.
// expres means condition expression.
// it converts data to []map[column]value.
func (o *querySet) Values(results *[]Params, exprs ...string) (int64, error) {
	return 0, nil
}

// query all data and map to [][]interface
// it converts data to [][column_index]value
func (o *querySet) ValuesList(results *[]ParamsList, exprs ...string) (int64, error) {
	return 0, nil
}

// query all data and map to []interface.
// it's designed for one row record set, auto change to []value, not [][column]value.
func (o *querySet) ValuesFlat(result *ParamsList, expr string) (int64, error) {
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
func (o *querySet) RowsToMap(result *Params, keyCol, valueCol string) (int64, error) {
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
func (o *querySet) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	return 0, nil
}

// set context to QuerySeter.
func (o querySet) WithContext(ctx context.Context) QuerySeter {
	return nil
}

// create new QuerySeter.
func newQuerySet(orm *orm, mi *modelInfo) QuerySeter {
	return nil
}
