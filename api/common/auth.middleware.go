package common

import (
	"DulceDayServer/database/models"
	serviceAuth "DulceDayServer/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

const KAuthUserIdentifierContextKey = "user_identifier"
const KAuthUsernameContextKey = "username"
const KAuthUserLevelContextKey = "user_level"
const KAuthErrorKey = "auth_error"
const KIsAuthKey = "is_auth"
const KIsSensitiveAuthKey = "is_sensitive_auth"

// 请求头中 token 字段为空错误
type ErrorAuthEmptyTokenHeader struct{}

func (e ErrorAuthEmptyTokenHeader) Error() string {
	return "token请求头为空"
}

// token 验证失败错误
type ErrorAuthValidateWrong struct{}

func (e ErrorAuthValidateWrong) Error() string {
	return "token验证错误"
}

// 鉴权等级错误
type ErrorAuthorityLevel struct{}

func (e ErrorAuthorityLevel) Error() string {
	return "错误的鉴权等级"
}

// MiddleWareAuth 的通过策略
type MiddleWareAuthPolicy int

const (
	// MiddleWareAuth 的通过策略: 未登陆成功，直接拒绝，不会执行 handler
	MiddleWareAuthPolicyReject MiddleWareAuthPolicy = iota

	// MiddleWareAuth 的通过策略: 即时未登陆成功，也会执行 handler
	MiddleWareAuthPolicyAccess
)

// gin 鉴权等级中间件
// 注意：在 MiddleWareAuth 之后使用
func MiddleWareAuthorityLevel(
	level models.AuthorityLevel,
) gin.HandlerFunc {
	return func(context *gin.Context) {
		detail := GetAuthDetail(context)
		print(detail.AuthUserLevel, level)
		if detail.AuthUserLevel < level {
			HandleHttpErr(ErrorAuthorityLevel{}, context)
			context.Abort()
			return
		}
	}
}

// gin 鉴权中间件
func MiddleWareAuth(service serviceAuth.Service, accessPolicy MiddleWareAuthPolicy) gin.HandlerFunc {
	doUnLogin := func(context *gin.Context, err error) {
		if accessPolicy == MiddleWareAuthPolicyReject {
			context.JSON(http.StatusUnauthorized, BaseResponse{
				Code:    40001,
				Message: err.Error(),
			})
			context.Abort()
		} else if accessPolicy == MiddleWareAuthPolicyAccess {
			context.Set(KIsAuthKey, false)
			context.Set(KAuthErrorKey, err)
			context.Set(KAuthUserIdentifierContextKey, nil)
			context.Set(KAuthUsernameContextKey, nil)
			context.Set(KAuthUserLevelContextKey, nil)
			context.Set(KIsSensitiveAuthKey, nil)
		}
	}

	return func(context *gin.Context) {
		// 获取 Token
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			doUnLogin(context, ErrorAuthEmptyTokenHeader{})
			return
		}

		// 校验 Token
		isValidate, claims, err := service.Authorize(tokenStr)
		if err != nil {
			HttpLogger(context, err, gin.H{
				"bearer_token": tokenStr,
			}).Debug("用户授权时发生错误")
			doUnLogin(context, err)
			return
		}
		if isValidate {
			HttpLogger(context, nil, gin.H{
				"bearer_token": tokenStr,
				"claims":       claims,
			}).Debug("用户授权成功")
			context.Set(KIsAuthKey, true)
			context.Set(KAuthUserIdentifierContextKey, claims.UserIdentifier)
			context.Set(KAuthUsernameContextKey, claims.Username)
			context.Set(KAuthUserLevelContextKey, int(claims.UserAuthority))
			context.Set(KIsSensitiveAuthKey, claims.SensitiveVerification)
			return
		} else {
			doUnLogin(context, ErrorAuthValidateWrong{})
			return
		}
	}
}
