package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Person 用户信息
type Person struct {
	gorm.Model
	Name string
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

	db.AutoMigrate(&Person{})

	p1 := Person{Name: "七米", Age: 18}
	p2 := Person{Name: "沙河娜扎", Age: 29}
	db.Create(&p1)
	db.Create(&p2)

	var people []Person
	db.Debug().Find(&people)
	fmt.Printf("users:%#v\n", people)

}
