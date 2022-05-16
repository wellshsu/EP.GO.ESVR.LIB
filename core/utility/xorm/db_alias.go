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
	"time"
)

// DriverType database driver constant int.
type DriverType int

// Enum the Database driver
const (
	// int enum type
	_          DriverType = iota // int enum type
	_          DriverType = iota // int enum type
	DRMySQL                      // mysql
	DRSqlite                     // sqlite
	DROracle                     // oracle
	DRPostgres                   // pgsql
	DRTiDB                       // TiDB
)

// database driver string.
type driver string

// get type constant int of current driver..
func (d driver) Type() DriverType {
	return 0
}

// get name of current driver
func (d driver) Name() string {
	return ""
}

// check driver iis implemented Driver interface or not.
// github.com/mattn/go-oci8
//https://github.com/rana/ora
// database alias cacher.
type _dbCache struct {
}

// add database alias with original name.
func (ac *_dbCache) add(name string, al *alias) (added bool) {
	return added
}

// get database alias if cached.
func (ac *_dbCache) get(name string) (al *alias, ok bool) {
	return al, ok
}

// get default alias.
func (ac *_dbCache) getDefault() (al *alias) {
	return al
}

type alias struct {
	Name         string
	Driver       DriverType
	DriverName   string
	DataSource   string
	MaxIdleConns int
	MaxOpenConns int
	DB           *sql.DB
	DbBaser      dbBaser
	TZ           *time.Location
	Engine       string
}

func detectTZ(al *alias) {
	return
}
func addAliasWthDB(aliasName, driverName string, db *sql.DB) (*alias, error) {
	return nil, nil
}

// AddAliasWthDB add a aliasName for the drivename
func AddAliasWthDB(aliasName, driverName string, db *sql.DB) error {
	return nil
}

// RegisterDataBase Setting the database connect params. Use the database driver self dataSource args.
func RegisterDataBase(aliasName, driverName, dataSource string, params ...int) error {
	return nil
}

// RegisterDriver Register a database driver use specify driver name, this can be definition the driver is which database type.
func RegisterDriver(driverName string, typ DriverType) error {
	return nil
}

// SetDataBaseTZ Change the database default used timezone
func SetDataBaseTZ(aliasName string, tz *time.Location) error {
	return nil
}

// SetMaxIdleConns Change the max idle conns for *sql.DB, use specify database alias name
func SetMaxIdleConns(aliasName string, maxIdleConns int) {
	return
}

// SetMaxOpenConns Change the max open conns for *sql.DB, use specify database alias name
func SetMaxOpenConns(aliasName string, maxOpenConns int) {
	return
}

// GetDB Get *sql.DB from registered database by db alias name.
// Use "default" as alias name if you not set.
func GetDB(aliasNames ...string) (*sql.DB, error) {
	return nil, nil
}
