package user

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

// 存储用户 Token 以支持灵活的基于 Token 的操作，如支持后期的可能出现的 OAuth 需求，以及 Token 的颁发与撤回功能
type TokenStore interface {
	addNewTokenToUser(token *models.TokenAuth, user *models.User)
	updateToken(token *models.TokenAuth, ip string, tokenStr string)
	deleteToken(token *models.TokenAuth)

	revokeToken(token *models.TokenAuth)
	checkTokenIsRevoke(tokenStr string) bool
	removeTokenFromRevokeList(tokenStr string)
	inactiveToken(token *models.TokenAuth)
	checkTokenIsInActive(tokenStr string) bool
	removeTokenFromInActiveList(tokenStr string)

	findTokenByUserAndDeviceName(user *models.User, deviceName string) *models.TokenAuth
	findTokenByTokenStr(tokenStr string) *models.TokenAuth
}

// 该实现采用 Redis 来实现 Token 的存储
type TokenStoreImpl struct {
	db  *gorm.DB
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
	tokenAuth.ExpireTime = time.Unix(time.Now().Unix()+config.SiteConfig.AuthTokenExpiresTime, 0)
	t.db.Save(tokenAuth)
}

func (t TokenStoreImpl) deleteToken(token *models.TokenAuth) {
	t.db.Delete(token)
}

func (t TokenStoreImpl) revokeToken(token *models.TokenAuth) {
	revokeListName := config.SiteConfig.CacheConfig.RevokeTokenListName
	t.rdb.ZAdd(context.Background(), revokeListName, &redis.Z{Member: token.TokenStr, Score: kTokenRevokeListScore})
}

func (t TokenStoreImpl) checkTokenIsRevoke(tokenStr string) bool {
	revokeListName := config.SiteConfig.CacheConfig.RevokeTokenListName
	val := t.rdb.ZScore(context.Background(), revokeListName, tokenStr).Val()
	if val == kTokenRevokeListScore {
		return true
	}
	return false
}

func (t TokenStoreImpl) removeTokenFromRevokeList(tokenStr string) {
	revokeListName := config.SiteConfig.CacheConfig.RevokeTokenListName
	t.rdb.ZRem(context.Background(), revokeListName, tokenStr)
}

func (t TokenStoreImpl) inactiveToken(token *models.TokenAuth) {
	revokeListName := config.SiteConfig.CacheConfig.InActiveTokenListName
	t.rdb.ZAdd(context.Background(), revokeListName, &redis.Z{Member: token.TokenStr, Score: kTokenInActiveListScore})
}

func (t TokenStoreImpl) checkTokenIsInActive(tokenStr string) bool {
	revokeListName := config.SiteConfig.CacheConfig.InActiveTokenListName
	val := t.rdb.ZScore(context.Background(), revokeListName, tokenStr).Val()
	if val == kTokenInActiveListScore {
		return true
	}
	return false
}

func (t TokenStoreImpl) removeTokenFromInActiveList(tokenStr string) {
	revokeListName := config.SiteConfig.CacheConfig.InActiveTokenListName
	t.rdb.ZRem(context.Background(), revokeListName, tokenStr)
}

func (t TokenStoreImpl) findTokenByUserAndDeviceName(user *models.User, deviceName string) *models.TokenAuth {
	token := &models.TokenAuth{}
	t.db.Where(&models.TokenAuth{UserID: user.ID, DeviceName: deviceName}).Take(token)
	return token
}

func (t TokenStoreImpl) findTokenByTokenStr(tokenStr string) *models.TokenAuth {
	token := &models.TokenAuth{}
	t.db.Where(&models.TokenAuth{TokenStr: tokenStr}).Take(token)
	return token
}
