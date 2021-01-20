package store

import (
	"DulceDayServer/api/common"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type putPurchasesParameter struct {
	Purchases *[]models.PurchaseItem `json:"purchases"` // 商品信息
}

type putPurchasesResponse struct {
	common.BaseResponse
}

// @Summary 重设内购商品信息
// @Produce json
// @Security ApiKeyAuth
// @Param purchases body putPurchasesParameter true "参数"
// @Success 200 {object} putPurchasesResponse 获取成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 未登录
// @Router /store/purchases [put]
func (e EndpointsImpl) putPurchases(context *gin.Context) {
	parameter := &putPurchasesParameter{}
	if err := context.ShouldBindJSON(parameter); err == nil {
		e.purchasesProvider.PutPurchases(parameter.Purchases)
		context.JSON(http.StatusCreated, putPurchasesResponse{
			BaseResponse: common.BaseResponse{
				Code:    2001,
				Message: "替换商品成功",
			},
		})
	} else {
		common.HandleHttpErr(err, context)
	}
}
