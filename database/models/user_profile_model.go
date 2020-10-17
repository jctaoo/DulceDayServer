package models

import "gorm.io/gorm"

// 用于表现用户信息的，区别于用于登录的 User
type UserProfile struct {
	gorm.Model `json:"-"`
	Uid string `json:"uid"`
	UserIdentifier string `json:"-"`
	Nickname string `json:"nickname"`
	AvatarFileKey string `json:"avatar_file_key"`
}

func (u UserProfile) IsEmpty() bool {
	return u == UserProfile{}
}
