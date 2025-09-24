package main

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
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

func main() {
	fmt.Println("判断是否有效字符串", isValidStr("()[]{}"))
}
