package utils

import "strings"

// 判断是否包含特殊字符
func ContainsSpecialChar(str string) bool {
	// 定义特殊字符的集合
	specialChars := "!@#$%^&*()_+-=[]{};:'\"\\|<>,./?"

	// 遍历字符串的每个字符
	for _, ch := range str {
		// 将字符转换为字符串形式
		chStr := string(ch)

		// 判断字符是否在特殊字符集合中
		if strings.Contains(specialChars, chStr) {
			return true
		}
	}

	return false
}
