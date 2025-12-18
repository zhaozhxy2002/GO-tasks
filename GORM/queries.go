package main

import (
	"fmt"

	"gorm.io/gorm"
)

// QueryComments 查询某个用户发布的所有文章及其对应的评论信息。
func QueryComments(db *gorm.DB, userid uint) error {
	var user User
	if err := db.Preload("Posts.Comments").First(&user, userid).Error; err != nil {
		return err
	}
	for _, post := range user.Posts {
		fmt.Printf("\n文章标题：%s\n", post.Title)
		fmt.Println("评论：")
		if len(post.Comments) == 0 {
			fmt.Println("（暂无评论）")
		}
		for _, comment := range post.Comments {
			fmt.Println("-", comment.Content)
		}
	}
	return nil
}

// GetMostComments 查询评论最多的文章并打印结果（内部打印，返回 error）
func GetMostComments(db *gorm.DB) error {
	var result PostResult
	if err := db.Model(&Post{}).
		Select("posts.id, posts.title, posts.content, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id, posts.title, posts.content").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error; err != nil {
		return err
	}

	fmt.Printf("文章ID: %d\n", result.ID)
	fmt.Printf("标题: %s\n", result.Title)
	fmt.Printf("内容: %s\n", result.Content)
	fmt.Printf("评论数量: %d\n", result.CommentCount)
	return nil
}
