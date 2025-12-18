package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 连接数据库并根据表是否存在决定是否执行 AutoMigrate 与插入测试数据
// 返回值说明: (*gorm.DB, created bool, error)
// - created == true: 说明有任一表之前不存在（刚刚创建），调用方可插入示例数据
// - created == false: 所有表之前均已存在，仅执行了 AutoMigrate
func InitDB(dsn string) (*gorm.DB, bool, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, false, err
	}
	fmt.Println("连接数据库成功！")

	hasUsers := db.Migrator().HasTable(&User{})
	hasPosts := db.Migrator().HasTable(&Post{})
	hasComments := db.Migrator().HasTable(&Comment{})

	// 记录表是否已存在：如果表已存在，我们仍然执行 AutoMigrate
	// 以便新增的字段/列可以被创建；仅当表原本不存在时才插入示例数据。
	allExist := hasUsers && hasPosts && hasComments

	// 执行自动迁移（创建表/新增字段）
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		return db, false, err
	}

	fmt.Println("✅ 表迁移完成（如有缺失的表/列已创建）")

	if allExist {
		fmt.Println("表已存在，已执行 AutoMigrate（新增字段已添加），跳过示例数据插入")
		return db, false, nil
	}

	// 任一表之前不存在——说明刚刚创建表，调用方应插入示例数据
	return db, true, nil
}
