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
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io/ioutil"
	_ "os"
	_ "strings"
	"sync"
)

// JSONConfig is a json config parser and implements Config interface.
type JSONConfig struct {
}

// Parse returns a ConfigContainer with parsed json config map.
func (js *JSONConfig) Parse(filename string) (Configer, error) {
	return nil, nil
}

// ParseData returns a ConfigContainer with json string
func (js *JSONConfig) ParseData(data []byte) (Configer, error) {
	return nil, nil
}

// JSONConfigContainer A Config represents the json configuration.
// Only when get value, support key as section:name type.
type JSONConfigContainer struct {
	sync.RWMutex
}

// Bool returns the boolean value for a given key.
func (c *JSONConfigContainer) Bool(key string) (bool, error) {
	return false, nil
}

// DefaultBool return the bool value if has no error
// otherwise return the defaultval
func (c *JSONConfigContainer) DefaultBool(key string, defaultval bool) bool {
	return false
}

// Int returns the integer value for a given key.
func (c *JSONConfigContainer) Int(key string) (int, error) {
	return 0, nil
}

// DefaultInt returns the integer value for a given key.
// if err != nil return defaultval
func (c *JSONConfigContainer) DefaultInt(key string, defaultval int) int {
	return 0
}

// Int64 returns the int64 value for a given key.
func (c *JSONConfigContainer) Int64(key string) (int64, error) {
	return 0, nil
}

// DefaultInt64 returns the int64 value for a given key.
// if err != nil return defaultval
func (c *JSONConfigContainer) DefaultInt64(key string, defaultval int64) int64 {
	return 0
}

// Float returns the float value for a given key.
func (c *JSONConfigContainer) Float(key string) (float64, error) {
	return 0, nil
}

// DefaultFloat returns the float64 value for a given key.
// if err != nil return defaultval
func (c *JSONConfigContainer) DefaultFloat(key string, defaultval float64) float64 {
	return 0
}

// String returns the string value for a given key.
func (c *JSONConfigContainer) String(key string) string {
	return ""
}

// DefaultString returns the string value for a given key.
// if err != nil return defaultval
func (c *JSONConfigContainer) DefaultString(key string, defaultval string) string {
	return ""
}

// Strings returns the []string value for a given key.
func (c *JSONConfigContainer) Strings(key string) []string {
	return []string{}
}

// DefaultStrings returns the []string value for a given key.
// if err != nil return defaultval
func (c *JSONConfigContainer) DefaultStrings(key string, defaultval []string) []string {
	return []string{}
}

// GetSection returns map for the given section
func (c *JSONConfigContainer) GetSection(section string) (map[string]string, error) {
	return nil, nil
}

// SaveConfigFile save the config into file
func (c *JSONConfigContainer) SaveConfigFile(filename string) (err error) {
	return err
}

// Set writes a new value for key.
func (c *JSONConfigContainer) Set(key, val string) error {
	return nil
}

// DIY returns the raw value by a given key.
func (c *JSONConfigContainer) DIY(key string) (v interface{}, err error) {
	return v, err
}

// section.key or key
func (c *JSONConfigContainer) getData(key string) interface{} {
	return nil
}
func init() {
	return
}
