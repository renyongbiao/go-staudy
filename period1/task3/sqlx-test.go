package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL驱动
	"github.com/jmoiron/sqlx"
)

// 定义员工结构体
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"` // 修改为string类型
	Salary     float64 `db:"salary"`     // 修改为float64类型以支持小数
}

func selectEmployeeByDeptName(db *sqlx.DB, department string) ([]Employee, error) {

	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", department)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func selectMaxSalary(db *sqlx.DB) (Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees order by salary desc limit 1")
	if err != nil {
		return employee, err
	}
	return employee, nil
}

// 定义book结构体

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  string `db:"price"`
}

func selectBooks(db *sqlx.DB) ([]Book, error) {

	var books []Book
	err := db.Select(&books, "SELECT * FROM books where price > 50")
	if err != nil {
		return nil, err
	}

	return books, nil
}

func main() {
	// 本地数据库连接字符串
	dsn := "test:123456@tcp(106.42.31.5:3306)/go-test?charset=utf8mb4&parseTime=true"

	// 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("数据库连接成功！")

	employees, err := selectEmployeeByDeptName(db, "技术部")
	if err != nil {
		log.Fatalf("查询员工失败: %v", err)
	}

	for _, employee := range employees {
		fmt.Printf("员工信息: %+v\n", employee)
	}

	// 查询工资最高的员工
	employee, err := selectMaxSalary(db)
	if err != nil {
		log.Fatalf("查询最高工资失败: %v", err)
	}
	fmt.Printf("最高工资: %+v", employee)

	books, err := selectBooks(db)
	if err != nil {
		log.Fatalf("查询书籍失败: %v", err)
	}
	for _, book := range books {
		fmt.Printf("书籍信息: %+v\n", book)
	}

}
