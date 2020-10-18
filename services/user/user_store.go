package user

import (
	"DulceDayServer/database/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Store interface {
	NewUser(userIdentifier string, username string, userProfile models.UserProfile)
}

type StoreImpl struct {
	db *gorm.DB
	cdb *redis.Client
}

func NewStoreImpl(db *gorm.DB, cdb *redis.Client) *StoreImpl {
	return &StoreImpl{db: db, cdb: cdb}
}

func (s StoreImpl) NewUser(userIdentifier string, username string, userProfile models.UserProfile) {
	s.db.Create(&models.User{
		UserIdentifier: userIdentifier,
		Username: username,
		UserProfile: userProfile,
	})
}
