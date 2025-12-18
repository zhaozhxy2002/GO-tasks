package main

import (
	"fmt"
	"log"
)

func main() {
	dsn := "root:ZXYzxy20021113@tcp(127.0.0.1:3306)/go_test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, created, err := InitDB(dsn)
	if err != nil {
		log.Fatal("连接或迁移数据库失败！", err)
	}

	// 只有当任一表是刚刚创建（首次运行/表不存在）时才插入示例数据
	if created {
		if err := CreateSample(db); err != nil {
			log.Printf("插入示例数据失败：%v", err)
		}
	}

	fmt.Println("=== 查询用户及其评论 ===")
	if err := QueryComments(db, 1); err != nil {
		log.Printf("查询用户文章失败: %v", err)
	}

	fmt.Println("=== 查询评论最多的文章 ===")
	if err := GetMostComments(db); err != nil {
		log.Printf("查询失败: %v", err)
	}

	// 验证钩子函数
	fmt.Println("\n=== 验证钩子函数 ===")
	if err := VerifyHooks(db); err != nil {
		log.Printf("钩子验证失败: %v", err)
	}

}
