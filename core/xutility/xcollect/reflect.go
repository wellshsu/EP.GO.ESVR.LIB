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
package xcollect

import (
	_ "reflect"

	_ "github.com/wellshsu/EP.GO.ESVR.LIB/core/xutility/xlog"
)

// [Deprecated]判断集合是否存在元素（支持指针类型的map和slice），compare为元素或对比函数
func Contains(collection interface{}, compare interface{}) bool {
	return false
}

// [Deprecated]从集合中查找元素（支持指针类型的map和slice），compare为元素或对比函数
func Find(collection interface{}, compare interface{}) interface{} {
	return nil
}

// [Deprecated]从集合中索引元素（支持指针类型的map和slice），compare为元素或对比函数，当集合为slice时返回number，map则返回相应key
func Index(collection interface{}, compare interface{}) interface{} {
	return nil
}

// [Deprecated]在集合中新增元素（支持指针类型的slice），element为所要新增的元素
func Append(collection interface{}, element interface{}) bool {
	return false
}

// [Deprecated]从集合中移除元素（支持指针类型的map和slice），compare为元素或对比函数
func Remove(collection interface{}, compare interface{}) bool {
	return false
}

// [Deprecated]从集合中移除元素（支持指针类型的map和slice），compare为索引或key值
func Delete(collection interface{}, compare interface{}) bool {
	return false
}

// [Deprecated]清空集合中的元素（支持指针类型的map和slice）
func Clear(collection interface{}) bool {
	return false
}
