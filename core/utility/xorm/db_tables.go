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
	"time"
)

// table info struct.
type dbTable struct {
}

// tables collection struct, contains some tables.
type dbTables struct {
}

// set table info to collection.
// if not exist, create new.
func (t *dbTables) set(names []string, mi *modelInfo, fi *fieldInfo, inner bool) *dbTable {
	return nil
}

// add table info to collection.
func (t *dbTables) add(names []string, mi *modelInfo, fi *fieldInfo, inner bool) (*dbTable, bool) {
	return nil, false
}

// get table info in collection.
func (t *dbTables) get(name string) (*dbTable, bool) {
	return nil, false
}

// get related fields info in recursive depth loop.
// loop once, depth decreases one.
func (t *dbTables) loopDepth(depth int, prefix string, fi *fieldInfo, related []string) []string {
	return []string{}
}

// parse related fields.
func (t *dbTables) parseRelated(rels []string, depth int) {
	return
}

// generate join string.
func (t *dbTables) getJoinSQL() (join string) {
	return join
}

// parse orm model struct field tag expression.
func (t *dbTables) parseExprs(mi *modelInfo, exprs []string) (index, name string, info *fieldInfo, success bool) {
	return index, name, info, success
}

// generate condition sql.
func (t *dbTables) getCondSQL(cond *Condition, sub bool, tz *time.Location) (where string, params []interface{}) {
	return where, params
}

// generate group sql.
func (t *dbTables) getGroupSQL(groups []string) (groupSQL string) {
	return groupSQL
}

// generate order sql.
func (t *dbTables) getOrderSQL(orders []string) (orderSQL string) {
	return orderSQL
}

// generate limit sql.
func (t *dbTables) getLimitSQL(mi *modelInfo, offset int64, limit int64) (limits string) {
	return limits
}

// crete new tables collection.
func newDbTables(mi *modelInfo, base dbBaser) *dbTables {
	return nil
}
