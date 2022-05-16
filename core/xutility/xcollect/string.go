//go:binary-only-package
package xcollect

// 从数组中索引元素（string），ele为元素
func IndexForStr(arr []string, ele string) int {
	return 0
}

// 判断数组是否存在元素（string），ele为元素
func ContainsForStr(arr []string, ele string) bool {
	return false
}

// 从数组中移除元素（string），ele为元素
func RemoveForStr(arr []string, ele string) []string {
	return []string{}
}

// 从数组中移除元素（string），ele为索引
func DeleteForStr(arr []string, ele int) []string {
	return []string{}
}

// 在数组中新增元素（string），ele为元素
func AppendForStr(arr []string, ele string) []string {
	return []string{}
}
