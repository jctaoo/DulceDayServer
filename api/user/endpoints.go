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
	// 注册
	register(context *gin.Context)
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
	return userGroup
}
