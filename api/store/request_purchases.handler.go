package store

import (
	"DulceDayServer/api/common"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestPurchasesResponse struct {
	common.BaseResponse
	Purchases *[]models.PurchaseItem `json:"purchases"` // 商品信息
}

// @Summary 获取所有内购商品
// @Produce json
// @Success 200 {object} requestPurchasesResponse 获取成功
// @Router /store/purchases [get]
func (e EndpointsImpl) requestPurchases(context *gin.Context) {
	purchaseItems := e.purchasesProvider.GetPurchases()
	rsp := requestPurchasesResponse{
		BaseResponse: common.BaseResponse{
			Code: 2000,
			Message: "获取成功",
		},
		Purchases: purchaseItems,
	}
	context.JSON(http.StatusOK, rsp)
}
