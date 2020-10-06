package user

import (
	"DulceDayServer/database/models"
	"github.com/sirupsen/logrus"
)

type TokenGranter interface {
	// 授予 Token (鉴权的末尾--颁发 Token)
	grantToken(user *models.User, ip string, deviceName string) string

	// 授权
	authorize(tokenStr string) (isValidate bool, claims TokenClaims, err error)

	RevokeToken(token *models.TokenAuth)
	CheckTokenIsRevoke(tokenStr string) bool
	RemoveTokenFromRevokeList(tokenStr string)

	CheckTokenIsInActive(tokenStr string) bool
	ActiveToken(tokenStr string)
}

type TokenGranterImpl struct {
	tokenStore   TokenStore
	tokenAdaptor TokenAdaptor
}

func NewTokenGranterImpl(tokenStore TokenStore, tokenAdaptor TokenAdaptor) *TokenGranterImpl {
	return &TokenGranterImpl{tokenStore: tokenStore, tokenAdaptor: tokenAdaptor}
}

func (t TokenGranterImpl) grantToken(user *models.User, ip string, deviceName string) string {
	// 查看是否有相应的 Token 记录
	tokenAuth := t.tokenStore.findTokenByUserAndDeviceName(user, deviceName)
	if !tokenAuth.IsEmpty() {
		// 如果有，就将原 Token 放入 Revoke 列表，更新 Token 并返回新的 Token 字符串和 IP
		logrus.WithField("token", tokenAuth.TokenStr).Debug("由于用户重新生成Token，原Token被撤回")
		t.tokenStore.revokeToken(tokenAuth)
		newToken := t.tokenAdaptor.generateTokenStr(tokenAuth, user)
		// 更新
		t.tokenStore.updateToken(tokenAuth, ip, newToken)
		return newToken
	} else {
		// 否则写入新的 Token 记录
		tokenAuth := models.NewTokenAuthWithoutTokenStr(user, ip, deviceName)
		// 生成字符串 Token
		tokenStr := t.tokenAdaptor.generateTokenStr(tokenAuth, user)
		tokenAuth.TokenStr = tokenStr
		// 将 Token 信息持久化到数据库
		t.tokenStore.addNewTokenToUser(tokenAuth, user)
		return tokenStr
	}
}

func (t TokenGranterImpl) authorize(tokenStr string) (isValidate bool, claims TokenClaims, err error) {
	// 检验 Token 字符串是否在 RevokeTokens 中
	if t.CheckTokenIsRevoke(tokenStr) {
		return false, TokenClaims{}, ErrorTokenRevoke{}
	}
	// 检验 Token 字符串是否在 InActiveTokens 中
	if t.CheckTokenIsInActive(tokenStr) {
		return false, TokenClaims{}, ErrorTokenInActive{}
	}
	// JWT 检查 Token 字符串
	isValidate, claims = t.tokenAdaptor.verifyToken(tokenStr)
	if !isValidate {
		return false, claims, ErrorTokenBad{}
	}
	return isValidate, claims, nil
}

func (t TokenGranterImpl) RevokeToken(token *models.TokenAuth) {
	// 将 Token 字符串存入缓存数据库 RevokeTokens 中
	t.tokenStore.revokeToken(token)
	// 删除相应 Token 记录 (软删)
	t.tokenStore.deleteToken(token)
}

func (t TokenGranterImpl) CheckTokenIsRevoke(tokenStr string) bool {
	return t.tokenStore.checkTokenIsRevoke(tokenStr)
}

func (t TokenGranterImpl) RemoveTokenFromRevokeList(tokenStr string) {
	t.tokenStore.removeTokenFromRevokeList(tokenStr)
}

func (t TokenGranterImpl) CheckTokenIsInActive(tokenStr string) bool {
	return t.tokenStore.checkTokenIsInActive(tokenStr)
}

func (t TokenGranterImpl) ActiveToken(tokenStr string) {
	t.tokenStore.removeTokenFromInActiveList(tokenStr)
}
