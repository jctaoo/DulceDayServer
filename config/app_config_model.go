package config

type Config struct {
	// 应用名字
	AppName string
	// 应用运行的地址
	AppAddress string
	// token 过期时间， 单位为
	AuthTokenExpiresTime int64
	// 用于生成鉴权 token 的 secret
	AuthTokenSecret string
	// 数据库设置
	DataBaseConfig DataBaseConfig
	// 缓存数据库设置
	CacheConfig CacheConfig
}
