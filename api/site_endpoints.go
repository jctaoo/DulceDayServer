package api

import (
	"DulceDayServer/api/user"
	"DulceDayServer/api/user_profile"
	"github.com/gin-gonic/gin"
)

type SiteEndpoints struct {
	UserEndpoints user.Endpoints
	UserProfileEndpoints user_profile.Endpoints
}

func (se SiteEndpoints) RouteGroups(router *gin.RouterGroup) []*gin.RouterGroup {
	return []*gin.RouterGroup{
		se.UserEndpoints.MapHandlersToRoutes(router),
		se.UserProfileEndpoints.MapHandlersToRoutes(router),
	}
}