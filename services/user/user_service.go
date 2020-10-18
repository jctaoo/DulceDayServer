package user

import (
	"DulceDayServer/database/models"
	"DulceDayServer/helpers"
	"DulceDayServer/services/auth"
	"DulceDayServer/services/user_profile"
)

// user.Service 仅用作 user 模块的服务，其他模块对于授权鉴权服务的调用直接使用 auth 模块
type Service interface {
	// 注册
	Register(username string, email string, password string)

	// 使用密码鉴权
	AuthenticateWithPassword(username string, email string, password string, ip string, deviceName string) (token string, err error)

	// 为敏感信息鉴权做准备，生成验证码
	PrepareForAuthForSensitiveVerification(username string, email string, ip string, deviceName string) (verificationCode string, err error)

	// 敏感操作鉴权
	AuthenticateForSensitiveVerification(email string, verificationCode string, ip string, deviceName string) (token string, err error)

	CheckUserExisting(username string, email string) bool
}

type ServiceImpl struct {
	store Store
	authService auth.Service
	profileService user_profile.Service
}

func NewServiceImpl(store Store, authService auth.Service, profileService user_profile.Service) *ServiceImpl {
	return &ServiceImpl{store: store, authService: authService, profileService: profileService}
}

func (s ServiceImpl) Register(username string, email string, password string) {
	authUser := &models.AuthUser{
		Username: username,
		Password: password,
		Email:    email,
	}
	identifier := helpers.GenerateRandomUserID()
	// 在user模块加入相应记录
	up := s.profileService.CreateNewProfileByUserIdentifier(identifier)
	s.store.NewUser(identifier, authUser.Username, *up)
	// 在授权鉴权模块加入记录
	authUser.Identifier = identifier
	s.authService.NewUser(authUser)
}

func (s ServiceImpl) AuthenticateWithPassword(username string, email string, password string, ip string, deviceName string) (token string, err error) {
	return s.authService.AuthenticateWithPassword(username, email, password, ip, deviceName)
}

func (s ServiceImpl) PrepareForAuthForSensitiveVerification(username string, email string, ip string, deviceName string) (verificationCode string, err error) {
	return s.authService.PrepareForAuthForSensitiveVerification(username, email, ip, deviceName)
}

func (s ServiceImpl) AuthenticateForSensitiveVerification(email string, verificationCode string, ip string, deviceName string) (token string, err error) {
	return s.authService.AuthenticateForSensitiveVerification(email, verificationCode, ip, deviceName)
}

func (s ServiceImpl) CheckUserExisting(username string, email string) bool {
	return s.authService.CheckUserExisting(&models.AuthUser{
		Username: username,
		Email:    email,
	})
}
