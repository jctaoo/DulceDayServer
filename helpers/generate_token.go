package helpers

import (
	"DulceDayServer/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(audience string, identifier string) (token string, err error) {
	expiresTime := time.Now().Unix() + config.SiteConfig.AuthTokenExpiresTime
	claims := jwt.StandardClaims{
		Audience:  audience,                  // 受众
		ExpiresAt: expiresTime,               // 失效时间
		Id:        identifier,                // 编号
		IssuedAt:  time.Now().Unix(),         // 签发时间
		Issuer:    config.SiteConfig.AppName, // 签发人
		NotBefore: time.Now().Unix(),         // 生效时间
		Subject:   "generate-token",          // 主题
	}
	jwtSecret := []byte(config.SiteConfig.AuthTokenSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}