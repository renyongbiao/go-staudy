package main

import "fmt"

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
	fmt.Println("删除数组中的重复项，并且递增排序；严格递增意味着重复的元素是相邻的；返回新数组的长度", removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}))
}
