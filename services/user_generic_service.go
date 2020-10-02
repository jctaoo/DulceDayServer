package services

import (
	models2 "DulceDayServer/database/models"
	"gorm.io/gorm"
)
import "github.com/satori/go.uuid"

type GenericService interface {
	NewUser(user *models2.User) error
	FindUserByUserName(username string) *models2.User
	FindUserByIdentifier(identifier string) *models2.User
	FindUserByEmail(email string) *models2.User
	CheckMatchPassword(user *models2.User, password string) bool
	CheckExisting(user *models2.User) bool
	GenerateIdentifier() string
}

type GenericServiceImpl struct {
	db *gorm.DB
}

func NewGenericService(db *gorm.DB) *GenericServiceImpl {
	return &GenericServiceImpl{
		db: db,
	}
}

func (g *GenericServiceImpl) NewUser(user *models2.User) error {
	err := g.db.Create(user).Error
	return err
}

func (g *GenericServiceImpl) FindUserByUserName(username string) *models2.User {
	user := &models2.User{}
	g.db.Where("username = ?", username).First(user)
	return user
}

func (g *GenericServiceImpl) FindUserByIdentifier(identifier string) *models2.User {
	user := &models2.User{}
	g.db.Where("identifier = ?", identifier).First(user)
	return user
}

func (g *GenericServiceImpl) FindUserByEmail(email string) *models2.User {
	user := &models2.User{}
	g.db.Where("email = ?", email).First(user)
	return user
}

func (g *GenericServiceImpl) CheckMatchPassword(user *models2.User, password string) bool {
	resUser := &models2.User{}
	g.db.Where(user).First(resUser)
	return resUser.Password == password
}

func (g *GenericServiceImpl) CheckExisting(user *models2.User) bool {
	// 此处同时检查用户名和邮箱
	resUser := &models2.User{}
	if !user.Validate() {
		return false
	}
	rows := g.db.Where("username = ? OR (email = ? AND email != \"\")", user.Username, user.Email).Find(resUser).Statement.RowsAffected
	return rows > 0
}

func (g *GenericServiceImpl) GenerateIdentifier() string {
	return uuid.NewV4().String()
}