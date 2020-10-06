package user_profile

import (
	"DulceDayServer/api/common"
	serviceUser "DulceDayServer/services/user"
	"DulceDayServer/services/user_profile"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	common.BaseEndpoints
	getProfile(context *gin.Context)
	getSelfProfile(context *gin.Context)
	updateProfile(context *gin.Context)
}

type EndpointsImpl struct {
	Endpoints
	service user_profile.Service
	userService serviceUser.Service
}

func NewEndpointsImpl(service user_profile.Service, userService serviceUser.Service) *EndpointsImpl {
	return &EndpointsImpl{service: service, userService: userService}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	userGroup := router.Group("/user/profile")
	userGroup.GET("/", common.MiddleWareAuth(e.userService), e.getSelfProfile)
	userGroup.GET("/:username", e.getProfile)
	userGroup.PUT("/update", common.MiddleWareAuth(e.userService), e.updateProfile)
	return userGroup
}

