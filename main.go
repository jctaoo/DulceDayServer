package main

import (
	"DulceDayServer/api"
	"DulceDayServer/config"
	_ "DulceDayServer/docs"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	"path"
)

// @title DulceDayServer
// @version 1.0
// @description 一个轻社区app, 内置直播, 帮助用户找到身边的兴趣圈
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /v1
func main() {
	fmt.Println("Starting...")

	// 从命令行参数获取是否为生产模式
	var release = false
	flag.BoolVar(&release, "release", false, "is product environment")
	flag.Parse()

	// 读取配置
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Some Error Occurred When Run.")
	}
	config.ReadConfig(path.Join(wd, "config.toml"), release)

	// gin 的初始化配置
	engine := gin.Default()
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	// api 配置
	v1 := engine.Group("/v1")
	api.SiteEndpoints{
		UserEndpoints: UserEndpoints(),
	}.RouteGroups(v1)

	// api 文档配置
	url := ginSwagger.URL("/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 运行 gin
	err = engine.Run(":" + config.SiteConfig.AppAddress)
	if err != nil {
		fmt.Println("Some Error Occurred When Run.")
	}
}
