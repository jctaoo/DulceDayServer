package api

import (
	"DulceDayServer/api/moment"
	"DulceDayServer/api/static_storage"
	"DulceDayServer/api/user"
	"DulceDayServer/api/user_profile"
	"github.com/gin-gonic/gin"
)

type SiteEndpoints struct {
	UserEndpoints          user.Endpoints
	UserProfileEndpoints   user_profile.Endpoints
	StaticStorageEndpoints static_storage.Endpoints
	MomentEndpoints        moment.Endpoints
}

func (se SiteEndpoints) RouteGroups(router *gin.RouterGroup) []*gin.RouterGroup {
	return []*gin.RouterGroup{
		se.UserEndpoints.MapHandlersToRoutes(router),
		se.UserProfileEndpoints.MapHandlersToRoutes(router),
		se.StaticStorageEndpoints.MapHandlersToRoutes(router),
		se.MomentEndpoints.MapHandlersToRoutes(router),
	}
}
