package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) []int {
	var index1 int
	var index2 int
	for i := 0; i < len(nums); i++ {
		num3 := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == num3 {
				index1 = i
				index2 = j
				return []int{index1, index2}
			}
		}

	}
	return nil
}

func main() {
	fmt.Println("两数之和", twoSum([]int{2, 7, 11, 15}, 13))
}
