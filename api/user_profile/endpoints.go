package user_profile

import (
	"DulceDayServer/api/common"
	serviceAuth "DulceDayServer/services/auth"
	"DulceDayServer/services/static_storage"
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
	service       user_profile.Service
	userService   serviceAuth.Service
	staticStorage static_storage.Service
}

func NewEndpointsImpl(service user_profile.Service, userService serviceAuth.Service, staticStorage static_storage.Service) *EndpointsImpl {
	return &EndpointsImpl{service: service, userService: userService, staticStorage: staticStorage}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	userGroup := router.Group("/user/profile")
	userGroup.GET("/:Username", e.getProfile)

	userGroup.GET(
		"/",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		e.getSelfProfile,
	)
	userGroup.PATCH(
		"/update",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		e.updateProfile,
	)
	userGroup.PATCH(
		"/update/avatar",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		e.updateAvatar,
	)

	return userGroup
}
