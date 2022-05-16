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
	"database/sql"
	_ "fmt"
	"io"
	"log"
	_ "strings"
	"time"
)

// Log implement the log.Logger
type Log struct {
	*log.Logger
}

//costomer log func
var LogFunc func(query map[string]interface{})

// NewLog set io.Writer to create a Logger.
func NewLog(out io.Writer) *Log {
	return nil
}
func debugLogQueies(alias *alias, operaton, query string, t time.Time, err error, args ...interface{}) {
	return
}

// statement query logger struct.
// if dev mode, use stmtQueryLog, or use stmtQuerier.
type stmtQueryLog struct {
}

func (d *stmtQueryLog) Close() error {
	return nil
}
func (d *stmtQueryLog) Exec(args ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (d *stmtQueryLog) Query(args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
func (d *stmtQueryLog) QueryRow(args ...interface{}) *sql.Row {
	return nil
}
func newStmtQueryLog(alias *alias, stmt stmtQuerier, query string) stmtQuerier {
	return nil
}

// database query logger struct.
// if dev mode, use dbQueryLog, or use dbQuerier.
type dbQueryLog struct {
}

func (d *dbQueryLog) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}
func (d *dbQueryLog) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, nil
}
func (d *dbQueryLog) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (d *dbQueryLog) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (d *dbQueryLog) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
func (d *dbQueryLog) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
func (d *dbQueryLog) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}
func (d *dbQueryLog) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return nil
}
func (d *dbQueryLog) Begin() (*sql.Tx, error) {
	return nil, nil
}
func (d *dbQueryLog) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return nil, nil
}
func (d *dbQueryLog) Commit() error {
	return nil
}
func (d *dbQueryLog) Rollback() error {
	return nil
}
func (d *dbQueryLog) SetDB(db dbQuerier) {
	return
}
func newDbQueryLog(alias *alias, db dbQuerier) dbQuerier {
	return nil
}
