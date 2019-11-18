package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 数据库
)

// DB 数据库实例
var DB *sql.DB

// InitDB 初始化数据库
func InitDB(url string) {
	DB, _ = sql.Open("mysql", url)
	defer DB.Close()
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	// 连接数据库成功
	fmt.Println("connect mysql success")
}
