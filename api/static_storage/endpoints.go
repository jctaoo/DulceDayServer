package static_storage

import (
	"DulceDayServer/api/common"
	"DulceDayServer/services/static_storage"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	common.BaseEndpoints
	getFile(gin *gin.Context)
}

type EndpointsImpl struct {
	Endpoints
	staticStorageService static_storage.Service
}

func NewEndpointsImpl(staticStorageService static_storage.Service) *EndpointsImpl {
	return &EndpointsImpl{staticStorageService: staticStorageService}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	fileGroup := router.Group("/static")
	fileGroup.GET("/*key", e.getFile)
	return fileGroup
}
