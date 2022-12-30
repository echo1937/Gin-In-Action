package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

func main() {
	// 创建一个服务
	ginServer := gin.Default()

	//	当前的working_directory为GO-In-Action

	// 1、加载html页面
	ginServer.LoadHTMLGlob("gin/templates/*")

	// 2、加载静态资源，relativePath指url的地址，root指对应文件系统的地址
	// 访问localhost:8082/static/{file}可获取${working_directory}/gin/static目录下的file文件
	ginServer.Static("./static", "gin/static")

	// 3、404页面设置
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// GET方法
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})

	// GET方法(取得查询参数): user/info?userid=1&username=tom
	ginServer.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")

		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username})
	})

	// GET方法(取得PathVariable): user/info/{userid}/{username}
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")

		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username})
	})

	// GET方法(返回HTML页面)
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "来自后台的程序",
		})
	})

	// GET方法(中间件,即拦截器)
	// ginServer.Use(myHandler())  // 全局使用
	ginServer.GET("/handler", myHandler(), func(context *gin.Context) {

		value := context.MustGet("session")
		context.JSON(http.StatusOK, gin.H{
			"session": value,
			"message": "OK",
		})
	})

	ginServer.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.qq.com")
	})

	// POST方法
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "post  -user"})
	})

	// POST方法(反序列化): {"message":"hello"}
	ginServer.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]any // key为string,value为any
		json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	// POST方法(接收form表单)
	ginServer.POST("/user/add", func(context *gin.Context) {
		username, _ := context.GetPostForm("username")
		password, _ := context.GetPostForm("password")

		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由组
	group := ginServer.Group("/group")
	{
		group.GET("/add", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "这是路由组",
				"path":    "/group/add",
			})
		})
		group.GET("/show", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "这是路由组",
				"path":    "/group/show",
			})
		})
	}

	ginServer.Run(":8082")
}

func myHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.Set("session", "这是通过myHandler注入的")
		//context.Next()
		//context.Abort()
	}
}
