package main

import (
	"gormtest/task"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:1234@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func main() {
	InitDB()

	//基本CRUD操作
	DB.AutoMigrate(&task.Student{})
	// s1 := task.Student{Name: "张三", Age: 20, Grade: "三年级"}
	// task.CreateStudent(DB, &s1)
	// var students []task.Student
	// task.FindStudent(DB, students)
	// task.UpdateStudnet(DB)
	// task.DeleteStudent(DB)

	//事务语句
	DB.AutoMigrate(&task.Account{})
	DB.AutoMigrate(&task.Transaction{})
	// accounts := []task.Account{
	// 	task.Account{Id: 1, Balance: 0},
	// 	task.Account{Id: 2, Balance: 100},
	// }
	// DB.Save(&accounts)
	// fmt.Println(task.Transfer(DB))

	// 使用SQL扩展库进行查询
	DB.AutoMigrate(&task.Employee{})

	// 实现类型安全映射
	DB.AutoMigrate(&task.Book{})

	// 模型定义
	DB.AutoMigrate(&task.User{})
	DB.AutoMigrate(&task.Post{})
	DB.AutoMigrate(&task.Comment{})

	// 关联查询
	// DB.Save(&task.User{Id: 1, PostCounts: 2, Posts: []task.Post{{Id: 1, CommentCounts: 1, Comments: []task.Comment{{Id: 1, Com: "1"}}}, {Id: 2, CommentCounts: 2, Comments: []task.Comment{{Id: 2, Com: "1"}, {Id: 3, Com: "2"}}}}})
	// var post task.Post
	// task.FindMaxCommentsPost(DB, &post)
	// fmt.Println(post)
}
