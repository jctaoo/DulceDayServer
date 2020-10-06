package user_profile

import (
	"DulceDayServer/database/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserProfileStore interface {
	getProfileByUsername(username string) *models.UserProfile
	createNewProfile(username string) *models.UserProfile
	updateUserProfile(original *models.UserProfile, new *models.UserProfile)
}

type UserProfileStoreImpl struct {
	db *gorm.DB
	cdb *redis.Client
}

func NewUserProfileStoreImpl(db *gorm.DB, cdb *redis.Client) *UserProfileStoreImpl {
	return &UserProfileStoreImpl{db: db, cdb: cdb}
}

func (u UserProfileStoreImpl) getProfileByUsername(username string) *models.UserProfile {
	userProfile := &models.UserProfile{}
	u.db.Where("username = ?", username).Find(userProfile)
	return userProfile
}

func (u UserProfileStoreImpl) createNewProfile(username string) *models.UserProfile {
	profile := &models.UserProfile{Username: username}
	u.db.Create(profile)
	return profile
}

func (u UserProfileStoreImpl) updateUserProfile(original *models.UserProfile, new *models.UserProfile) {
	u.db.Model(original).Where("username = ?", original.Username).Updates(new)
}