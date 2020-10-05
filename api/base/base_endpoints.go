package base

import "github.com/gin-gonic/gin"

type Endpoints interface {
	MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup
}