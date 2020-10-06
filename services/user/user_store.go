package user

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Store interface {
	newUser(user *models.User) *models.User
	checkUserInBlackList(userId string) bool
	addUserInBlackList(user *models.User)
	removeUserFromBlackList(user *models.User)
	findUserByUserName(username string) *models.User
	findUserByIdentifier(identifier string) *models.User
	findUserByEmail(email string) *models.User
	checkUserExisting(user *models.User) bool
}

type StoreImpl struct {
	db *gorm.DB
	rdb *redis.Client
}

func NewStoreImpl(db *gorm.DB, rdb *redis.Client) *StoreImpl {
	return &StoreImpl{db: db, rdb: rdb}
}

func (u StoreImpl) newUser(user *models.User) *models.User {
	u.db.Create(user)
	return user
}

func (u StoreImpl) checkUserInBlackList(userId string) bool {
	userIdBlackListName := config.SiteConfig.CacheConfig.BlackListName
	val := u.rdb.ZScore(context.Background(), userIdBlackListName, userId).Val()
	if val == kUserIdBlackListScore {
		return true
	}
	return false
}

func (u StoreImpl) addUserInBlackList(user *models.User) {
	userIdBlackListName := config.SiteConfig.CacheConfig.BlackListName
	u.rdb.ZAdd(context.Background(), userIdBlackListName, &redis.Z{Member: user.Identifier, Score: kUserIdBlackListScore})
}

func (u StoreImpl) removeUserFromBlackList(user *models.User) {
	userIdBlackListName := config.SiteConfig.CacheConfig.BlackListName
	u.rdb.ZRem(context.Background(), userIdBlackListName, user.Identifier)
}

func (u StoreImpl) findUserByUserName(username string) *models.User {
	user := &models.User{}
	u.db.Where("Username = ?", username).First(user)
	return user
}

func (u StoreImpl) findUserByIdentifier(identifier string) *models.User {
	user := &models.User{}
	u.db.Where("identifier = ?", identifier).First(user)
	return user
}

func (u StoreImpl) findUserByEmail(email string) *models.User {
	user := &models.User{}
	u.db.Where("email = ?", email).First(user)
	return user
}

func (u StoreImpl) checkUserExisting(user *models.User) bool {
	// 此处同时检查用户名和邮箱
	var resUsers []models.User
	if !user.Validate() {
		return false
	}
	rows := u.db.Where("Username = ? OR email = ?", user.Username, user.Email).Find(&resUsers).RowsAffected
	return rows > 0
}
