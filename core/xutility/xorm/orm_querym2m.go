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

import "reflect"

// model to model struct
type queryM2M struct {
}

// add models to origin models when creating queryM2M.
// example:
// 	m2m := orm.QueryM2M(post,"Tag")
// 	m2m.Add(&Tag1{},&Tag2{})
//  for _,tag := range post.Tags{}
//
// make sure the relation is defined in post model struct tag.
func (o *queryM2M) Add(mds ...interface{}) (int64, error) {
	return 0, nil
}

// remove models following the origin model relationship
func (o *queryM2M) Remove(mds ...interface{}) (int64, error) {
	return 0, nil
}

// check model is existed in relationship of origin model
func (o *queryM2M) Exist(md interface{}) bool {
	return false
}

// clean all models in related of origin model
func (o *queryM2M) Clear() (int64, error) {
	return 0, nil
}

// count all related models of origin model
func (o *queryM2M) Count() (int64, error) {
	return 0, nil
}

// create new M2M queryer.
func newQueryM2M(md interface{}, o *orm, mi *modelInfo, fi *fieldInfo, ind reflect.Value) QueryM2Mer {
	return nil
}
