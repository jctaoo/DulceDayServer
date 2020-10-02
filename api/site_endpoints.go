package api

import (
	"DulceDayServer/api/user"
	"github.com/gin-gonic/gin"
)

type SiteEndpoints struct {
	UserEndpoints user.Endpoints
}

func (se SiteEndpoints) RouteGroups(router *gin.RouterGroup) []*gin.RouterGroup {
	return []*gin.RouterGroup{
		se.UserEndpoints.MapHandlersToRoutes(router),
	}
}