package main

import (
	_ "database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

var sqlx *sqlx

func main() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(192.168.81.129:3306)/test")
}
