package auth

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strings"
)

type Store interface {
	newUser(user *models.AuthUser) *models.AuthUser
	checkUserInBlackList(userId string) bool
	addUserInBlackList(user *models.AuthUser)
	removeUserFromBlackList(user *models.AuthUser)
	findUserByUserName(username string) *models.AuthUser
	findUserByIdentifier(identifier string) *models.AuthUser
	findUserByEmail(email string) *models.AuthUser
	checkUserExisting(user *models.AuthUser) bool

	saveVerificationCode(key string, value string, tokenStr string)
	getVerificationCode(key string) (verificationCode string, tokenStr string)
	removeVerificationCode(key string)
}

type StoreImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewStoreImpl(db *gorm.DB, rdb *redis.Client) *StoreImpl {
	return &StoreImpl{db: db, rdb: rdb}
}

func (u StoreImpl) newUser(user *models.AuthUser) *models.AuthUser {
	u.db.Create(user)
	return user
}

func (u StoreImpl) checkUserInBlackList(userId string) bool {
	userIdBlackListName := config.SiteConfig.CacheConfig.BlackListName
	val := u.rdb.ZScore(context.Background(), userIdBlackListName, userId).Val()
	if val == KUserIdBlackListScore {
		return true
	}
	return false
}

func (u StoreImpl) addUserInBlackList(user *models.AuthUser) {
	userIdBlackListName := config.SiteConfig.CacheConfig.BlackListName
	u.rdb.ZAdd(context.Background(), userIdBlackListName, &redis.Z{Member: user.Identifier, Score: KUserIdBlackListScore})
}

func (u StoreImpl) removeUserFromBlackList(user *models.AuthUser) {
	userIdBlackListName := config.SiteConfig.CacheConfig.BlackListName
	u.rdb.ZRem(context.Background(), userIdBlackListName, user.Identifier)
}

func (u StoreImpl) findUserByUserName(username string) *models.AuthUser {
	user := &models.AuthUser{}
	u.db.Where("Username = ?", username).First(user)
	return user
}

func (u StoreImpl) findUserByIdentifier(identifier string) *models.AuthUser {
	user := &models.AuthUser{}
	u.db.Where("identifier = ?", identifier).First(user)
	return user
}

func (u StoreImpl) findUserByEmail(email string) *models.AuthUser {
	user := &models.AuthUser{}
	u.db.Where("email = ?", email).First(user)
	return user
}

func (u StoreImpl) checkUserExisting(user *models.AuthUser) bool {
	// 此处同时检查用户名和邮箱
	var resUsers []models.AuthUser
	if !user.Validate() {
		return false
	}
	rows := u.db.Where("Username = ? OR email = ?", user.Username, user.Email).Find(&resUsers).RowsAffected
	return rows > 0
}

func (u StoreImpl) saveVerificationCode(key string, value string, tokenStr string) {
	verificationCodeListName := config.SiteConfig.CacheConfig.VerificationCodeListName
	u.rdb.HSet(context.Background(), verificationCodeListName, map[string]interface{}{key: value + ":" + tokenStr})
}

func (u StoreImpl) getVerificationCode(key string) (verificationCode string, tokenStr string) {
	verificationCodeListName := config.SiteConfig.CacheConfig.VerificationCodeListName
	codeAndToken := u.rdb.HGet(context.Background(), verificationCodeListName, key).Val()
	if codeAndToken == "" {
		return "", ""
	}
	splitCodeAndToken := strings.Split(codeAndToken, ":")
	verificationCode = splitCodeAndToken[0]
	tokenStr = splitCodeAndToken[1]
	return
}

func (u StoreImpl) removeVerificationCode(key string) {
	verificationCodeListName := config.SiteConfig.CacheConfig.VerificationCodeListName
	u.rdb.HDel(context.Background(), verificationCodeListName, key)
}
