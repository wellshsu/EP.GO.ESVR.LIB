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
package xconfig

import (
	_ "errors"
	_ "strconv"
	_ "strings"
)

type fakeConfigContainer struct {
}

func (c *fakeConfigContainer) getData(key string) string {
	return ""
}
func (c *fakeConfigContainer) Set(key, val string) error {
	return nil
}
func (c *fakeConfigContainer) String(key string) string {
	return ""
}
func (c *fakeConfigContainer) DefaultString(key string, defaultval string) string {
	return ""
}
func (c *fakeConfigContainer) Strings(key string) []string {
	return []string{}
}
func (c *fakeConfigContainer) DefaultStrings(key string, defaultval []string) []string {
	return []string{}
}
func (c *fakeConfigContainer) Int(key string) (int, error) {
	return 0, nil
}
func (c *fakeConfigContainer) DefaultInt(key string, defaultval int) int {
	return 0
}
func (c *fakeConfigContainer) Int64(key string) (int64, error) {
	return 0, nil
}
func (c *fakeConfigContainer) DefaultInt64(key string, defaultval int64) int64 {
	return 0
}
func (c *fakeConfigContainer) Bool(key string) (bool, error) {
	return false, nil
}
func (c *fakeConfigContainer) DefaultBool(key string, defaultval bool) bool {
	return false
}
func (c *fakeConfigContainer) Float(key string) (float64, error) {
	return 0, nil
}
func (c *fakeConfigContainer) DefaultFloat(key string, defaultval float64) float64 {
	return 0
}
func (c *fakeConfigContainer) DIY(key string) (interface{}, error) {
	return nil, nil
}
func (c *fakeConfigContainer) GetSection(section string) (map[string]string, error) {
	return nil, nil
}
func (c *fakeConfigContainer) SaveConfigFile(filename string) error {
	return nil
}

// NewFakeConfig return a fake Configer
func NewFakeConfig() Configer {
	return nil
}
