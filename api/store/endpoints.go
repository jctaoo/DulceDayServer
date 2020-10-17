package store

import (
	"DulceDayServer/api/common"
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
	purchasesProvider store.PurchasesProvider
}

func NewEndpointsImpl(purchasesProvider store.PurchasesProvider) *EndpointsImpl {
	return &EndpointsImpl{purchasesProvider: purchasesProvider}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	storeGroup := router.Group("/store")
	storeGroup.GET("/purchases", e.requestPurchases)
	return storeGroup
}
