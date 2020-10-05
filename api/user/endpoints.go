package user

import (
	"DulceDayServer/api/base"
	"DulceDayServer/services/user"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	base.Endpoints
	// 登陆
	login(context *gin.Context)
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
	userGroup.POST("/login", e.login)
	return userGroup
}
