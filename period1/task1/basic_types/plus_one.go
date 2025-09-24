package main

import "fmt"

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

func main() {
	fmt.Println("大整数加 1，并返回结果的数字数组。", plusOne([]int{1, 2, 3}))
}
