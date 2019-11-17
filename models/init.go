package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Db数据库连接池
var DB *sql.DB

func InitDB(mysql_dsn string) {
	DB, _ = sql.Open("mysql", mysql_dsn)
	defer DB.Close()
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	// 连接数据库成功
	fmt.Println("connect mysql success")
}
