package user

import (
	"DulceDayServer/api/common"
	"DulceDayServer/services/auth"
	"DulceDayServer/services/user"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	common.BaseEndpoints
	// 登陆
	loginWithEmail(context *gin.Context)
	loginWithUsername(context *gin.Context)
	loginForSensitiveWithEmail(context *gin.Context) // 敏感登录, 需要事先使用以上方式登录
	// 注册
	register(context *gin.Context)
	registerForSensitiveWithEmail(context *gin.Context) // 敏感注册，用于生成验证码等, 需要事先登录
}

type EndpointsImpl struct {
	Endpoints
	service user.Service
	authService auth.Service
}

func NewEndpointsImpl(service user.Service, authService auth.Service) *EndpointsImpl {
	return &EndpointsImpl{
		service: service,
		authService: authService,
	}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	userGroup := router.Group("/user")
	userGroup.POST("/register", e.register)
	userGroup.POST("/login/email", e.loginWithEmail)
	userGroup.POST("/login/username", e.loginWithUsername)
	userGroup.POST(
		"/login/sensitive/email",
		common.MiddleWareAuth(e.authService, common.MiddleWareAuthPolicyReject),
		e.loginForSensitiveWithEmail,
	)
	userGroup.POST(
		"/register/sensitive/email",
		common.MiddleWareAuth(e.authService, common.MiddleWareAuthPolicyReject),
		e.registerForSensitiveWithEmail,
	)
	return userGroup
}
