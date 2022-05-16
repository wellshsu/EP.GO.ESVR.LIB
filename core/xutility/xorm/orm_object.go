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
	_ "reflect"
)

// an insert queryer struct
type insertSet struct {
}

// insert model ignore it's registered or not.
func (o *insertSet) Insert(md interface{}) (int64, error) {
	return 0, nil
}

// close insert queryer statement
func (o *insertSet) Close() error {
	return nil
}

// create new insert queryer.
func newInsertSet(orm *orm, mi *modelInfo) (Inserter, error) {
	return nil, nil
}
