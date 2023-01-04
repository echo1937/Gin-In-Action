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

	// xorm.NewEngine
	engine, err := xorm.NewEngine("mysql", dataSourceName)
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

	user1000 := User{
		Id:     1000,
		Name:   "测试一号",
		Passwd: "123",
		Age:    18,
	}

	affected, err := engine.Insert(&user1000)
	fmt.Println(affected)

	user1001 := User{
		Id:     1001,
		Name:   "测试一号",
		Passwd: "123",
		Age:    18,
	}

	user1002 := User{
		Id:     1002,
		Name:   "测试一号",
		Passwd: "123",
		Age:    18,
	}

	affected, err = engine.Insert(&user1001, &user1002)
	fmt.Println(affected)

	// 切片
	var users []User
	users = append(users, User{
		Id:     1003,
		Name:   "测试一号",
		Passwd: "123",
		Age:    18,
	})

	users = append(users, User{
		Id:     1004,
		Name:   "测试一号",
		Passwd: "123",
		Age:    18,
	})

	affected, err = engine.Insert(&users)
	fmt.Println(affected)
}
