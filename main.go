package main

import (
	"DulceDayServer/api"
	"DulceDayServer/api/common"
	"DulceDayServer/config"
	_ "DulceDayServer/docs"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /v1
func main() {
	fmt.Println("starting...")

	// 从命令行参数获取是否为生产模式
	var release = false
	flag.BoolVar(&release, "release", false, "is product environment")
	flag.Parse()

	// 读取配置
	wd, err := os.Getwd()
	if err != nil {
		log.Error("获取运行目录失败")
		return
	}
	config.ReadConfigOrExit(path.Join(wd, "config.toml"), release)

	// 日志配置
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05", ForceColors: true})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout)) // todo log to file in release
	log.SetLevel(log.PanicLevel)
	log.SetReportCaller(true)

	// gin 的初始化配置
	engine := gin.Default()
	if release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// gin 通用中间件配置
	//engine.Use(common.MiddleWareLog())

	// 配置国际化
	common.ValidatorTransInit()

	// api 配置
	v1 := engine.Group("/v1")
	api.SiteEndpoints{
		UserEndpoints: UserEndpoints(),
		UserProfileEndpoints: UserProfileEndpoints(),
		StaticStorageEndpoints: StaticStorageEndpoints(),
		MomentEndpoints: MomentEndpoints(),
		StoreEndpoints: StoreEndpoints(),
	}.RouteGroups(v1)

	// api 文档配置
	url := ginSwagger.URL("/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 运行 gin
	err = engine.Run(":" + config.SiteConfig.AppAddress)
	if err != nil {
		log.Error("运行失败")
		return
	}
}
