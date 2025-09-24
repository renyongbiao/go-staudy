package main

import "fmt"

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func addTen(num *int) int {
	*num += 10
	return *num
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiplyByTwo(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
}

func main() {

	//指针
	num := 15
	fmt.Println("传入的值为：", num)
	fmt.Println("修改后的值为：", addTen(&num))

	//切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片为：", slice)
	multiplyByTwo(&slice)
	fmt.Println("修改后的切片为：", slice)
}
