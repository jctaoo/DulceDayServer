package user

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 存储用户 Token 以支持灵活的基于 Token 的操作，如支持后期的可能出现的 OAuth 需求，以及 Token 的颁发与撤回功能
type TokenStore interface {
	addNewTokenToUser(token *models.TokenAuth, user *models.User)
	updateToken(token *models.TokenAuth, ip string, tokenStr string)
	deleteToken(token *models.TokenAuth)

	revokeToken(token *models.TokenAuth)
	checkTokenIsRevoke(tokenStr string) bool
	removeTokenFromRevokeList(tokenStr string)
	checkTokenIsInActive(tokenStr string) bool
	removeTokenFromInActiveList(tokenStr string)

	findTokenByUserAndDeviceName(user *models.User, deviceName string) *models.TokenAuth
}

// 该实现采用 Redis 来实现 Token 的存储
type TokenStoreImpl struct {
	db *gorm.DB
	rdb *redis.Client
}

func NewTokenStoreImpl(db *gorm.DB, rdb *redis.Client) *TokenStoreImpl {
	return &TokenStoreImpl{db: db, rdb: rdb}
}

func (t TokenStoreImpl) addNewTokenToUser(token *models.TokenAuth, user *models.User) {
	user.Tokens = append(user.Tokens, *token)
	t.db.Save(user)
}

func (t TokenStoreImpl) updateToken(tokenAuth *models.TokenAuth, ip string, newToken string) {
	tokenAuth.TokenStr = newToken
	tokenAuth.Ip = ip
	t.db.Save(tokenAuth)
}

func (t TokenStoreImpl) deleteToken(token *models.TokenAuth) {
	t.db.Delete(token)
}

func (t TokenStoreImpl) revokeToken(token *models.TokenAuth) {
	revokeListName := config.SiteConfig.CacheConfig.RevokeTokenListName
	_, err := t.rdb.ZAdd(context.Background(), revokeListName, &redis.Z{Member: token.TokenStr, Score: 20}).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{"token": token}).WithError(err).Error("撤回 Token 发生错误")
	}
}

func (t TokenStoreImpl) checkTokenIsRevoke(tokenStr string) bool {
	revokeListName := config.SiteConfig.CacheConfig.RevokeTokenListName
	val, err := t.rdb.ZScore(context.Background(), revokeListName, tokenStr).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{"tokenStr": tokenStr}).WithError(err).Error("检查 Token 是否被撤回时发生错误")
		return false
	}
	if val == kTokenRevokeListScore {
		return false
	}
	return true
}

func (t TokenStoreImpl) removeTokenFromRevokeList(tokenStr string) {
	revokeListName := config.SiteConfig.CacheConfig.RevokeTokenListName
	_, err := t.rdb.ZRem(context.Background(), revokeListName, tokenStr).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{"tokenStr": tokenStr}).WithError(err).Error("将 Token 移除 RevokeList 时发生错误")
	}
}

func (t TokenStoreImpl) checkTokenIsInActive(tokenStr string) bool {
	revokeListName := config.SiteConfig.CacheConfig.InActiveTokenListName
	val, err := t.rdb.ZScore(context.Background(), revokeListName, tokenStr).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{"tokenStr": tokenStr}).WithError(err).Error("检查 Token 是否被激活时发生错误")
		return false
	}
	if val == kTokenInActiveListScore {
		return false
	}
	return true
}

func (t TokenStoreImpl) removeTokenFromInActiveList(tokenStr string) {
	revokeListName := config.SiteConfig.CacheConfig.InActiveTokenListName
	_, err := t.rdb.ZRem(context.Background(), revokeListName, tokenStr).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{"tokenStr": tokenStr}).WithError(err).Error("激活 Token 时发生错误")
	}
}

func (t TokenStoreImpl) findTokenByUserAndDeviceName(user *models.User, deviceName string) *models.TokenAuth {
	token := &models.TokenAuth{}
	t.db.Where(&models.TokenAuth{UserID: user.ID, DeviceName: deviceName}).Take(token)
	return token
}
