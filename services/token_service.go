package services

import (
	models2 "DulceDayServer/database/models"
)

type TokenService interface {
	// 生成并保存新的 Token
	GenerateAndStoreNewToken(user *models2.User, ip string) string
	// 撤回某个 Token
	RevokeToken(token *models2.TokenAuth)
	// 撤回某个用户所有的 Token
	RevokeAllToken(user *models2.User)
	// 获取某个用户所有的 Token
	GetAllToken(user *models2.User) []*models2.TokenAuth
}

type TokenServiceImpl struct {
	tokenStore TokenStore
}

func NewTokenServiceImpl(tokenStore TokenStore) *TokenServiceImpl {
	return &TokenServiceImpl{tokenStore: tokenStore}
}

func (t TokenServiceImpl) GenerateNewToken(user *models2.User, ip string) string {
	panic("implement me")
}

func (t TokenServiceImpl) RevokeToken(token *models2.TokenAuth) {
	panic("implement me")
}

func (t TokenServiceImpl) RevokeAllToken(user *models2.User) {
	panic("implement me")
}

func (t TokenServiceImpl) GetAllToken(user *models2.User) []*models2.TokenAuth {
	panic("implement me")
}


