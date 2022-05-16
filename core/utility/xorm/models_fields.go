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
	"time"
)

// Define the Type enum
const (
	TypeBooleanField = 1 << iota
	TypeVarCharField
	TypeCharField
	TypeTextField
	TypeTimeField
	TypeDateField
	TypeDateTimeField
	TypeBitField
	TypeSmallIntegerField
	TypeIntegerField
	TypeBigIntegerField
	TypePositiveBitField
	TypePositiveSmallIntegerField
	TypePositiveIntegerField
	TypePositiveBigIntegerField
	TypeFloatField
	TypeDecimalField
	TypeJSONField
	TypeJsonbField
	RelForeignKey
	RelOneToOne
	RelManyToMany
	RelReverseOne
	RelReverseMany
)

// Define some logic enum
const (
	IsIntegerField         = ^-TypePositiveBigIntegerField >> 6 << 7
	IsPositiveIntegerField = ^-TypePositiveBigIntegerField >> 10 << 11
	IsRelField             = ^-RelReverseMany >> 18 << 19
	IsFieldType            = ^-RelReverseMany<<1 + 1
)

// BooleanField A true/false field.
type BooleanField bool

// Value return the BooleanField
func (e BooleanField) Value() bool {
	return false
}

// Set will set the BooleanField
func (e *BooleanField) Set(d bool) {
	return
}

// String format the Bool to string
func (e *BooleanField) String() string {
	return ""
}

// FieldType return BooleanField the type
func (e *BooleanField) FieldType() int {
	return 0
}

// SetRaw set the interface to bool
func (e *BooleanField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return the current value
func (e *BooleanField) RawValue() interface{} {
	return nil
}

// verify the BooleanField implement the Fielder interface
// CharField A string field
// required values tag: size
// The size is enforced at the database level and in models’s validation.
// eg: `orm:"size(120)"`
type CharField string

// Value return the CharField's Value
func (e CharField) Value() string {
	return ""
}

// Set CharField value
func (e *CharField) Set(d string) {
	return
}

// String return the CharField
func (e *CharField) String() string {
	return ""
}

// FieldType return the enum type
func (e *CharField) FieldType() int {
	return 0
}

// SetRaw set the interface to string
func (e *CharField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return the CharField value
func (e *CharField) RawValue() interface{} {
	return nil
}

// verify CharField implement Fielder
// TimeField A time, represented in go by a time.Time instance.
// only time values like 10:00:00
// Has a few extra, optional attr tag:
//
// auto_now:
// Automatically set the field to now every time the object is saved. Useful for “last-modified” timestamps.
// Note that the current date is always used; it’s not just a default value that you can override.
//
// auto_now_add:
// Automatically set the field to now when the object is first created. Useful for creation of timestamps.
// Note that the current date is always used; it’s not just a default value that you can override.
//
// eg: `orm:"auto_now"` or `orm:"auto_now_add"`
type TimeField time.Time

// Value return the time.Time
func (e TimeField) Value() time.Time {
	return time.Time{}
}

// Set set the TimeField's value
func (e *TimeField) Set(d time.Time) {
	return
}

// String convert time to string
func (e *TimeField) String() string {
	return ""
}

// FieldType return enum type Date
func (e *TimeField) FieldType() int {
	return 0
}

// SetRaw convert the interface to time.Time. Allow string and time.Time
func (e *TimeField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return time value
func (e *TimeField) RawValue() interface{} {
	return nil
}

// DateField A date, represented in go by a time.Time instance.
// only date values like 2006-01-02
// Has a few extra, optional attr tag:
//
// auto_now:
// Automatically set the field to now every time the object is saved. Useful for “last-modified” timestamps.
// Note that the current date is always used; it’s not just a default value that you can override.
//
// auto_now_add:
// Automatically set the field to now when the object is first created. Useful for creation of timestamps.
// Note that the current date is always used; it’s not just a default value that you can override.
//
// eg: `orm:"auto_now"` or `orm:"auto_now_add"`
type DateField time.Time

// Value return the time.Time
func (e DateField) Value() time.Time {
	return time.Time{}
}

// Set set the DateField's value
func (e *DateField) Set(d time.Time) {
	return
}

// String convert datetime to string
func (e *DateField) String() string {
	return ""
}

// FieldType return enum type Date
func (e *DateField) FieldType() int {
	return 0
}

// SetRaw convert the interface to time.Time. Allow string and time.Time
func (e *DateField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return Date value
func (e *DateField) RawValue() interface{} {
	return nil
}

// verify DateField implement fielder interface
// DateTimeField A date, represented in go by a time.Time instance.
// datetime values like 2006-01-02 15:04:05
// Takes the same extra arguments as DateField.
type DateTimeField time.Time

// Value return the datetime value
func (e DateTimeField) Value() time.Time {
	return time.Time{}
}

// Set set the time.Time to datetime
func (e *DateTimeField) Set(d time.Time) {
	return
}

// String return the time's String
func (e *DateTimeField) String() string {
	return ""
}

// FieldType return the enum TypeDateTimeField
func (e *DateTimeField) FieldType() int {
	return 0
}

// SetRaw convert the string or time.Time to DateTimeField
func (e *DateTimeField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return the datetime value
func (e *DateTimeField) RawValue() interface{} {
	return nil
}

// verify datetime implement fielder
// FloatField A floating-point number represented in go by a float32 value.
type FloatField float64

// Value return the FloatField value
func (e FloatField) Value() float64 {
	return 0
}

// Set the Float64
func (e *FloatField) Set(d float64) {
	return
}

// String return the string
func (e *FloatField) String() string {
	return ""
}

// FieldType return the enum type
func (e *FloatField) FieldType() int {
	return 0
}

// SetRaw converter interface Float64 float32 or string to FloatField
func (e *FloatField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return the FloatField value
func (e *FloatField) RawValue() interface{} {
	return nil
}

// verify FloatField implement Fielder
// SmallIntegerField -32768 to 32767
type SmallIntegerField int16

// Value return int16 value
func (e SmallIntegerField) Value() int16 {
	return 0
}

// Set the SmallIntegerField value
func (e *SmallIntegerField) Set(d int16) {
	return
}

// String convert smallint to string
func (e *SmallIntegerField) String() string {
	return ""
}

// FieldType return enum type SmallIntegerField
func (e *SmallIntegerField) FieldType() int {
	return 0
}

// SetRaw convert interface int16/string to int16
func (e *SmallIntegerField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return smallint value
func (e *SmallIntegerField) RawValue() interface{} {
	return nil
}

// verify SmallIntegerField implement Fielder
// IntegerField -2147483648 to 2147483647
type IntegerField int32

// Value return the int32
func (e IntegerField) Value() int32 {
	return 0
}

// Set IntegerField value
func (e *IntegerField) Set(d int32) {
	return
}

// String convert Int32 to string
func (e *IntegerField) String() string {
	return ""
}

// FieldType return the enum type
func (e *IntegerField) FieldType() int {
	return 0
}

// SetRaw convert interface int32/string to int32
func (e *IntegerField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return IntegerField value
func (e *IntegerField) RawValue() interface{} {
	return nil
}

// verify IntegerField implement Fielder
// BigIntegerField -9223372036854775808 to 9223372036854775807.
type BigIntegerField int64

// Value return int64
func (e BigIntegerField) Value() int64 {
	return 0
}

// Set the BigIntegerField value
func (e *BigIntegerField) Set(d int64) {
	return
}

// String convert BigIntegerField to string
func (e *BigIntegerField) String() string {
	return ""
}

// FieldType return enum type
func (e *BigIntegerField) FieldType() int {
	return 0
}

// SetRaw convert interface int64/string to int64
func (e *BigIntegerField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return BigIntegerField value
func (e *BigIntegerField) RawValue() interface{} {
	return nil
}

// verify BigIntegerField implement Fielder
// PositiveSmallIntegerField 0 to 65535
type PositiveSmallIntegerField uint16

// Value return uint16
func (e PositiveSmallIntegerField) Value() uint16 {
	return 0
}

// Set PositiveSmallIntegerField value
func (e *PositiveSmallIntegerField) Set(d uint16) {
	return
}

// String convert uint16 to string
func (e *PositiveSmallIntegerField) String() string {
	return ""
}

// FieldType return enum type
func (e *PositiveSmallIntegerField) FieldType() int {
	return 0
}

// SetRaw convert Interface uint16/string to uint16
func (e *PositiveSmallIntegerField) SetRaw(value interface{}) error {
	return nil
}

// RawValue returns PositiveSmallIntegerField value
func (e *PositiveSmallIntegerField) RawValue() interface{} {
	return nil
}

// verify PositiveSmallIntegerField implement Fielder
// PositiveIntegerField 0 to 4294967295
type PositiveIntegerField uint32

// Value return PositiveIntegerField value. Uint32
func (e PositiveIntegerField) Value() uint32 {
	return 0
}

// Set the PositiveIntegerField value
func (e *PositiveIntegerField) Set(d uint32) {
	return
}

// String convert PositiveIntegerField to string
func (e *PositiveIntegerField) String() string {
	return ""
}

// FieldType return enum type
func (e *PositiveIntegerField) FieldType() int {
	return 0
}

// SetRaw convert interface uint32/string to Uint32
func (e *PositiveIntegerField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return the PositiveIntegerField Value
func (e *PositiveIntegerField) RawValue() interface{} {
	return nil
}

// verify PositiveIntegerField implement Fielder
// PositiveBigIntegerField 0 to 18446744073709551615
type PositiveBigIntegerField uint64

// Value return uint64
func (e PositiveBigIntegerField) Value() uint64 {
	return 0
}

// Set PositiveBigIntegerField value
func (e *PositiveBigIntegerField) Set(d uint64) {
	return
}

// String convert PositiveBigIntegerField to string
func (e *PositiveBigIntegerField) String() string {
	return ""
}

// FieldType return enum type
func (e *PositiveBigIntegerField) FieldType() int {
	return 0
}

// SetRaw convert interface uint64/string to Uint64
func (e *PositiveBigIntegerField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return PositiveBigIntegerField value
func (e *PositiveBigIntegerField) RawValue() interface{} {
	return nil
}

// verify PositiveBigIntegerField implement Fielder
// TextField A large text field.
type TextField string

// Value return TextField value
func (e TextField) Value() string {
	return ""
}

// Set the TextField value
func (e *TextField) Set(d string) {
	return
}

// String convert TextField to string
func (e *TextField) String() string {
	return ""
}

// FieldType return enum type
func (e *TextField) FieldType() int {
	return 0
}

// SetRaw convert interface string to string
func (e *TextField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return TextField value
func (e *TextField) RawValue() interface{} {
	return nil
}

// verify TextField implement Fielder
// JSONField postgres json field.
type JSONField string

// Value return JSONField value
func (j JSONField) Value() string {
	return ""
}

// Set the JSONField value
func (j *JSONField) Set(d string) {
	return
}

// String convert JSONField to string
func (j *JSONField) String() string {
	return ""
}

// FieldType return enum type
func (j *JSONField) FieldType() int {
	return 0
}

// SetRaw convert interface string to string
func (j *JSONField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return JSONField value
func (j *JSONField) RawValue() interface{} {
	return nil
}

// verify JSONField implement Fielder
// JsonbField postgres json field.
type JsonbField string

// Value return JsonbField value
func (j JsonbField) Value() string {
	return ""
}

// Set the JsonbField value
func (j *JsonbField) Set(d string) {
	return
}

// String convert JsonbField to string
func (j *JsonbField) String() string {
	return ""
}

// FieldType return enum type
func (j *JsonbField) FieldType() int {
	return 0
}

// SetRaw convert interface string to string
func (j *JsonbField) SetRaw(value interface{}) error {
	return nil
}

// RawValue return JsonbField value
func (j *JsonbField) RawValue() interface{} {
	return nil
}

// verify JsonbField implement Fielder
