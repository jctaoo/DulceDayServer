package user

import (
	"DulceDayServer/database/models"
	"github.com/satori/go.uuid"
)

type Service interface {
	// 鉴权, 并返回 Token 字符串
	// Username: 可空
	// email: 可空
	// password: 密码
	AuthenticateWithPassword(username string, email string, password string, ip string, deviceName string) (token string, err error)

	// 授权
	Authorize(tokenStr string) (isValidate bool, claims TokenClaims, err error)

	// 检查用户密码是否正确
	// user: 完整的用户模型
	// password: 需要检验的密码
	checkMatchPassword(user *models.User, password string) bool

	NewUser(user *models.User) *models.User

	CheckUserInBlackList(userId *models.User) bool
	CheckUserInBlackListByUserIdentifier(userId string) bool
	AddUserInBlackList(user *models.User)
	RemoveUserFromBlackList(user *models.User)

	CheckUserExisting(user *models.User) bool
	GenerateUserIdentifier() string
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
	var user *models.User
	if username != "" {
		user = g.userStore.findUserByUserName(username)
	} else if email != "" {
		user = g.userStore.findUserByEmail(email)
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

func (g *ServiceImpl) checkMatchPassword(user *models.User, password string) bool {
	return g.encryptionAdaptor.verify(password, user.Password)
}

func (g *ServiceImpl) NewUser(user *models.User) *models.User {
	// 将用户密码由 MD5 处理
	user.Password = g.encryptionAdaptor.toHash(user.Password)
	g.userStore.newUser(user)
	return user
}

func (g *ServiceImpl) CheckUserInBlackList(user *models.User) bool {
	return g.CheckUserInBlackListByUserIdentifier(user.Identifier)
}

func (g *ServiceImpl) CheckUserInBlackListByUserIdentifier(userId string) bool {
	return g.userStore.checkUserInBlackList(userId)
}

func (g *ServiceImpl) AddUserInBlackList(user *models.User) {
	g.userStore.addUserInBlackList(user)
}

func (g *ServiceImpl) RemoveUserFromBlackList(user *models.User) {
	g.userStore.removeUserFromBlackList(user)
}

func (g *ServiceImpl) CheckUserExisting(user *models.User) bool {
	return g.userStore.checkUserExisting(user)
}

func (g *ServiceImpl) GenerateUserIdentifier() string {
	return uuid.NewV4().String()
}
