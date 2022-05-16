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
	_ "unsafe"
)

// ExprSep define the expression separation
const (
	ExprSep = "__"
)

type condValue struct {
}

// Condition struct.
// work for WHERE conditions.
type Condition struct {
}

// NewCondition return new condition struct
func NewCondition() *Condition {
	return nil
}

// Raw add raw sql to condition
func (c Condition) Raw(expr string, sql string) *Condition {
	return nil
}

// And add expression to condition
func (c Condition) And(expr string, args ...interface{}) *Condition {
	return nil
}

// AndNot add NOT expression to condition
func (c Condition) AndNot(expr string, args ...interface{}) *Condition {
	return nil
}

// AndCond combine a condition to current condition
func (c *Condition) AndCond(cond *Condition) *Condition {
	return nil
}

// AndNotCond combine a AND NOT condition to current condition
func (c *Condition) AndNotCond(cond *Condition) *Condition {
	return nil
}

// Or add OR expression to condition
func (c Condition) Or(expr string, args ...interface{}) *Condition {
	return nil
}

// OrNot add OR NOT expression to condition
func (c Condition) OrNot(expr string, args ...interface{}) *Condition {
	return nil
}

// OrCond combine a OR condition to current condition
func (c *Condition) OrCond(cond *Condition) *Condition {
	return nil
}

// OrNotCond combine a OR NOT condition to current condition
func (c *Condition) OrNotCond(cond *Condition) *Condition {
	return nil
}

// IsEmpty check the condition arguments are empty or not.
func (c *Condition) IsEmpty() bool {
	return false
}

// clone clone a condition
func (c Condition) clone() *Condition {
	return nil
}

//================== Custom20200916 ==================//
type CondLink struct {
	Params []CondValue `json:"Params,omitempty"`
}
type CondValue struct {
	Exprs  []string      `json:"Exprs,omitempty"`
	Args   []interface{} `json:"Args,omitempty"`
	Cond   *CondLink     `json:"Cond,omitempty"`
	IsOr   bool
	IsNot  bool
	IsCond bool
	IsRaw  bool   `json:"-"`
	Sql    string `json:"Sql,omitempty"`
}

func ReadableConds(cond *Condition) []CondValue {
	return []CondValue{}
}

//================== Custom20200916 ==================//
//================== Custom20210223 ==================//
func init() {
	return
}
func Parse(expression string, params ...interface{}) *Condition {
	return nil
}
func getLRIndex(expression string) map[int]int {
	return nil
}
func doParse(expression string, lrMap map[int]int, left int, right int, lianStr string, feiStr string, cond *Condition, params ...interface{}) *Condition {
	return nil
}
func parseArr(expression string, _strs []string, lianStr string, cond *Condition, params ...interface{}) (string, string, *Condition) {
	return "", "", nil
}
func subStr(str string, from int, to int) string {
	return ""
}
func toInt(str string) int {
	return 0
}

//================== Custom20210223 ==================//
