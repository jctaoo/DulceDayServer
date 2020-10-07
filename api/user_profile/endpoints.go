package user_profile

import (
	"DulceDayServer/api/common"
	"DulceDayServer/services/static_storage"
	serviceUser "DulceDayServer/services/user"
	"DulceDayServer/services/user_profile"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	common.BaseEndpoints
	getProfile(context *gin.Context)
	getSelfProfile(context *gin.Context)
	updateProfile(context *gin.Context)
	updateAvatar(context *gin.Context)
}

type EndpointsImpl struct {
	Endpoints
	service user_profile.Service
	userService serviceUser.Service
	staticStorage static_storage.Service
}

func NewEndpointsImpl(service user_profile.Service, userService serviceUser.Service, staticStorage static_storage.Service) *EndpointsImpl {
	return &EndpointsImpl{service: service, userService: userService, staticStorage: staticStorage}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	userGroup := router.Group("/user/profile")
	userGroup.GET("/:username", e.getProfile)

	userGroup.GET("/", common.MiddleWareAuth(e.userService), e.getSelfProfile)
	userGroup.PUT("/update", common.MiddleWareAuth(e.userService), e.updateProfile)
	userGroup.PUT("/update/avatar", common.MiddleWareAuth(e.userService), e.updateAvatar)

	return userGroup
}

