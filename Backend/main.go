package main

import (
	"NetworkControl/app/controller"
	"NetworkControl/app/mapper"
	"NetworkControl/app/middleware"
	"NetworkControl/config"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
)

func main() {

	err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}

	err = mapper.Init(viper.GetString("source"))
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	e.Use(middleware.Authenticate)

	//  编写一个接受请求后输出pong!的 handler
	e.GET("/pong", controller.Pong)

	// Print 第一个接口 接受一个query参数，直接print这个参数
	e.GET("/print/query", controller.Print)

	// Analybody 第二个接口 解析request body并print
	e.POST("/print/body", controller.Analybody)

	// AddTodo 添加todo接口
	e.POST("/api/todo/add", controller.AddTodo)

	// GetTodo 查询todo接口
	e.GET("/api/todo/get", controller.GetTodo)

	// 启动http server
	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
