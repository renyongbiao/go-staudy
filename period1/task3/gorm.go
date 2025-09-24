package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义用户结构体
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:32;not null"`
	PostCount int    `gorm:"size:8"`                                        //文章总数
	Posts     []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // 一对多
}

// Post 文章：一对多
type Post struct {
	ID       uint      `gorm:"primaryKey"`
	Title    string    `gorm:"size:128;not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   uint      `gorm:"index;not null"`                                // 外键
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` // 一对多
}

// Comment 评论
type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"index;not null"` // 外键
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量
	err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error

	log.Println("文章创建后回调函数执行", err)

	return err
}

// 查询某个用户发布的所有文章以及文章评论
func selectCommentAndPostByUser(userId uint, db *gorm.DB) []Post {
	var posts []Post
	err := db.Where("user_id = ?", userId).Preload("Comments").Find(&posts).Error
	if err != nil {
		return nil
	}
	return posts

}

// 查询评论最多的文章
func selectMostCommentedPost(db *gorm.DB) Post {
	var post Post
	err := db.Where("id = (select post_id from comments group by post_id order by count(*) desc limit 1)").Find(&post).Error
	if err != nil {
		return Post{} // 返回空的Post结构体
	}
	return post
}

func main() {

	// 1. 连接 MySQL
	dsn := "test:123456@tcp(106.42.31.5:3306)/go-test?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

	// 查询用户ID为1的文章
	posts := selectCommentAndPostByUser(1, db)
	for _, post := range posts {
		fmt.Printf("文章信息: %+v\n", post)
	}

	// 查询评论最多的文章
	post := selectMostCommentedPost(db)
	fmt.Printf("评论最多的文章: %+v\n", post)

	// 为用户1添加文章
	post = Post{Title: "gorm测试", Content: "gorm测试内容", UserID: 1}
	db.Create(&post)

}
