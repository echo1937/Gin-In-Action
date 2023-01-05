package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
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

	session := engine.NewSession()
	defer session.Close()

	session.Begin()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	// 这条sql总是能成功的
	rand.Seed(time.Now().UnixNano())
	primary := int64(rand.Intn(1000) + 1000)
	fmt.Println(primary)
	if _, err := session.Insert(&User{Id: primary, Name: "测试一号", Passwd: "123", Age: 18}); err != nil {
		panic(err)
	}

	// 这条sql总是会失败的(1000是重复的)
	if _, err := session.Insert(&User{Id: 1000, Name: "测试一号", Passwd: "123", Age: 18}); err != nil {
		panic(err)
	}

}
