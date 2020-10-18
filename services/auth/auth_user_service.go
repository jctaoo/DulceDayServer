package auth

import (
	"DulceDayServer/database/models"
)

type Service interface {
	// 鉴权, 并返回 Token 字符串
	// username: 可空
	// email: 可空 (优先使用 username)
	// password: 密码
	AuthenticateWithPassword(username string, email string, password string, ip string, deviceName string) (token string, err error)

	// 为敏感信息鉴权做准备，生成验证码，并存在缓存数据库中
	// username: 必填
	// email: 可空 (优先使用 email)
	PrepareForAuthForSensitiveVerification(username string, email string, ip string, deviceName string) (verificationCode string, err error)
	// 为敏感信息鉴权, 并返回 Token 字符串
	// email: 可空 (优先使用 email)
	// verificationCode: 验证码
	AuthenticateForSensitiveVerification(email string, verificationCode string, ip string, deviceName string) (token string, err error)

	// 授权
	Authorize(tokenStr string) (isValidate bool, claims TokenClaims, err error)

	// 检查用户密码是否正确
	// user: 完整的用户模型
	// password: 需要检验的密码
	checkMatchPassword(user *models.AuthUser, password string) bool

	NewUser(user *models.AuthUser) *models.AuthUser

	CheckUserInBlackList(userId *models.AuthUser) bool
	CheckUserInBlackListByUserIdentifier(userId string) bool
	AddUserInBlackList(user *models.AuthUser)
	RemoveUserFromBlackList(user *models.AuthUser)

	CheckUserExisting(user *models.AuthUser) bool
}

type ServiceImpl struct {
	encryptionAdaptor EncryptionAdaptor
	tokenGranter      TokenGranter
	userStore         Store
}

func NewServiceImpl(encryptionAdaptor EncryptionAdaptor, tokenGranter TokenGranter, userStore Store) *ServiceImpl {
	return &ServiceImpl{encryptionAdaptor: encryptionAdaptor, tokenGranter: tokenGranter, userStore: userStore}
}

func (g *ServiceImpl) AuthenticateWithPassword(username string, email string, password string, ip string, deviceName string) (token string, err error) {
	var user *models.AuthUser
	if username != "" {
		user = g.userStore.findUserByUserName(username)
	} else if email != "" {
		user = g.userStore.findUserByEmail(email)
	}
	if user == nil {
		err = ErrorUserNotFound{}
		return
	}

	// 检验 UserId 字符串是否在 BlackList 中
	inBlackList := g.CheckUserInBlackList(user)
	if inBlackList {
		err = ErrorUserIdInBlackList{}
		return
	}
	// 检查密码是否正确
	if !g.checkMatchPassword(user, password) {
		err = ErrorPasswordWrong{}
		return
	}
	token = g.tokenGranter.grantToken(user, ip, deviceName)
	return
}

func (g *ServiceImpl) PrepareForAuthForSensitiveVerification(username string, email string, ip string, deviceName string) (verificationCode string, err error) {
	var user *models.AuthUser
	if username != "" {
		user = g.userStore.findUserByUserName(username)
	} else if email != "" {
		user = g.userStore.findUserByEmail(email)
	}
	if user == nil {
		err = ErrorUserNotFound{}
		return
	}

	// 检验 UserId 字符串是否在 BlackList 中
	inBlackList := g.CheckUserInBlackList(user)
	if inBlackList {
		err = ErrorUserIdInBlackList{}
		return
	}

	// 生成 Token
	tokenStr := g.tokenGranter.grantTokenForSensitiveVerification(user, ip, deviceName)

	// 生成验证码，验证方式(邮箱等)为键，验证码为值(包含 tokenStr)
	verificationKey := email
	verificationCode = g.encryptionAdaptor.generateVerificationCode()
	// 存入缓存数据库
	g.userStore.saveVerificationCode(verificationKey, g.encryptionAdaptor.toHash(verificationCode), tokenStr)

	return
}

func (g *ServiceImpl) AuthenticateForSensitiveVerification(email string, verificationCode string, ip string, deviceName string) (token string, err error) {
	var code, tokenStr string
	doFailure := func() {
		// 若不成功 Revoke 相应 Token
		g.tokenGranter.RevokeTokenStr(tokenStr)
	}

	// 校验验证码
	code, tokenStr = g.userStore.getVerificationCode(email)
	g.userStore.removeVerificationCode(email) // 删除验证码键值
	if code == "" || tokenStr == "" {
		// 为空错误
		err = ErrorBadVerificationCode{}
		doFailure()
		return
	}
	if !g.encryptionAdaptor.verify(verificationCode, code) {
		// 验证码错误
		err = ErrorBadVerificationCode{}
		doFailure()
		return
	}
	// 验证成功后从 InActiveTokens 中移除
	g.tokenGranter.ActiveToken(tokenStr)

	// 校验 tokenStr
	var isValidate bool
	var claims TokenClaims
	isValidate, claims, err = g.tokenGranter.checkTokenForSensitiveVerification(tokenStr, ip, deviceName)
	if !isValidate || err != nil {
		return
	}

	// 检验 UserId 字符串是否在 BlackList 中
	inBlackList := g.CheckUserInBlackListByUserIdentifier(claims.UserIdentifier)
	if inBlackList {
		err = ErrorUserIdInBlackList{}
		return
	}

	// 当校验成功时，返回 TokenStr
	token = tokenStr
	return
}

func (g *ServiceImpl) Authorize(tokenStr string) (isValidate bool, claims TokenClaims, err error) {
	isValidate, claims, err = g.tokenGranter.authorize(tokenStr)
	if !isValidate || err != nil {
		return
	}
	// 检验 UserId 字符串是否在 BlackList 中
	inBlackList := g.CheckUserInBlackListByUserIdentifier(claims.UserIdentifier)
	if inBlackList {
		err = ErrorUserIdInBlackList{}
		return
	}
	return
}

func (g *ServiceImpl) checkMatchPassword(user *models.AuthUser, password string) bool {
	return g.encryptionAdaptor.verify(password, user.Password)
}

func (g *ServiceImpl) NewUser(user *models.AuthUser) *models.AuthUser {
	// 将用户密码由 MD5 处理
	user.Password = g.encryptionAdaptor.toHash(user.Password)
	g.userStore.newUser(user)
	return user
}

func (g *ServiceImpl) CheckUserInBlackList(user *models.AuthUser) bool {
	return g.CheckUserInBlackListByUserIdentifier(user.Identifier)
}

func (g *ServiceImpl) CheckUserInBlackListByUserIdentifier(userId string) bool {
	return g.userStore.checkUserInBlackList(userId)
}

func (g *ServiceImpl) AddUserInBlackList(user *models.AuthUser) {
	g.userStore.addUserInBlackList(user)
}

func (g *ServiceImpl) RemoveUserFromBlackList(user *models.AuthUser) {
	g.userStore.removeUserFromBlackList(user)
}

func (g *ServiceImpl) CheckUserExisting(user *models.AuthUser) bool {
	return g.userStore.checkUserExisting(user)
}
