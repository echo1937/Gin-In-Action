package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {

	var (
		username  string = "root"
		password  string = "redhat"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, ipAddress, port, dbName, charset)
	fmt.Printf("dataSourceName字符串为: %s\n", dataSourceName)

	var (
		engine *xorm.Engine
		err    error
	)

	// xorm.NewEngine
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	type User struct {
		Id      int64
		Name    string
		Salt    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	// engine.Sync, 同步结构体到数据库
	err = engine.Sync(new(User))
	if err != nil {
		fmt.Println("Sync操作发生异常")
	}

}
