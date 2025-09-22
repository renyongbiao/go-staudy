package main

import "fmt"

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

func singleNumber(nums []int)  {

	var map1 = make(map[int]int)

	for i:=0 ;i<len(nums);i++{

		if _,exists := map1[nums[i]];exists {
			map1[nums[i]]++
		}else{
			map1[nums[i]] = 1
		}

	}

	for k,v := range map1{
		if v == 1{
			fmt.Println("出现一次的元素为：",k)
		}
	}

}

// 判断是否回文数
func isPalindrome(x int) bool {

	if x<=0 || (x%10 == 0 && x!=0){
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

	for _,c := range arrStr  {

		if c == '(' || c == '[' || c == '{'{
			stack = append(stack,c)
		}else if len(stack) == 0 || stack[len(stack)-1] == c{
			return false
		}else{
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

    var prefix []byte  // 使用字节切片存储字符

    // 逐位比较
    for i := 0; i < minLen; i++ {
        char := strs[0][i]  // 第一个字符串的第i个字符

        // 检查所有字符串的第i位是否相同
        allMatch := true
        for j := 1; j < len(strs); j++ {
            if strs[j][i] != char {
                allMatch = false
                break
            }
        }

        if !allMatch {
            break  // 发现不匹配，停止
        }

        // 所有字符串的第i位都相同，添加到前缀
        prefix = append(prefix, char)
    }

    return string(prefix)
}


func main() {
	singleNumber([]int{2, 2, 1,3,3})  // 出现一次的元素
	fmt.Println("判断是否回文数：",isPalindrome(1211))
	fmt.Println("判断是否有效字符串",isValidStr("()[]{}"))
	fmt.Println("查找字符串最长前缀",findLongestPrefix([]string{"abcdefg","abcde","abc"}))
}
