package common

import (
	serviceUser "DulceDayServer/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

const KAuthUserIdentifierContextKey = "user_identifier"
const KAuthUsernameContextKey = "username"
const KAuthUserLevelContextKey = "user_level"
const KAuthErrorKey = "auth_error"
const KIsAuthKey = "is_auth"
const KIsSensitiveAuth = "is_sensitive_auth"

func MiddleWareAuth(service serviceUser.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 Token
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			context.JSON(http.StatusUnauthorized, BaseResponse{
				Code: 40001,
				Message: "未登录",
			})
			return
		}

		// 校验 Token
		isValidate, claims, err := service.Authorize(tokenStr)
		if err != nil {
			HttpLogger(context, err, gin.H{
				"bearer_token": tokenStr,
			}).Info("用户授权时发生错误")
			context.Set(KAuthErrorKey, err)
			context.JSON(http.StatusUnauthorized, BaseResponse{
				Code: 40001,
				Message: "未登录",
			})
			return
		}
		if isValidate {
			HttpLogger(context, nil, gin.H{
				"bearer_token": tokenStr,
				"claims": claims,
			}).Debug("用户授权成功")
			context.Set(KIsAuthKey, true)
			context.Set(KAuthUserIdentifierContextKey, claims.UserIdentifier)
			context.Set(KAuthUsernameContextKey, claims.Username)
			context.Set(KAuthUserLevelContextKey, claims.UserAuthority)
			context.Set(KIsSensitiveAuth, claims.SensitiveVerification)
			context.Next()
		} else {
			context.JSON(http.StatusUnauthorized, BaseResponse{
				Code: 40001,
				Message: "未登录",
			})
			return
		}
	}
}
