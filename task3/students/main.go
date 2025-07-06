package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Students struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

建表语句：
CREATE TABLE students (

	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	age INT NOT NULL,
	grade VARCHAR(20) NOT NULL

);
*/
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	student := Students{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	db.Create(&student)

	var students []Students
	db.Where("age > ?", 18).Find(&students)
	fmt.Println("所有大于18岁的学生: ", students)

	result := db.Model(&student).Update("grade", "四年级")
	fmt.Println("更新成功, count: ", result.RowsAffected)

	students = []Students{}
	db.Where("grade = ?", "四年级").Find(&students)
	fmt.Println("所有四年级的学生: ", students)

	result = db.Where("age < ?", 15).Delete(&Students{})
	fmt.Println("删除成功, count: ", result.RowsAffected)

	result = db.Where("name = ?", "张三").Delete(&Students{})
	fmt.Println("删除所有张三, count: ", result.RowsAffected)
}
