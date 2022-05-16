//go:binary-only-package
package xcollect

// 从数组中索引元素（int），ele为元素
func IndexForInt(arr []int, ele int) int {
	return 0
}

// 判断数组是否存在元素（int），ele为元素
func ContainsForInt(arr []int, ele int) bool {
	return false
}

// 从数组中移除元素（int），ele为元素
func RemoveForInt(arr []int, ele int) []int {
	return []int{}
}

// 从数组中移除元素（int），ele为索引
func DeleteForInt(arr []int, ele int) []int {
	return []int{}
}

// 在数组中新增元素（int），ele为元素
func AppendForInt(arr []int, ele int) []int {
	return []int{}
}
