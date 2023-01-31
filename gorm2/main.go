package main

import (
	"Gin-In-Action/gorm2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
)

var (
	db *gorm.DB
)

func initDB() (err error) {

	var (
		username  string = "root"
		password  string = "redhat"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, ipAddress, port, dbName, charset)

	// https://stackoverflow.com/questions/44589060/how-to-set-singular-name-for-a-table-in-gorm
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return err
	}

	// https://gorm.io/zh_CN/docs/generic_interface.html
	sqlDB, err := db.DB()
	return sqlDB.Ping()
}

func main() {

	// 连接数据库
	err := initDB()
	if err != nil {
		panic(err)
	}

	// 模型绑定
	db.AutoMigrate(&models.Todo{})

	ginServer := gin.Default()
	ginServer.LoadHTMLGlob("gorm2/templates/*")
	ginServer.Static("static", "gorm2/static")

	ginServer.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	// 待办事项
	v1Group := ginServer.Group("/v1")
	{
		// 添加
		v1Group.POST("/todo", func(context *gin.Context) {
			// 前端页面填写待办事项 点击提交 会发请求到这里
			// 1. 从请求中把数据拿出来
			var todo models.Todo
			context.BindJSON(&todo)
			// 2. 存入数据库
			if err = db.Create(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(context *gin.Context) {
			// 查询todo这个表里的所有数据
			var todoList []models.Todo
			if err = db.Find(&todoList).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, todoList)
			}
		})
		// 查看
		v1Group.GET("/todo/:id", func(context *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(context *gin.Context) {
			id := context.Param("id")
			var todo models.Todo
			if err = db.Where("id=?", id).First(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			context.BindJSON(&todo)
			if err = db.Save(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, &todo)
			}
		})
		// 删除一个待办事项
		v1Group.DELETE("/todo/:id", func(context *gin.Context) {
			id := context.Param("id")

			if err = db.Where("id=?", id).Delete(models.Todo{}).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{id: "deleted"})
			}

		})
	}
	ginServer.Run()

}
