package common

import (
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
)

// 授权细节模型，详见 common.MiddleWareAuth
type AuthDetail struct {
	// 已登陆用户的用户标示
	UserIdentifier string `json:"-"`

	// 已登陆用户的唯一用户名
	Username string `json:"-"`

	// 已登陆用户的权限
	AuthUserLevel models.AuthorityLevel `json:"-"`

	// 授权过程中发生的错误
	Error error `json:"-"`

	// 是否授权
	// 要在 handler 中处理未授权的情况，common.MiddleWareAuth 的第二个参数需要为 common.MiddleWareAuthPolicyAccess
	// 即即时未授权成功，也会执行 handler
	IsAuth bool `json:"-"`

	// 是否为敏感操作专属 token，即是否被允许进行敏感操作，如更改密码，邮箱，或唯一用户名(Username)等
	IsSensitiveAuth bool `json:"-"`
}

func GetAuthDetail(context *gin.Context) *AuthDetail {
	var err error = nil
	e, exists := context.Get(KAuthErrorKey)
	if exists {
		contextErr, ok := e.(error)
		if ok {
			err = contextErr
		}
	}

	return &AuthDetail{
		UserIdentifier:  context.GetString(KAuthUserIdentifierContextKey),
		Username:        context.GetString(KAuthUsernameContextKey),
		AuthUserLevel:   models.AuthorityLevel(context.GetInt(KAuthUserLevelContextKey)),
		Error:           err,
		IsAuth:          context.GetBool(KIsAuthKey),
		IsSensitiveAuth: context.GetBool(KIsSensitiveAuthKey),
	}
}
