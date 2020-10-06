package common

import (
	serviceUser "DulceDayServer/services/user"
	"github.com/gin-gonic/gin"
)

const KAuthUserIdentifierContextKey = "user_identifier"
const KAuthUsernameContextKey = "username"
const KAuthUserLevelContextKey = "user_level"
const KAuthErrorKey = "auth_error"
const KIsAuthKey = "is_auth"

func MiddleWareAuth(service serviceUser.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 Token
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			context.Set(KIsAuthKey, false)
			context.Next()
			return
		}

		// 校验 Token
		isValidate, claims, err := service.Authorize(tokenStr)
		if err != nil {
			HttpLogger(context, err, gin.H{
				"bearer_token": tokenStr,
			}).Info("用户授权时发生错误")
			context.Set(KAuthErrorKey, err)
		}
		if isValidate {
			HttpLogger(context, nil, gin.H{
				"bearer_token": tokenStr,
				"claims": claims,
			}).Info("用户授权成功")
			context.Set(KIsAuthKey, true)
			context.Set(KAuthUserIdentifierContextKey, claims.UserIdentifier)
			context.Set(KAuthUsernameContextKey, claims.Username)
			context.Set(KAuthUserLevelContextKey, claims.UserAuthority)
		} else {
			context.Set(KIsAuthKey, false)
		}
		context.Next()
	}
}
