package config

import "fmt"

// 持久化数据库配置
type DataBaseConfig struct {
	Host string
	Port string
	// collection (for nosql) or table (for sql)
	Collection string
	User       string
	Password   string
	QueryParameters string
}

// 缓存数据库配置
type CacheConfig struct {
	Host string
	Port string
	DB int

	// 用户 ID 黑名单列表名字
	BlackListName,
	// 撤销 Token 列表名字
	RevokeTokenListName,
	// IP 黑名单名字
	IPBlackListName,
	// 未激活 Token 列表名字
	InActiveTokenListName string
}

// like `user=gorm password=gorm dbname=gorm port=9920`
// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
func (config DataBaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Collection,
		config.QueryParameters,
	)
}
