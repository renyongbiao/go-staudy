package main

import "fmt"

// 定义父结构体：人
type Person struct {
	name string
	age  int
}

// 定义员工结构体
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) printInfo() {
	fmt.Printf("姓名：%s，年龄：%d，员工编号：%d\n", e.name, e.age, e.EmployeeID)
}

func main() {
	employee := Employee{Person: Person{name: "张三", age: 30}, EmployeeID: 12345}
	employee.printInfo()
}
