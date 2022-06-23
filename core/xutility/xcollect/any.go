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

// 从数组中索引元素（interface{}），ele为元素
func IndexForAny(arr []interface{}, ele interface{}) int {
	return 0
}

// 判断数组是否存在元素（interface{}），ele为元素
func ContainsForAny(arr []interface{}, ele interface{}) bool {
	return false
}

// 从数组中移除元素（interface{}），ele为元素
func RemoveForAny(arr []interface{}, ele interface{}) []interface{} {
	return []interface{}{}
}

// 从数组中移除元素（interface{}），ele为索引
func DeleteForAny(arr []interface{}, ele int) []interface{} {
	return []interface{}{}
}

// 在数组中新增元素（interface{}），ele为元素
func AppendForAny(arr []interface{}, ele interface{}) []interface{} {
	return []interface{}{}
}
