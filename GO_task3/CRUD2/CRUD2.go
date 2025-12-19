/*
题目2：事务语句 - 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
- 要求 ： - 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Transfer(db *sql.DB, from_ID string, to_ID string, amount float64) error {
	//开启事务，想要A给B
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	//一旦函数中出现错误，事务整个回滚
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var balance float64
	//查询A的账户余额,并scan赋值给balance
	QuerySql := `SELECT balance FROM accounts WHERE id = ?`
	row := tx.QueryRow(QuerySql, from_ID)
	err = row.Scan(&balance)
	if err != nil {
		return err
	}
	//检查余额是否大于amounts
	if balance < amount {
		return fmt.Errorf("余额不足，当前余额 %.2f，转账金额 %.2f", balance, amount)
	}
	//开始转账，给A扣款
	DeductSql := `UPDATE accounts SET balance=balance-? WHERE id=?`
	_, err = tx.Exec(DeductSql, amount, from_ID)
	if err != nil {
		return err
	}
	//给B账户增加余额
	AddSql := `UPDATE accounts SET balance=balance+? WHERE id=?`
	_, err = tx.Exec(AddSql, amount, to_ID)
	if err != nil {
		return err
	}
	//把转账记录到transaction表中，录入数据库
	TranSql := `INSERT INTO transactions (from_account_id,to_account_id,amount) VALUES(?,?,?)`
	_, err = tx.Exec(TranSql, from_ID, to_ID, amount)

	if err != nil {
		return err
	}
	return nil //到这里 err == nil，defer 会自动 Commit

}

func main() {
	dsn := "root:ZXYzxy20021113@tcp(127.0.0.1:3306)/mysql"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() //最后关闭连接池

	err = db.Ping()
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接成功！")

	//调用转账函数
	err = Transfer(db, "A", "B", 800)
	if err != nil {
		log.Println("转账失败！", err)
	} else {
		log.Println("转账成功！")
	}
}
