package model

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

// DB mysql 应该是引用类型
var DB = &sql.DB{}

// Redis redis 处理
var Redis = &redis.Client{}

// Init 数据库初始化链接
func Init() {
	var err error
	DB, err = sql.Open("mysql", "root:liqian521@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 地址
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
}
