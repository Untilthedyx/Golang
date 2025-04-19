package task

import (
	"errors"

	"gorm.io/gorm"
)

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	Id    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Age   int
	Grade string
}

func CreateStudent(db *gorm.DB, student *Student) error {
	result := db.Debug().Table("students").Create(student)
	return result.Error
}
func FindStudent(db *gorm.DB, students []Student) error {
	result := db.Debug().Table("students").Where("Age>?", 18).Find(&students)
	return result.Error
}
func UpdateStudnet(db *gorm.DB) error {
	result := db.Debug().Table("students").Where("Name=?", "张三").Update("Grade", "四年级")
	return result.Error
}
func DeleteStudent(db *gorm.DB) error {
	result := db.Debug().Table("students").Where("Age<?", 15).Delete(&Student{})
	return result.Error
}

// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
// 如果余额不足，则回滚事务。

type Account struct {
	Id      uint
	Balance float64
}
type Transaction struct {
	Id            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func Transfer(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	var A, B Account = Account{}, Account{}
	if err := tx.Model(&Account{}).Find(&A, 1).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&Account{}).Find(&B, 2).Error; err != nil {
		tx.Rollback()
		return err
	}
	if A.Balance < 100 {
		tx.Rollback()
		return errors.New("余额不足")
	} else {
		A.Balance -= 100
		B.Balance += 100
		if err := tx.Save(&A).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Save(&B).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Create(&Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 100}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
