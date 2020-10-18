package user_profile

import (
	"DulceDayServer/database/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserProfileStore interface {
	createNewProfile(userIdentifier string) *models.UserProfile
	updateUserProfile(original *models.UserProfile, new *models.UserProfile)
	findUserProfileByUserIdentifier(userIdentifier string) *models.UserProfile
	findFullUserByUserIdentifier(userIdentifier string) *FullUser
	findFullUserByUsername(username string) *FullUser
}

type UserProfileStoreImpl struct {
	db  *gorm.DB
	cdb *redis.Client
}

func NewUserProfileStoreImpl(db *gorm.DB, cdb *redis.Client) *UserProfileStoreImpl {
	return &UserProfileStoreImpl{db: db, cdb: cdb}
}

func (u UserProfileStoreImpl) createNewProfile(userIdentifier string) *models.UserProfile {
	profile := &models.UserProfile{UserIdentifier: userIdentifier}
	u.db.Create(profile)
	return profile
}

func (u UserProfileStoreImpl) updateUserProfile(original *models.UserProfile, new *models.UserProfile) {
	u.db.Model(original).Where(&models.UserProfile{
		UserIdentifier: original.UserIdentifier,
	}).Updates(new)
}

func (u UserProfileStoreImpl) findUserProfileByUserIdentifier(userIdentifier string) *models.UserProfile {
	profile := &models.UserProfile{}
	u.db.Where(&models.UserProfile{
		UserIdentifier: userIdentifier,
	}).Take(profile)
	return profile
}

func (u UserProfileStoreImpl) findFullUserByUserIdentifier(userIdentifier string) *FullUser {
	fullUser := FullUser{
		UserProfile: &models.UserProfile{},
	}
	query := buildBaseQueryForFullUser(u.db).Where(&models.UserProfile{
		UserIdentifier: userIdentifier,
	})
	query.Scan(&fullUser)
	return &fullUser
}

func (u UserProfileStoreImpl) findFullUserByUsername(username string) *FullUser {
	fullUser := FullUser{}
	query := buildBaseQueryForFullUser(u.db).Where(&models.User{
		Username: username,
	})
	query.Scan(&fullUser)
	return &fullUser
}
