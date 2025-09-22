package main

import "fmt"

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

func singleNumber(nums []int) {

	var map1 = make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if _, exists := map1[nums[i]]; exists {
			map1[nums[i]]++
		} else {
			map1[nums[i]] = 1
		}

	}

	for k, v := range map1 {
		if v == 1 {
			fmt.Println("出现一次的元素为：", k)
		}
	}

}

// 判断是否回文数
func isPalindrome(x int) bool {

	if x <= 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var num = 0

	for x > num {

		num = num*10 + x%10
		x /= 10
	}

	return num == x || num/10 == x

}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValidStr(str string) bool {

	arrStr := []byte(str)

	var stack []byte

	for _, c := range arrStr {

		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) == 0 || stack[len(stack)-1] == c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0

}

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

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
// plusOne 将 digits 表示的大整数加 1，返回新的数字切片
func plusOne(digits []int) []int {
	n := len(digits)
	// 从最后一位向前遍历
	for i := n - 1; i >= 0; i-- {
		digits[i]++ // 先加 1
		if digits[i] < 10 {
			return digits // 无进位，直接返回
		}
		digits[i] = 0 // 有进位，当前位归零，继续循环处理前一位
	}
	// 跳出循环说明所有位都是 0（原数为 999...），需要在最前面补 1
	return append([]int{1}, digits...)
}

// 删除数组中的重复项，并且递增排序；严格递增意味着重复的元素是相邻的；返回新数组的长度
func removeDuplicates(nums []int) int {

	k := 1 //第一个元素保留

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k

}

func main() {
	singleNumber([]int{2, 2, 1, 3, 3}) // 出现一次的元素
	fmt.Println("判断是否回文数：", isPalindrome(1211))
	fmt.Println("判断是否有效字符串", isValidStr("()[]{}"))
	fmt.Println("查找字符串最长前缀", findLongestPrefix([]string{"abcdefg", "abcde", "abc"}))
	fmt.Println("大整数加 1，并返回结果的数字数组。", plusOne([]int{1, 2, 3}))
	fmt.Println("删除数组中的重复项，并且递增排序；严格递增意味着重复的元素是相邻的；返回新数组的长度", removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}))
}
