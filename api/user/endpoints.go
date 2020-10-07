package user

import (
	"DulceDayServer/api/common"
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
	service           user.Service
}

func NewEndpointsImpl(genericService user.Service) *EndpointsImpl {
	return &EndpointsImpl{
		service:           genericService,
	}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	userGroup := router.Group("/user")
	userGroup.POST("/register", e.register)
	userGroup.POST("/login/email", e.loginWithEmail)
	userGroup.POST("/login/username", e.loginWithUsername)
	userGroup.POST("/login/sensitive/email", common.MiddleWareAuth(e.service), e.loginForSensitiveWithEmail)
	userGroup.POST("/register/sensitive/email", common.MiddleWareAuth(e.service), e.registerForSensitiveWithEmail)
	return userGroup
}
