package task

import "gorm.io/gorm"

// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
type Employee struct {
	Id         uint
	Name       string
	Department string
	Salary     float64
}

func FindEmpoloyee(db *gorm.DB, employees []Employee) error {
	result := db.Table("employees").Where("department=?", "技术部").Find(&employees)
	return result.Error
}

func FindMaxSalary(db *gorm.DB, employee *Employee) error {
	result := db.Table("employees").Order("salary desc").First(employee)
	return result.Error
}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	Id     uint
	Title  string
	Author string
	Price  float64
}

func FindBook(db *gorm.DB, books []Book) error {
	result := db.Table("books").Where("price>?", 50).Find(&books)
	return result.Error
}
