//go:binary-only-package
package xcollect

// 从数组中索引元素（int32），ele为元素
func IndexForInt32(arr []int32, ele int32) int {
	return 0
}

// 判断数组是否存在元素（int32），ele为元素
func ContainsForInt32(arr []int32, ele int32) bool {
	return false
}

// 从数组中移除元素（int32），ele为元素
func RemoveForInt32(arr []int32, ele int32) []int32 {
	return []int32{}
}

// 从数组中移除元素（int32），ele为索引
func DeleteForInt32(arr []int32, ele int) []int32 {
	return []int32{}
}

// 在数组中新增元素（int32），ele为元素
func AppendForInt32(arr []int32, ele int32) []int32 {
	return []int32{}
}
