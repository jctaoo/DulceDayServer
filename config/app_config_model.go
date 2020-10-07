package config

type Config struct {
	// 应用名字
	AppName string
	// 应用运行的地址
	AppAddress string
	// token 过期时间， 单位为秒
	AuthTokenExpiresTime int64
	// 验证 token 过期时间
	VerificationTokenExpiresTime int64
	// 用于生成鉴权 token 的 secret
	AuthTokenSecret string
	// 数据库设置
	DataBaseConfig DataBaseConfig
	// 缓存数据库设置
	CacheConfig CacheConfig
	// 默认设备名
	DefaultDeviceName string
	// 头像大小限制
	AvatarSizeMB int
	// 静态存储配置
	AliOssStaticStorageConfig AliOssStaticStorageConfig
}
