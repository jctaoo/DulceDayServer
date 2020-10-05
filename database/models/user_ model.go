package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Identifier string
	Username string
	Password string
	Email string
	Tokens []TokenAuth
	Authority  AuthorityLevel
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