package main

import (
	"time"

	"gorm.io/gorm"
)

// 通用基类：包含主键、时间戳、软删除
type BaseModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 用户
type User struct {
	BaseModel
	Name      string `gorm:"type:varchar(64);not null;index"`
	Email     string `gorm:"type:varchar(128);uniqueIndex"`
	Nickname  string `gorm:"type:varchar(64)"`
	PostCount int    `gorm:"default:0"` // 文章数量统计，后续由 Post 钩子更新

	Posts []Post `gorm:"foreignKey:UserID;references:ID"`
}

// 文章
type Post struct {
	BaseModel
	Title         string    `gorm:"type:varchar(128);not null;index"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null;index"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Comments      []Comment `gorm:"foreignKey:PostID;references:ID"`
	CommentStatus string    `gorm:"type:varchar(32);default:'有评论'"`
}

// 评论
type Comment struct {
	BaseModel
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"not null;index"`
}

// 查询结果结构体
type PostResult struct {
	ID           uint
	Title        string
	Content      string
	CommentCount int64
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段，更新对应 User.PostCount
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {

	// 如果文章没有关联用户，直接返回
	if p.UserID == 0 {
		return nil
	}
	// 更新用户的文章数量：PostCount + 1
	err = tx.Model(&User{}).
		Where("id=?", p.UserID).
		Update("post_count", gorm.Expr("post_count+1")).Error
	return err
}

// - 为 Comment 添加 AfterDelete 钩子，在删除评论后检查对应 Post 的评论数量并更新 Post.CommentStatus
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	if c.PostID == 0 {
		return nil
	}
	var count int64
	//统计该文章剩余评论数
	err = tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error
	if err != nil {
		return err
	}
	//根据评论数量更新文章状态
	status := "有评论"
	if count == 0 {
		status = "无评论"
	}
	//更新文章字段
	err = tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_status", status).Error

	return err
}
