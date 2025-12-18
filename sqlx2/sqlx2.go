/*
实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func QueryPrice(db *sqlx.DB) ([]Book, error) { //[]Employee是一个切片类型，函数要返回类型
	books := []Book{}
	Query := `SELECT id, title, author, price FROM books WHERE price > ?`
	err := db.Select(&books, Query, 50)
	return books, err //函数里要返回具体值
}

func main() {
	db, err := sqlx.Open("mysql", "root:ZXYzxy20021113@tcp(127.0.0.1:3306)/go_test_db") //打开数据库，但不连接
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() //最后关闭连接池

	err = db.Ping() //连接数据库
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接成功！")

	QueryPriceBook, err := QueryPrice(db)
	if err != nil {
		log.Fatal("查询书籍失败！", err)
	}
	fmt.Println("价格大于 50 元的书籍：")
	for _, result := range QueryPriceBook {
		fmt.Printf("%+v\n", result)
	}

}
