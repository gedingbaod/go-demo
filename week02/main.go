package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//用户结构体
type Users struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

//数据库指针
var db *sqlx.DB

//初始化数据库连接，init()方法系统会在动在main方法之前执行。
func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(192.168.81.129:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func main() {

	//insertValue()
	//selectValue()
	//updateValue()
	//selectValue()
	//deleteValue()
	tranX()
}
func tranX() {
	//开启事务
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	//执行插入语句
	r, err := conn.Exec("insert into user(username, sex, email)values(?, ?, ?)", "user01", "man", "usre01@163.com")

	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback() //出现异常，进行回滚操作
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into user(username2, sex, email)values(?, ?, ?)", "user02", "man", "user02@163.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()

}
func deleteValue() {
	sql := "delete from user where user_id=?"

	res, err := db.Exec(sql, 4)
	if err != nil {
		fmt.Println("exce failed,", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row failed, ", err)
	}
	fmt.Println("delete succ: ", row)
}
func updateValue() {
	//执行SQL语句
	sql := "update user set username =? where user_id = ?"
	res, err := db.Exec(sql, "user002", 2)

	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	//查询影响的行数，判断修改插入成功
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}

	fmt.Println("update succ:", row)
}
func selectValue() {
	var users []Users
	sql := "select user_id, username,sex,email from user where user_id=? "
	err := db.Select(&users, sql, 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select succ:", users)
}
func insertValue() {
	sql := "insert into user(username,sex, email)values (?,?,?)"
	value := [3]string{"user01", "man", "user01@163.com"}

	//执行SQL语句
	r, err := db.Exec(sql, value[0], value[1], value[2])
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	//查询最后一天用户ID，判断是否插入成功
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("insert succ", id)
}
