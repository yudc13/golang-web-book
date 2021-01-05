package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:yudachao@/book")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	// 设置与数据库建立连接的最大数目
	db.SetMaxOpenConns(10)
	// 设置连接池中最大空闲连接数
	db.SetMaxIdleConns(10)
	Db = db
}
