package user

import (
	"DulceDayServer/api/base"
	"DulceDayServer/services"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	base.BaseEndpoints
	// 登陆
	login(context *gin.Context)
	// 注册
	register(context *gin.Context)
}

type EndpointsImpl struct {
	Endpoints
	authService       services.AuthService
	genericService    services.GenericService
	httpStatusPackage base.HttpPackage
}

func NewEndpointsImpl(authService services.AuthService, genericService services.GenericService, httpStatusPackage base.HttpPackage) *EndpointsImpl {
	return &EndpointsImpl{
		authService: authService,
		genericService: genericService,
		httpStatusPackage: httpStatusPackage,
	}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	userGroup := router.Group("/user")
	userGroup.POST("/register", e.register)
	userGroup.POST("/login", e.login)
	return userGroup
}
