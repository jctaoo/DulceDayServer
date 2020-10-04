package user

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 与存储无关的无状态 Token 适配器
type TokenAdaptor interface {
	generateTokenStr(tokenAuth *models.TokenAuth, user *models.User) string
	verifyToken(tokenStr string) (isValidate bool, claims TokenClaims)
}

// JWT 荷载
type customClaims struct {
	userAuthority models.AuthorityLevel
	userIdentifier string
	jwt.StandardClaims
}

type TokenClaims struct {
	UserAuthority models.AuthorityLevel
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
		userAuthority: user.Authority,
		userIdentifier: user.Identifier,
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
		// todo log
	}
	return token
}

func (t TokenAdaptorImpl) verifyToken(tokenStr string) (isValidate bool, claims TokenClaims) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := []byte(config.SiteConfig.AuthTokenSecret)
		return jwtSecret, nil
	})
	if err != nil {
		return false, TokenClaims{}
	}
	if claims, ok := token.Claims.(customClaims); ok && token.Valid {
		return true, TokenClaims{
			UserIdentifier: claims.userIdentifier,
			UserAuthority: claims.userAuthority,
		}
	} else {
		return false, TokenClaims{}
	}
}


