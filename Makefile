build: 	## 为当前系统构建可执行文件
	go build -gcflags='-N -l' -o ./build/outer-server-build DulceDayServer

help: ## 显示帮助信息
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

debug: ## 为当前系统构建可执行文件并以调试模式启动
	build \
        && dlv --listen=:2345 --headless=true --api-version=2 exec ./outer-server-build

gen: swagger wire ## 执行所有开发时生成动作

swagger: ## 生成 swagger 文档
	go run github.com/swaggo/swag/cmd/swag init

wire: ## 生成 wire 依赖注入代码
	go run github.com/google/wire/cmd/wire

clear: ## 清除编译或生成的文件及其代码
	rm -Rf ./build \
		&& rm -Rf ./docs \
		&& rm ./wire_gen.go

#graph:
#	go run github.com/99designs/gqlgen