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
	_ "strings"
)

// CommaSpace is the separation
const CommaSpace = ", "

// MySQLQueryBuilder is the SQL build
type MySQLQueryBuilder struct {
	Tokens []string
}

// Select will join the fields
func (qb *MySQLQueryBuilder) Select(fields ...string) QueryBuilder {
	return nil
}

// ForUpdate add the FOR UPDATE clause
func (qb *MySQLQueryBuilder) ForUpdate() QueryBuilder {
	return nil
}

// From join the tables
func (qb *MySQLQueryBuilder) From(tables ...string) QueryBuilder {
	return nil
}

// InnerJoin INNER JOIN the table
func (qb *MySQLQueryBuilder) InnerJoin(table string) QueryBuilder {
	return nil
}

// LeftJoin LEFT JOIN the table
func (qb *MySQLQueryBuilder) LeftJoin(table string) QueryBuilder {
	return nil
}

// RightJoin RIGHT JOIN the table
func (qb *MySQLQueryBuilder) RightJoin(table string) QueryBuilder {
	return nil
}

// On join with on cond
func (qb *MySQLQueryBuilder) On(cond string) QueryBuilder {
	return nil
}

// Where join the Where cond
func (qb *MySQLQueryBuilder) Where(cond string) QueryBuilder {
	return nil
}

// And join the and cond
func (qb *MySQLQueryBuilder) And(cond string) QueryBuilder {
	return nil
}

// Or join the or cond
func (qb *MySQLQueryBuilder) Or(cond string) QueryBuilder {
	return nil
}

// In join the IN (vals)
func (qb *MySQLQueryBuilder) In(vals ...string) QueryBuilder {
	return nil
}

// OrderBy join the Order by fields
func (qb *MySQLQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	return nil
}

// Asc join the asc
func (qb *MySQLQueryBuilder) Asc() QueryBuilder {
	return nil
}

// Desc join the desc
func (qb *MySQLQueryBuilder) Desc() QueryBuilder {
	return nil
}

// Limit join the limit num
func (qb *MySQLQueryBuilder) Limit(limit int) QueryBuilder {
	return nil
}

// Offset join the offset num
func (qb *MySQLQueryBuilder) Offset(offset int) QueryBuilder {
	return nil
}

// GroupBy join the Group by fields
func (qb *MySQLQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	return nil
}

// Having join the Having cond
func (qb *MySQLQueryBuilder) Having(cond string) QueryBuilder {
	return nil
}

// Update join the update table
func (qb *MySQLQueryBuilder) Update(tables ...string) QueryBuilder {
	return nil
}

// Set join the set kv
func (qb *MySQLQueryBuilder) Set(kv ...string) QueryBuilder {
	return nil
}

// Delete join the Delete tables
func (qb *MySQLQueryBuilder) Delete(tables ...string) QueryBuilder {
	return nil
}

// InsertInto join the insert SQL
func (qb *MySQLQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	return nil
}

// Values join the Values(vals)
func (qb *MySQLQueryBuilder) Values(vals ...string) QueryBuilder {
	return nil
}

// Subquery join the sub as alias
func (qb *MySQLQueryBuilder) Subquery(sub string, alias string) string {
	return ""
}

// String join all Tokens
func (qb *MySQLQueryBuilder) String() string {
	return ""
}
