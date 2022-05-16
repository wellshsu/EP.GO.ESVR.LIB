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
	_ "flag"
	_ "fmt"
	_ "os"
	_ "strings"
)

type commander interface {
	Parse([]string)
	Run() error
}

// print help.
func printHelp(errs ...string) {
	return
}

// RunCommand listen for orm command and then run it if command arguments passed.
func RunCommand() {
	return
}

// sync database struct command interface.
type commandSyncDb struct {
}

// parse orm command line arguments.
func (d *commandSyncDb) Parse(args []string) {
	return
}

// run orm line command.
func (d *commandSyncDb) Run() error {
	return nil
}

// database creation commander interface implement.
type commandSQLAll struct {
}

// parse orm command line arguments.
func (d *commandSQLAll) Parse(args []string) {
	return
}

// run orm line command.
func (d *commandSQLAll) Run() error {
	return nil
}
func init() {
	return
}

// RunSyncdb run syncdb command line.
// name means table's alias name. default is "default".
// force means run next sql if the current is error.
// verbose means show all info when running command or not.
func RunSyncdb(name string, force bool, verbose bool) error {
	return nil
}
