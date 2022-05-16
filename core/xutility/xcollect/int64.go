//go:binary-only-package
package xcollect

// 从数组中索引元素（int64），ele为元素
func IndexForInt64(arr []int64, ele int64) int {
	return 0
}

// 判断数组是否存在元素（int64），ele为元素
func ContainsForInt64(arr []int64, ele int64) bool {
	return false
}

// 从数组中移除元素（int64），ele为元素
func RemoveForInt64(arr []int64, ele int64) []int64 {
	return []int64{}
}

// 从数组中移除元素（int64），ele为索引
func DeleteForInt64(arr []int64, ele int) []int64 {
	return []int64{}
}

// 在数组中新增元素（int64），ele为元素
func AppendForInt64(arr []int64, ele int64) []int64 {
	return []int64{}
}
