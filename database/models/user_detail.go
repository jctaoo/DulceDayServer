package models

import "gorm.io/gorm"

// 用于表现用户信息的，区别于用于登录的 User
type UserDetail struct {
	gorm.Model `json:"-"`


}
