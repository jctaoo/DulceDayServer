package main

import (
	"DulceDayServer/api"
	"DulceDayServer/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

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

	// 运行 gin
	err = engine.Run(":" + config.SiteConfig.AppAddress)
	if err != nil {
		fmt.Println("Some Error Occurred When Run.")
	}
}
