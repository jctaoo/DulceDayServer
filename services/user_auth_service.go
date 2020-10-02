package services

import (
	models2 "DulceDayServer/database/models"
	"DulceDayServer/helpers"
	"gorm.io/gorm"
)

type AuthService interface {
	GenerateToken(user *models2.User) (token string, err error)
}

type AuthServiceImpl struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthServiceImpl {
	return &AuthServiceImpl{
		db: db,
	}
}

func (as *AuthServiceImpl) GenerateToken(user *models2.User) (token string, err error) {
	t, e := helpers.GenerateToken(user.Username, user.Identifier)
	return t, e
}