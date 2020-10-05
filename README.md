# DulceDayServer
一个轻社区app, 内置直播, 帮助用户找到身边的兴趣圈

## 前置条件
- 安装 docker 以及 docker-compose
- [开发阶段] 运行 deploy/development/deploy_environment.sh
- [开发阶段] 安装 swagger 命令行工具: `go get -u github.com/swaggo/swag/cmd/swag`
- [开发阶段] 安装 wire 命令行工具: `go get github.com/google/wire/cmd/wire`
> 注意：在做完相应改动后应该及时运行 wire 和 swag init，因此建议在运行和提交之前运行这两个命令

## 目录说明
- web: web 站点，目前采用 vue vite 开发，承载网站首页，隐私条款等等
- deploy: 用于部署，开发环境下运行 development/deploy_environment.sh 来运行必要的 docker image， 生产模式下运行 production/docker-compose.yml 直接启动服务
- public: 静态资源
- api: 网络接口，包括 http 与 websocket 接口，不同任务的接口分别为 api 的子 package
- database: 持久化任务，包括持久化数据库与缓存数据库与存放与数据库中的模型
- config: 项目配置，包括 config 模型与读取配置文件（config.toml）等
- e2e: 端对端测试
- locales: 国际化目录