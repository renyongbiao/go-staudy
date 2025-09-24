package main

import (
	"fmt"
	"math"
)

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	//定义长宽
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

// 矩形的周长
func (s Rectangle) Perimeter() float64 {
	return 2 * (s.length + s.width)
}

// 圆形周长
func (s Circle) Perimeter() float64 {
	return math.Pi * 2 * s.radius
}

// 圆形面积
func (s Circle) Area() float64 {
	return math.Pi * s.radius * s.radius
}

// 矩形的面积
func (s Rectangle) Area() float64 {
	return s.length * s.width
}

func main() {
	// 结构体
	rect := Rectangle{length: 5, width: 3}
	circle := Circle{radius: 4}

	fmt.Println("矩形的周长为：", rect.Perimeter())
	fmt.Println("矩形的面积为：", rect.Area())
	fmt.Println("圆形的周长为：", circle.Perimeter())
	fmt.Println("圆形的面积为：", circle.Area())
}
