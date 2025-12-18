
//err使用有大问题！！！！！！

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type STUDENT struct {
	ID    int
	NAME  string
	AGE   int
	GRADE string
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func InsertRow(db *sql.DB) error {
	InsertSql := `INSERT INTO students(name,age,grade) VALUES(?,?,?)`
	result, err := db.Exec(InsertSql, "张三", 20, "三年级")
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	log.Println("插入数据成功，新学生id：", id)
	return nil //表示无错误发生
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func QueryRow(db *sql.DB) error {
	QuerySql := `SELECT * FROM students WHERE age>?`
	rows, err := db.Query(QuerySql, 18)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var stu STUDENT //数据库操作最好用局部变量
		err = rows.Scan(&stu.ID, &stu.NAME, &stu.AGE, &stu.GRADE)
		if err != nil {
			return err
		}
		log.Printf("查询到的学生ID为%v,姓名为%v，年龄%v,年级%v\n", stu.ID, stu.NAME, stu.AGE, stu.GRADE)
	}
	return rows.Err()
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func UpdateRow(db *sql.DB) error {
	UpdateSql := `UPDATE students SET grade = ? WHERE name=? `
	result, err := db.Exec(UpdateSql, "四年级", "张三")
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		log.Println("更新失败！")
	} else {
		log.Println("更新成功！变更行数为：", rows)
	}
	return nil
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func DeleteRow(db *sql.DB) error {
	DeleteSql := `DELETE FROM students WHERE age<?`
	result, err := db.Exec(DeleteSql, 15)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		log.Println("删除失败！")
	} else {
		log.Println("删除成功！删除行数为：", rows)
	}
	return nil //表示无错误发生
}

func main() {
	db, err := sql.Open("mysql", "root:ZXYzxy20021113@tcp(127.0.0.1:3306)/mysql") //打开数据库，但不连接
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() //最后关闭连接池

	err = db.Ping() //连接数据库
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接成功！")

	// 按顺序调用CRUD
	if err = InsertRow(db); err != nil {
		log.Println("插入失败：", err)
	}

	if err = QueryRow(db); err != nil {
		log.Println("查询失败：", err)
	}

	if err = UpdateRow(db); err != nil {
		log.Println("更新失败：", err)
	}

	if err = DeleteRow(db); err != nil {
		log.Println("删除失败：", err)
	}

}
