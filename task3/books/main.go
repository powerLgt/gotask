package main

/*

假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
建表语句
	CREATE TABLE books (
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(200) NOT NULL,
		author VARCHAR(100) NOT NULL,
		price DECIMAL(6,2) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

*/
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		panic("数据库连接失败")
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	return
}

func initData() {
	books := []Book{
		{Title: "书籍1", Author: "作者1", Price: 100.0},
		{Title: "书籍2", Author: "作者2", Price: 50.0},
		{Title: "书籍3", Author: "作者3", Price: 25.0},
		{Title: "书籍4", Author: "作者4", Price: 10.0},
	}
	sql := "INSERT INTO `books`(`title`, `author`, `price`) VALUES (:title, :author, :price)"
	_, err := db.NamedExec(sql, books)
	if err != nil {
		fmt.Println("初始化数据失败", err)
	}
}

func clearData() {
	db.Exec("delete from `books`")
}

func main() {
	initDB()
	clearData()
	initData()

	var books []Book
	sql := "SELECT * FROM `books` WHERE `price` > ?"
	db.Select(&books, sql, 50)
	fmt.Println("查询价格大于 50 元的书籍：", books)
}
