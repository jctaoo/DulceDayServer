package store

import (
	"DulceDayServer/api/common"
	"DulceDayServer/database/models"
	serviceAuth "DulceDayServer/services/auth"
	"DulceDayServer/services/store"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	common.BaseEndpoints
	// 获取所有内购商品
	requestPurchases(context *gin.Context)
}

type EndpointsImpl struct {
	Endpoints
	userService       serviceAuth.Service
	purchasesProvider store.PurchasesProvider
}

func NewEndpointsImpl(userService serviceAuth.Service, purchasesProvider store.PurchasesProvider) *EndpointsImpl {
	return &EndpointsImpl{userService: userService, purchasesProvider: purchasesProvider}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	storeGroup := router.Group("/store")
	storeGroup.GET("/purchases", e.requestPurchases)
	storeGroup.PUT(
		"/purchases",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		common.MiddleWareAuthorityLevel(models.AuthorityLevelRoot),
		e.putPurchases,
	)
	return storeGroup
}
