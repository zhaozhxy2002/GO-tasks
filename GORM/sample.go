package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// CreateSample 插入示例数据（只做示例用途）
func CreateSample(db *gorm.DB) error {
	user := User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Nickname: "Ali",
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	post1 := Post{Title: "Go 语言入门", Content: "第一篇文章内容……", UserID: user.ID}
	if err := db.Create(&post1).Error; err != nil {
		log.Fatal("创建文章1失败！", err)
	}
	post2 := Post{Title: "Gorm 实战", Content: "第二篇文章内容……", UserID: user.ID}
	if err := db.Create(&post2).Error; err != nil {
		log.Fatal("创建文章2失败！", err)
	}

	comments := []Comment{
		{Content: "写得不错", PostID: post1.ID},
		{Content: "继续加油", PostID: post1.ID},
		{Content: "有帮助，谢谢", PostID: post2.ID},
	}
	if err := db.Create(&comments).Error; err != nil {
		log.Fatal("创建评论失败！", err)
	}

	fmt.Println("测试数据插入成功！")
	return nil
}

// VerifyHooks 验证钩子函数是否正常工作
func VerifyHooks(db *gorm.DB) error {
	fmt.Println("=== 开始验证钩子函数 ===")

	// 步骤1: 创建一个新用户，初始 PostCount 应为 0
	user := User{
		Name:     "Bob",
		Email:    "bob@example.com",
		Nickname: "Bobby",
	}
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}
	fmt.Printf("创建用户 Bob，初始 PostCount: %d\n", user.PostCount) // 应为 0

	// 步骤2: 为用户创建一篇文章，钩子应更新 PostCount 为 1
	post := Post{
		Title:   "测试文章",
		Content: "这是用于测试钩子的文章",
		UserID:  user.ID,
	}
	if err := db.Create(&post).Error; err != nil {
		return fmt.Errorf("创建文章失败: %v", err)
	}

	// 重新查询用户，检查 PostCount 是否更新
	var updatedUser User
	if err := db.First(&updatedUser, user.ID).Error; err != nil {
		return fmt.Errorf("查询用户失败: %v", err)
	}
	fmt.Printf("创建文章后，用户的 PostCount: %d (应为 1)\n", updatedUser.PostCount)

	// 步骤3: 为文章添加一个评论
	comment := Comment{
		Content: "测试评论",
		PostID:  post.ID,
	}
	if err := db.Create(&comment).Error; err != nil {
		return fmt.Errorf("创建评论失败: %v", err)
	}

	// 重新查询文章，检查 CommentStatus
	var updatedPost Post
	if err := db.First(&updatedPost, post.ID).Error; err != nil {
		return fmt.Errorf("查询文章失败: %v", err)
	}
	fmt.Printf("添加评论后，文章的 CommentStatus: %s (应为 '有评论')\n", updatedPost.CommentStatus)

	// 步骤4: 删除评论，钩子应更新 CommentStatus 为 '无评论'
	if err := db.Delete(&comment).Error; err != nil {
		return fmt.Errorf("删除评论失败: %v", err)
	}

	// 重新查询文章，检查 CommentStatus 是否更新
	if err := db.First(&updatedPost, post.ID).Error; err != nil {
		return fmt.Errorf("查询文章失败: %v", err)
	}
	fmt.Printf("删除评论后，文章的 CommentStatus: %s (应为 '无评论')\n", updatedPost.CommentStatus)

	fmt.Println("=== 钩子验证完成 ===")
	return nil
}
