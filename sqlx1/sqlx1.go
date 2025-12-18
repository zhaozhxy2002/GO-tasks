/*
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
  - 要求 ：
  - 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
  - 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int    `db:"id"`
	NAME       string `db:"name"`
	DEPARTMENT string `db:"department"`
	SALARY     int    `db:"salary"`
}

// 查询并储存到Employee[]类型的空切片
func QueryTech(db *sqlx.DB) ([]Employee, error) { //[]Employee是一个切片类型，函数要返回类型
	Emp := []Employee{}
	Query := `SELECT * FROM employees WHERE department = ?`
	err := db.Select(&Emp, Query, "技术部")
	return Emp, err //函数里要返回具体值
}

func QueryTopSalary(db *sqlx.DB) (*Employee, error) {
	emp := Employee{}
	query := `SELECT* FROM employees ORDER BY salary DESC LIMIT 1`
	err := db.Get(&emp, query)
	return &emp, err

}

func main() {
	db, err := sqlx.Open("mysql", "root:ZXYzxy20021113@tcp(127.0.0.1:3306)/mysql") //打开数据库，但不连接
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() //最后关闭连接池

	err = db.Ping() //连接数据库
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接成功！")

	QueryTechEmp, err := QueryTech(db)
	if err != nil {
		log.Fatal("查询技术部员工失败！", err)
	}
	fmt.Println("技术部人员为：")
	for _, result := range QueryTechEmp {
		fmt.Printf("%+v\n", result)
	}

	highest, err := QueryTopSalary(db)
	if err != nil {
		log.Fatal("查询最高工资失败:", err)
	}
	fmt.Println("工资最高的员工：")
	fmt.Printf("%+v\n", highest)
}
