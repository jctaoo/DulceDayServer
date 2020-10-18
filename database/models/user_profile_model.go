package models

import "gorm.io/gorm"

// UserProfile 用与表现用户信息
type UserProfile struct {
	gorm.Model     `json:"-"`

	// User 的外键，User.UserIdentifier --> UserProfile.UserIdentifier
	UserIdentifier string `gorm:"type:VARCHAR(50);unique" json:"-"`

	Nickname       string `json:"nickname"`
	AvatarFileKey  string `json:"avatar_file_key"`
}

func (u UserProfile) IsEmpty() bool {
	return u == UserProfile{}
}
