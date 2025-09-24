package main

import "fmt"

// 查找字符串最长前缀
func findLongestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 找最短字符串长度，避免越界
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	var prefix []byte // 使用字节切片存储字符

	// 逐位比较
	for i := 0; i < minLen; i++ {
		char := strs[0][i] // 第一个字符串的第i个字符

		// 检查所有字符串的第i位是否相同
		allMatch := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				allMatch = false
				break
			}
		}

		if !allMatch {
			break // 发现不匹配，停止
		}

		// 所有字符串的第i位都相同，添加到前缀
		prefix = append(prefix, char)
	}

	return string(prefix)
}

func main() {
	fmt.Println("查找字符串最长前缀", findLongestPrefix([]string{"abcdefg", "abcde", "abc"}))
}
