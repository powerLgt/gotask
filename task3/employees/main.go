package main

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

建表语句
CREATE TABLE employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    department VARCHAR(50) NOT NULL,
    salary DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Employees struct {
	Id         int        `db:"id"`
	Name       string     `db:"name"`
	Department string     `db:"department"`
	Salary     float32    `db:"salary"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}

func initDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
}

func initData() {
	users := []Employees{
		{Name: "Bob1", Department: "运营部", Salary: 11000},
		{Name: "Bob2", Department: "技术部", Salary: 12000},
		{Name: "Bob3", Department: "人事部", Salary: 13000},
		{Name: "Bob4", Department: "技术部", Salary: 16000},
		{Name: "Bob5", Department: "技术部", Salary: 15000},
	}
	_, err := db.NamedExec(
		"INSERT INTO `employees` (name, department, salary) VALUES (:name, :department, :salary)",
		users)
	if err != nil {
		fmt.Println("数据初始化失败", err)
		panic("数据初始化失败")
	}
}

func cleanData() {
	sql := "delete from `employees`"
	db.Exec(sql)
}

func main() {
	initDB()
	cleanData()
	initData()

	var users []Employees
	var user Employees
	// 使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片
	sql := "SELECT * FROM `employees` WHERE department=?"
	db.Select(&users, sql, "技术部")
	fmt.Println("所有技术部的员工：", users)

	// 使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	sql = "SELECT * FROM `employees` ORDER BY `salary` DESC LIMIT 1"
	db.Get(&user, sql)
	fmt.Println("工资最高的员工：", user)
}
