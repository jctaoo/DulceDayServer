package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`

	Identifier string `json:"identifier"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email string `json:"email"`
	Tokens []TokenAuth
}

func (u User) Validate() bool {
	return u.Username != ""
}

func (u User) ValidateNew() bool {
	return u.Username != "" && u.Password != ""
}