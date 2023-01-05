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

	// 查询
	sql := "select * from user"
	engine.Query(sql)
	engine.QueryString(sql)
	engine.QueryInterface(sql)

	// SELECT * FROM user LIMIT 1
	user := User{}
	_, _ = engine.Get(&User{})
	fmt.Println(user)

	// SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1
	user = User{Name: "测试一号"}
	_, _ = engine.Where("name = ?", user.Name).Desc("id").Get(&user)
	fmt.Println(user)

	// SELECT name FROM user WHERE id = ?
	var name string
	_, _ = engine.Table(&user).Where("id = ?", 1003).Cols("name").Get(&name)
	fmt.Println(name)

	// SELECT * FROM user WHERE passwd = ? AND age >= 18 limit 10 offset 0
	var users []User
	_ = engine.Where("passwd = ?", 123).And("age >= ?", 18).Limit(10, 0).Find(&users)
	fmt.Println(users)

	// SELECT count(*) AS total FROM user
	counts, _ := engine.Count(&User{})
	fmt.Println(counts)

	// SELECT * FROM user
	_ = engine.Iterate(&User{}, func(idx int, bean interface{}) error {
		user := bean.(*User)
		fmt.Printf("序号%d是%+v\n", idx, user)
		return nil
	})

	// SELECT name, age FROM user
	rows, err := engine.Cols("name", "age").Rows(&User{}) // 传入的&User{}代表搜索条件
	//rows, err := engine.Cols("name", "age").Rows(&User{Name: "测试一号"})
	defer rows.Close()
	for rows.Next() {
		var name string
		var age int
		err = rows.Scan(&name, &age)
		fmt.Println(name, age)
	}

	// SELECT * FROM user where name ="测试一号"
	rows, err = engine.Rows(&User{Name: "测试一号"})
	defer rows.Close()
	for rows.Next() {
		var userBean User
		err = rows.Scan(&userBean)
		fmt.Println(userBean)
	}

}
