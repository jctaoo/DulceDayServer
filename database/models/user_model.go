package models

import "gorm.io/gorm"

// User 用于唯一确定一个用户，区别于用于登录的 AuthUser
type User struct {
	gorm.Model `json:"-"`

	// UserIdentifier 表示这个系统的用户唯一 id，为系统处理用户标记一个随机的字符串
	// 用户可见且全局不可修改
	// 同时作为 user 模块中其他与 User 相关联的模型的外键的引用
	// 与 AuthUser.Identifier 冗余
	UserIdentifier string `sql:"index" gorm:"type:VARCHAR(50),unique" json:"user_identifier"`

	// Username 表示整个系统唯一的用户名，作用与微信号twitter号类似
	// 用户可见且可修改
	// 与 AuthUser.Username 冗余，减少 join 使用并便于
	// 迁移到分布式系统与模块化系统架构
	// 注意冗余字段的相应处理
	Username string `json:"username"`

	// UserProfile 表示用户的信息，如昵称，性别，签名之类的信息
	// 引用 User.UserIdentifier
	// https://gorm.io/zh_CN/docs/has_one.html#%E9%87%8D%E5%86%99%E5%BC%95%E7%94%A8
	UserProfile UserProfile `gorm:"foreignKey:UserIdentifier;references:UserIdentifier"`
}
