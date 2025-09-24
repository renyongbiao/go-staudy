package main

import "fmt"

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

func main() {
	fmt.Println("判断是否回文数：", isPalindrome(1211))
}
