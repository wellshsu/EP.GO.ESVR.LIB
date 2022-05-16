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
	_ "strconv"
	_ "strings"
)

// TiDBQueryBuilder is the SQL build
type TiDBQueryBuilder struct {
	Tokens []string
}

// Select will join the fields
func (qb *TiDBQueryBuilder) Select(fields ...string) QueryBuilder {
	return nil
}

// ForUpdate add the FOR UPDATE clause
func (qb *TiDBQueryBuilder) ForUpdate() QueryBuilder {
	return nil
}

// From join the tables
func (qb *TiDBQueryBuilder) From(tables ...string) QueryBuilder {
	return nil
}

// InnerJoin INNER JOIN the table
func (qb *TiDBQueryBuilder) InnerJoin(table string) QueryBuilder {
	return nil
}

// LeftJoin LEFT JOIN the table
func (qb *TiDBQueryBuilder) LeftJoin(table string) QueryBuilder {
	return nil
}

// RightJoin RIGHT JOIN the table
func (qb *TiDBQueryBuilder) RightJoin(table string) QueryBuilder {
	return nil
}

// On join with on cond
func (qb *TiDBQueryBuilder) On(cond string) QueryBuilder {
	return nil
}

// Where join the Where cond
func (qb *TiDBQueryBuilder) Where(cond string) QueryBuilder {
	return nil
}

// And join the and cond
func (qb *TiDBQueryBuilder) And(cond string) QueryBuilder {
	return nil
}

// Or join the or cond
func (qb *TiDBQueryBuilder) Or(cond string) QueryBuilder {
	return nil
}

// In join the IN (vals)
func (qb *TiDBQueryBuilder) In(vals ...string) QueryBuilder {
	return nil
}

// OrderBy join the Order by fields
func (qb *TiDBQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	return nil
}

// Asc join the asc
func (qb *TiDBQueryBuilder) Asc() QueryBuilder {
	return nil
}

// Desc join the desc
func (qb *TiDBQueryBuilder) Desc() QueryBuilder {
	return nil
}

// Limit join the limit num
func (qb *TiDBQueryBuilder) Limit(limit int) QueryBuilder {
	return nil
}

// Offset join the offset num
func (qb *TiDBQueryBuilder) Offset(offset int) QueryBuilder {
	return nil
}

// GroupBy join the Group by fields
func (qb *TiDBQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	return nil
}

// Having join the Having cond
func (qb *TiDBQueryBuilder) Having(cond string) QueryBuilder {
	return nil
}

// Update join the update table
func (qb *TiDBQueryBuilder) Update(tables ...string) QueryBuilder {
	return nil
}

// Set join the set kv
func (qb *TiDBQueryBuilder) Set(kv ...string) QueryBuilder {
	return nil
}

// Delete join the Delete tables
func (qb *TiDBQueryBuilder) Delete(tables ...string) QueryBuilder {
	return nil
}

// InsertInto join the insert SQL
func (qb *TiDBQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	return nil
}

// Values join the Values(vals)
func (qb *TiDBQueryBuilder) Values(vals ...string) QueryBuilder {
	return nil
}

// Subquery join the sub as alias
func (qb *TiDBQueryBuilder) Subquery(sub string, alias string) string {
	return ""
}

// String join all Tokens
func (qb *TiDBQueryBuilder) String() string {
	return ""
}
