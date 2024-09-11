package _database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dsn = "root:admin123@tcp(127.0.0.1:3306)/bi-test?charset=utf8mb4&parseTime=True&loc=Local"

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// 测试连接
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("mysql 连接成功")
}
