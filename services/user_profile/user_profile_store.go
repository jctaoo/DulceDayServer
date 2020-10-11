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
	db *gorm.DB
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
	fullUser := FullUser{}
	query := u.db.Table("user_profiles")
	query = query.Select("user_profiles.uid, user_profiles.nickname, users.username, " +
		"users.identifier as user_identifier, user_profiles.avatar_file_key")
	query = query.Joins("LEFT OUTER JOIN users ON user_profiles.user_identifier = " +
		"users.identifier")
	query = query.Where("user_profiles.user_identifier = ?", userIdentifier)
	query = query.Group("user_profiles.id").Group("users.id")
	query.Scan(&fullUser)
	return &fullUser
}

func (u UserProfileStoreImpl) findFullUserByUsername(username string) *FullUser {
	fullUser := FullUser{}
	query := u.db.Table("user_profiles")
	query = query.Select("user_profiles.uid, user_profiles.nickname, users.username, " +
		"users.identifier as user_identifier, user_profiles.avatar_file_key")
	query = query.Joins("LEFT OUTER JOIN users ON user_profiles.user_identifier = " +
		"users.identifier")
	query = query.Where("users.username = ?", username)
	query = query.Group("user_profiles.id").Group("users.id")
	query.Scan(&fullUser)
	return &fullUser
}
