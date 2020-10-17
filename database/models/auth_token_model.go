package models

import (
	"DulceDayServer/config"
	"gorm.io/gorm"
	"time"
)

// 用于授权操作的 Token 模型，描述了与 Token 相关的元数据，不但支持与用户模型相关联，也支持与获取以及使用该 Token 的请求等的相关数据
// 进行关联，从而实现检测异地登陆，登陆设备，用户可控制的强行下线其他设备等的功能
type TokenAuth struct {
	gorm.Model `json:"-"`

	TokenStr   string `json:"-"`
	Ip         string `json:"-"`
	StartTime  time.Time
	ExpireTime time.Time
	DeviceName string `json:"-"`
	AuthUserID uint   `json:"-"`
}

func NewTokenAuthWithoutTokenStr(user *AuthUser, ip string, deviceName string) *TokenAuth {
	startTime := time.Now()
	expiresTime := startTime.Unix() + config.SiteConfig.AuthTokenExpiresTime
	return &TokenAuth{
		Ip:         ip,
		StartTime:  startTime,
		ExpireTime: time.Unix(expiresTime, 0),
		DeviceName: deviceName,
		AuthUserID: user.ID,
	}
}

func (t TokenAuth) IsEmpty() bool {
	return t == TokenAuth{}
}
