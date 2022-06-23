//go:binary-only-package
//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//
package xorm

import (
	_ "fmt"
	_ "reflect"
	_ "strconv"
	"time"
	_ "unsafe"

	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xjson"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xlog"
	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xobj"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xstring"
	_ "github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xtime"
)

const (
	ORM_STRING_UNDEFINED string = "orm_undefined"
	ORM_INT_UNDEFINED    int    = -0x7fffffffffffffff - 1
	ORM_INT64_UNDEFINED  int64  = -0x7fffffffffffffff - 1
	ORM_INT32_UNDEFINED  int32  = -0x7fffffff - 1
)
const (
	LOCK_SLEEP_DEVIDE time.Duration = 30 // 锁的间隔时间，1除该值
)

type ITable interface {
	// 构造函数
	CTOR(CHILD interface{})
	// 设置子对象
	SetCHILD(CHILD interface{})
	// ORM数据模型
	ORMModel() *ModelInfo
	// ORM数据字段
	ORMField(name string) *FieldInfo
	// Redis数据字段
	RDXField(name string) *FieldRedis
	// ORM或Redis指定字段的值
	FLDValue(name string, value ...interface{}) interface{}
	// 是否关联DB
	LinkDB() bool
	// 是否关联Redis
	LinkRedis() bool
	// 是否关联缓存
	LinkCache() bool
	// 仅关联缓存
	OnlyCache() bool
	// 仅关联Redis
	OnlyRedis() bool
	// 仅关联DB
	OnlyDB() bool
	// 唯一标识符（不可修改）
	UUID() string
	// 数据库名称
	DBName() string
	// 表格名称
	TableName() string
	// 表格别名（常用于Redis区分前缀）
	TableAlias() string
	// 主键名称
	PKeyName() string
	// 主键值
	PKeyValue() interface{}
	// Redis名称
	RedisName() string
	// Redis前缀（常用区分内存域，GList/GClear）
	RedisPrefix() string
	// Redis加锁（若未指定sec则默认20秒自动释放锁）
	// [ argcount1:block-是否阻塞（默认阻塞）]
	// [ argcount2:block-是否阻塞（默认阻塞）, sec-超时时间]
	RedisLock(args ...interface{}) bool
	// Redis解锁
	RedisUnlock()
	// Redis时效
	RedisExpire(sec int) int
	// 获取DB数据数量（仅针对LinkDB有效）
	DataCount(_cond ...*Condition) int
	// DB指定Column（列）的最大值，若未指定Column则获取主键的最大值（仅针对LinkDB有效）
	MaxIndex(column ...string) int
	// DB指定Column（列）的最小值，若未指定Column则获取主键的最小值（仅针对LinkDB有效）
	MinIndex(column ...string) int
	// 从DB和Redis中删除数据
	Delete() bool
	// 从DB中删除数据
	DeleteFromDB() (affect int, e error)
	// 从Redis中删除数据
	DeleteFromRedis() bool
	// 写入数据至DB和Redis
	Write(ttl ...int) bool
	// 写入数据至DB
	WriteToDB() (affect int, e error)
	// 写入数据至Redis
	// [ argcount1:ttl ]
	WriteToRedis(ttl ...int) bool
	// 从DB或Redis中读取数据
	// [ _cond: 查询条件]
	// [ 若LinkRedis为true则同步写入redis ]
	Read(_cond ...*Condition) bool
	// 从DB中读取数据
	// [ _cond: 查询条件]
	ReadFromDB(_cond ...*Condition) error
	// 从Redis中读取数据
	// [ argcount2:key/kvp-isnew ]
	// [ argcount1:key ]
	ReadFromRedis(args ...interface{}) (objptr interface{}, ok bool)
	// [ 从Redis或DB中读取list形式数据, rets := []*Table{} ]
	// [ __gt-大于 ]
	// [ __gte-大于等于 ]
	// [ __lt-小于 ]
	// [ __lte-小于等于 ]
	// [ __in-特定（最多65534个KEY）]
	// [ __exact-等于 ]
	// [ __ne-不等于]
	// [ __contains-包含 ]
	// [ __startswith-以...起始 ]
	// [ __endswith-以...结束 ]
	// [ __isnull-是否为空 ]
	// [ _cond: 查询条件]
	// [ 若LinkRedis为true则同步写入redis ]
	// [ 其中rets类型为 *[]*Type ]
	List(rets interface{}, _cond ...*Condition) bool
	// [ 从DB中读取list形式数据, rets := []*Table{} ]
	// [ __gt-大于 ]
	// [ __gte-大于等于 ]
	// [ __lt-小于 ]
	// [ __lte-小于等于 ]
	// [ __in-特定（最多65534个KEY）]
	// [ __exact-等于 ]
	// [ __ne-不等于]
	// [ __contains-包含 ]
	// [ __startswith-以...起始 ]
	// [ __endswith-以...结束 ]
	// [ __isnull-是否为空 ]
	// [ _cond: 查询条件]
	// [ 其中rets类型为 *[]*Type ]
	ListFromDB(rets interface{}, _cond ...*Condition) (count int, err error)
	// [ 从Redis中读取list形式数据, rets := []*Table{} ]
	// [ __gt-大于 ]
	// [ __gte-大于等于 ]
	// [ __lt-小于 ]
	// [ __lte-小于等于 ]
	// [ __in-特定（最多65534个KEY）]
	// [ __exact-等于 ]
	// [ __ne-不等于]
	// [ __contains-包含 ]
	// [ __startswith-以...起始 ]
	// [ __endswith-以...结束 ]
	// [ __isnull-是否为空 ]
	// [ _cond: 查询条件]
	// [ 其中rets类型为 *[]*Type ]
	ListFromRedis(rets interface{}, _cond ...*Condition) bool
	// 从DB和Redis中清空所有关联数据
	// [ _cond: 查询条件]
	Clear(_cond ...*Condition) bool
	// 从DB中清空所有关联数据
	// [ _cond: 查询条件]
	ClearFromDB(_cond ...*Condition) (affect int, e error)
	// 从Redis中清空所有关联数据
	// [ _cond: 查询条件]
	ClearFromRedis(_cond ...*Condition) bool
	// 设置是否关联DB（默认false）
	LDB(sig bool) ITable
	// 设置是否关联Redis（默认false）
	LRedis(sig bool) ITable
	// 设置是否关联Cache（sync-设置同步时间，<=0-不同步）（默认false）
	LCache(sig bool, sync ...int) ITable
	// 设置是否读写（全局）（默认false）
	LRW(sig bool) ITable
	// 是否读写（默认为可读可写-true）
	RW(value ...bool) bool
	// 是否读写（用于会话标记，0-未标记，1-只读，2-可读可写）
	SRW(value ...int) int
	// Wirte时不调用Encode函数
	NoEncode(value ...bool) bool
	// 对象封装（写数据之前被调用）
	Encode(toDB bool)
	// 对象解析（读数据之后被调用）
	Decode(fromDB bool)
	// 是否为有效对象（从DB/Redis中创建的或即将存入DB/Redis的对象）
	IsValid(value ...bool) bool
	// 浅拷贝对象（值类型）（init-调用CTOR和Decode初始化）（若传入src，则以src作为源对象进行拷贝，否则使用this进行拷贝）
	Clone(init bool, _src ...interface{}) interface{}
	// 对象ToString（json格式）
	ToString() string
	// 依据参数获取条件
	EnsureCond(_cond ...*Condition) *Condition
	// 条件（第一层）是否皆为pk和ok
	IsAllPOK(cond *Condition) bool
	// 附加pk和ok至条件（第一层）中（该函数亦会给pk和ok赋值）（一般用于db操作，或为redis生成name/prefix）
	AttachPOK(cond *Condition) *Condition
	// 从条件（第一层）中剔除pk和ok（一般用于redis操作，避免多余的查询）
	DetachPOK(cond *Condition) *Condition
}
type Table struct {
	xobj.OBJECT `orm:"-" json:"-"`
	REAL        ITable `orm:"-" json:"-"`
}

// 构造函数
func (_this *Table) CTOR(CHILD interface{}) {
	return
}

// 设置子对象
func (_this *Table) SetCHILD(CHILD interface{}) {
	return
}

// ORM数据模型
func (_this *Table) ORMModel() *ModelInfo {
	return nil
}

// ORM数据字段
func (_this *Table) ORMField(name string) *FieldInfo {
	return nil
}

// Redis数据字段
func (_this *Table) RDXField(name string) *FieldRedis {
	return nil
}

// ORM或Redis指定字段的值
func (_this *Table) FLDValue(name string, value ...interface{}) interface{} {
	return nil
}

// 是否关联DB
func (_this *Table) LinkDB() bool {
	return false
}

// 是否关联Redis
func (_this *Table) LinkRedis() bool {
	return false
}

// 是否关联缓存
func (_this *Table) LinkCache() bool {
	return false
}

// 仅关联缓存
func (_this *Table) OnlyCache() bool {
	return false
}

// 仅关联Redis
func (_this *Table) OnlyRedis() bool {
	return false
}

// 仅关联DB
func (_this *Table) OnlyDB() bool {
	return false
}

// 唯一标识符（不可修改）
func (_this *Table) UUID() string {
	return ""
}

// 数据库名称
func (_this *Table) DBName() string {
	return ""
}

// 表格名称
func (_this *Table) TableName() string {
	return ""
}

// 表格别名（常用于Redis区分前缀）
func (_this *Table) TableAlias() string {
	return ""
}

// 主键名称
func (_this *Table) PKeyName() string {
	return ""
}

// 主键值
// [ argcount1:value ]
func (_this *Table) PKeyValue() interface{} {
	return nil
}

// Redis名称
func (_this *Table) RedisName() string {
	return ""
}

// Redis前缀（常用区分内存域，GList/GClear）
func (_this *Table) RedisPrefix() string {
	return ""
}

// Redis加锁（若未指定sec则默认20秒自动释放锁）
// [ argcount1:block-是否阻塞（默认阻塞）]
// [ argcount2:block-是否阻塞（默认阻塞）, sec-超时时间]
func (_this *Table) RedisLock(args ...interface{}) bool {
	return false
}

// Redis解锁
func (_this *Table) RedisUnlock() {
	return
}

// Redis时效
func (_this *Table) RedisExpire(sec int) int {
	return 0
}

// 获取DB数据数量（仅针对LinkDB有效）
func (_this *Table) DataCount(_cond ...*Condition) int {
	return 0
}

// DB指定Column（列）的最大值，若未指定Column则获取主键的最大值（仅针对LinkDB有效）
func (_this *Table) MaxIndex(column ...string) int {
	return 0
}

// DB指定Column（列）的最小值，若未指定Column则获取主键的最小值（仅针对LinkDB有效）
func (_this *Table) MinIndex(column ...string) int {
	return 0
}

// 从DB和Redis中删除数据
func (_this *Table) Delete() bool {
	return false
}

// 从DB中删除数据
func (_this *Table) DeleteFromDB() (affect int, e error) {
	return affect, e
}

// 从Redis中删除数据
func (_this *Table) DeleteFromRedis() bool {
	return false
}

// 写入数据至DB和Redis
func (_this *Table) Write(ttl ...int) bool {
	return false
}

// 写入数据至DB
func (_this *Table) WriteToDB() (affect int, e error) {
	return affect, e
}

// 写入数据至Redis
// [ argcount1:ttl ]
func (_this *Table) WriteToRedis(ttl ...int) bool {
	return false
}

// 从DB或Redis中读取数据
// [ _cond: 查询条件]
// [ 若LinkRedis为true则同步写入redis ]
func (_this *Table) Read(_cond ...*Condition) bool {
	return false
}

// 从DB中读取数据
// [ _cond: 查询条件]
func (_this *Table) ReadFromDB(_cond ...*Condition) error {
	return nil
}

// 从Redis中读取数据
// [ argcount2:key-isnew ]
// [ argcount1:key ]
func (_this *Table) ReadFromRedis(args ...interface{}) (objptr interface{}, ok bool) {
	return objptr, ok
}

// [ 从Redis或DB中读取list形式数据, rets := []*Table{} ]
// [ __gt-大于 ]
// [ __gte-大于等于 ]
// [ __lt-小于 ]
// [ __lte-小于等于 ]
// [ __in-特定（最多65534个KEY）]
// [ __exact-等于 ]
// [ __ne-不等于]
// [ __contains-包含 ]
// [ __startswith-以...起始 ]
// [ __endswith-以...结束 ]
// [ __isnull-是否为空 ]
// [ _cond: 查询条件]
// [ 若LinkRedis为true则同步写入redis ]
// [ 其中rets类型为 *[]*Type ]
func (_this *Table) List(rets interface{}, _cond ...*Condition) bool {
	return false
}

// [ 从DB中读取list形式数据, rets := []*Table{} ]
// [ __gt-大于 ]
// [ __gte-大于等于 ]
// [ __lt-小于 ]
// [ __lte-小于等于 ]
// [ __in-特定（最多65534个KEY）]
// [ __exact-等于 ]
// [ __ne-不等于]
// [ __contains-包含 ]
// [ __startswith-以...起始 ]
// [ __endswith-以...结束 ]
// [ __isnull-是否为空 ]
// [ _cond: 查询条件]
// [ 其中rets类型为 *[]*Type ]
func (_this *Table) ListFromDB(rets interface{}, _cond ...*Condition) (count int, err error) {
	return count, err
}

// [ 从Redis中读取list形式数据, rets := []*Table{} ]
// [ __gt-大于 ]
// [ __gte-大于等于 ]
// [ __lt-小于 ]
// [ __lte-小于等于 ]
// [ __in-特定（最多65534个KEY）]
// [ __exact-等于 ]
// [ __ne-不等于]
// [ __contains-包含 ]
// [ __startswith-以...起始 ]
// [ __endswith-以...结束 ]
// [ __isnull-是否为空 ]
// [ _cond: 查询条件]
// [ 其中rets类型为 *[]*Type ]
func (_this *Table) ListFromRedis(rets interface{}, _cond ...*Condition) bool {
	return false
}

// 从DB和Redis中清空所有关联数据
// [ _cond: 查询条件]
func (_this *Table) Clear(_cond ...*Condition) bool {
	return false
}

// 从DB中清空所有关联数据
// [ _cond: 查询条件]
func (_this *Table) ClearFromDB(_cond ...*Condition) (affect int, e error) {
	return affect, e
}

// 从Redis中清空所有关联数据
// [ _cond: 查询条件]
func (_this *Table) ClearFromRedis(_cond ...*Condition) bool {
	return false
}

// 浅拷贝对象（值类型）（init-调用CTOR和Decode初始化）（若传入src，则以src作为源对象进行拷贝，否则使用this进行拷贝）
func (_this *Table) Clone(init bool, _src ...interface{}) interface{} {
	return nil
}

// 对象转为字符串（json格式）
func (_this *Table) ToString() string {
	return ""
}

// 设置是否关联DB（默认false）
func (_this *Table) LDB(sig bool) ITable {
	return nil
}

// 设置是否关联Redis（默认false）
func (_this *Table) LRedis(sig bool) ITable {
	return nil
}

// 设置是否关联Cache（sync-设置同步时间，<=0-不同步）（默认false）
func (_this *Table) LCache(sig bool, sync ...int) ITable {
	return nil
}

// 设置是否读写（全局）（默认false）
func (_this *Table) LRW(sig bool) ITable {
	return nil
}

// 是否读写（默认为可读可写-true）
func (_this *Table) RW(value ...bool) bool {
	return false
}

// 是否读写（用于会话标记，0-未标记，1-只读，2-可读可写）
func (_this *Table) SRW(value ...int) int {
	return 0
}

// Wirte时不调用Encode函数
func (_this *Table) NoEncode(value ...bool) bool {
	return false
}

// 对象封装（写数据之前被调用）
func (_this *Table) Encode(toDB bool) {
	return
}

// 对象解析（读数据之后被调用）
func (_this *Table) Decode(fromDB bool) {
	return
}

// 是否为有效对象（从DB/Redis中创建的或即将存入DB/Redis的对象）
func (_this *Table) IsValid(value ...bool) bool {
	return false
}

// 依据参数获取条件
func (_this *Table) EnsureCond(_cond ...*Condition) *Condition {
	return nil
}

// 条件（第一层）是否皆为pk和ok
func (_this *Table) IsAllPOK(cond *Condition) bool {
	return false
}

// 附加pk和ok至条件（第一层）中（该函数亦会给pk和ok赋值）（一般用于db操作，或为redis生成name/prefix）
func (_this *Table) AttachPOK(cond *Condition) *Condition {
	return nil
}

// 从条件（第一层）中剔除pk和ok（一般用于redis操作，避免多余的查询）
func (_this *Table) DetachPOK(cond *Condition) *Condition {
	return nil
}
