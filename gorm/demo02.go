package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 用户信息
type User struct {
	ID   int64
	Name string `gorm:"default:'小王子'"`
	Age  int
}

func main() {

	var (
		username  string = "root"
		password  string = "redhat"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, ipAddress, port, dbName, charset)
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//var age int8 = 18

	user := User{Name: "只因", Age: 18}
	fmt.Println(db.NewRecord(user))
	db.Debug().Create(&user)
	fmt.Println(db.NewRecord(&user))
}
