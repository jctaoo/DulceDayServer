package models

import (
	"gorm.io/gorm"
)

// 用于登录的 User, 区别于用于表现用户信息的 UserProfile
type User struct {
	gorm.Model `json:"-"`

	Identifier string `json:"-"`
	Username string
	Password string `json:"-"`
	Email string `json:"-"`
	Tokens []TokenAuth `json:"-"`
	Authority  AuthorityLevel `json:"-"`
}

func (u User) Validate() bool {
	return u.Username != ""
}

func (u User) ValidateNew() bool {
	return u.Username != "" && u.Password != ""
}

func (u User) IsEmpty() bool {
	return u.Identifier == "" && u.Validate()
}