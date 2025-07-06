package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*

题目1：模型定义
	假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	要求 ：
	使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
	基于上述博客系统的模型定义。
	要求 ：
	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
	继续使用博客系统的模型。
	要求 ：
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

var db *gorm.DB

type User struct {
	gorm.Model
	Id        int `gorm:"primaryKey;autoIncrement"`
	Username  string
	PostCount int
	Posts     []Post    `gorm:"foreignKey:UserId"`
	Comments  []Comment `gorm:"foreignKey:UserId"`
}

type Post struct {
	gorm.Model
	Id            int `gorm:"primaryKey;autoIncrement"`
	UserId        int
	Title         string
	Content       string
	CommentCount  int
	CommentStatus int
	Comments      []Comment `gorm:"foreignKey:PostId"`
	User          User      `gorm:"foreignKey:UserId"`
}

type Comment struct {
	gorm.Model
	Id      int `gorm:"primaryKey;autoIncrement"`
	PostId  int
	UserId  int
	Content string
	Post    Post `gorm:"foreignKey:PostId;references:Id"`
	User    User `gorm:"foreignKey:UserId;references:Id"`
}

func (u *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Post{}).Where("user_id = ?", u.UserId).Count(&count)
	tx.Model(&User{Id: u.UserId}).Update("post_count", count)
	fmt.Println("更新用户文章数", u.UserId, count)
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if c.PostId > 0 {
		tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
		tx.Model(&Post{Id: c.PostId}).Update("comment_count", count)
		if count == 0 {
			tx.Model(&Post{Id: c.PostId}).Update("comment_status", 1)
		}
		fmt.Println("更新文章的评论数", c.PostId, count)
	}
	return
}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		panic("连接数据库失败")
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("表结构初始化失败", err)
		panic("表结构初始化失败")
	}

	return nil
}

func initData() {
	users := []User{
		{
			Username:  "user1",
			PostCount: 2,
			Posts: []Post{
				{
					Title:         "xxxx1",
					Content:       "xxxxx1",
					CommentCount:  2,
					CommentStatus: 2,
					Comments: []Comment{
						{Content: "xxxxx1"},
						{Content: "xxxxx2"},
					},
				},
				{
					Title:         "xxxx2",
					Content:       "xxxxx2",
					CommentCount:  2,
					CommentStatus: 2,
					Comments: []Comment{
						{Content: "xxxxx1"},
						{Content: "xxxxx2"},
					},
				},
			},
		},
		{
			Username:  "user2",
			PostCount: 3,
			Posts: []Post{
				{
					Title:         "xxxx1",
					Content:       "xxxxx1",
					CommentCount:  3,
					CommentStatus: 2,
					Comments: []Comment{
						{Content: "xxxxx1"},
						{Content: "xxxxx2"},
						{Content: "xxxxx3"},
					},
				},
				{
					Title:         "xxxx2",
					Content:       "xxxxx2",
					CommentCount:  3,
					CommentStatus: 2,
					Comments: []Comment{
						{Content: "xxxxx1"},
						{Content: "xxxxx2"},
						{Content: "xxxxx3"},
					},
				},
				{
					Title:         "xxxx3",
					Content:       "xxxxx3",
					CommentCount:  3,
					CommentStatus: 2,
					Comments: []Comment{
						{Content: "xxxxx1"},
						{Content: "xxxxx2"},
						{Content: "xxxxx3"},
					},
				},
			},
		},
	}
	db.Create(&users)
}

func clearData() {
	db.Delete(&User{}, "id > 0")
	db.Delete(&Post{}, "id > 0")
	db.Delete(&Comment{}, "id > 0")
}

func main() {
	initDB()
	clearData()
	initData()

	//查询某个用户发布的所有文章及其对应的评论信息
	var user User
	db.Preload("Posts.Comments").Preload(clause.Associations).Find(&user, "username = ?", "user2")
	for i := 0; i < len(user.Posts); i++ {
		fmt.Println("user2发布文章: ", user.Posts[i].Content)
		for j := 0; j < len(user.Posts[i].Comments); j++ {
			fmt.Println("- 文章评论: ", user.Posts[i].Comments[j].Content)
		}
	}

	// 查询评论数量最多的文章信息
	var topPost Post
	db.Model(&Post{}).
		Select("*").
		Order("comment_count DESC").
		First(&topPost)
	fmt.Println("评论数最多的文章是", topPost)
}
