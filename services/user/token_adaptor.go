package user

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

// 与存储无关的无状态 Token 适配器
type TokenAdaptor interface {
	generateTokenStr(tokenAuth *models.TokenAuth, user *models.User) string
	verifyToken(tokenStr string) (isValidate bool, claims TokenClaims)
}

// JWT 荷载
type customClaims struct {
	UserAuthority  models.AuthorityLevel
	UserIdentifier string
	Username       string
	jwt.StandardClaims
}

type TokenClaims struct {
	UserAuthority models.AuthorityLevel
	Username string
	UserIdentifier string
}

// 采用 JWT 实现的 TokenAdaptor
type TokenAdaptorImpl struct {

}

func NewTokenAdaptorImpl() *TokenAdaptorImpl {
	return &TokenAdaptorImpl{}
}

func (t TokenAdaptorImpl) generateTokenStr(tokenAuth *models.TokenAuth, user *models.User) string {
	expiresTime := time.Now().Unix() + config.SiteConfig.AuthTokenExpiresTime
	claims := customClaims{
		UserAuthority:  user.Authority,
		UserIdentifier: user.Identifier,
		Username:       user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime,               // 失效时间
			IssuedAt:  time.Now().Unix(),         // 签发时间
			Issuer:    config.SiteConfig.AppName, // 签发人
			NotBefore: time.Now().Unix(),         // 生效时间
		},
	}
	jwtSecret := []byte(config.SiteConfig.AuthTokenSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		logrus.WithFields(logrus.Fields{"tokenAuth": tokenAuth, "user": user}).WithError(err).Error("生成 TokenStr 发生错误")
	}
	return token
}

func (t TokenAdaptorImpl) verifyToken(tokenStr string) (isValidate bool, claims TokenClaims) {
	token, err := jwt.ParseWithClaims(tokenStr, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := []byte(config.SiteConfig.AuthTokenSecret)
		return jwtSecret, nil
	})
	if err != nil {
		return false, TokenClaims{}
	}
	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return true, TokenClaims{
			UserIdentifier: claims.UserIdentifier,
			UserAuthority: claims.UserAuthority,
			Username: claims.Username,
		}
	} else {
		return false, TokenClaims{}
	}
}


