package moment

import (
	"DulceDayServer/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createMomentParameter struct {
	Content string `json:"content" binding:"required,gte=5,lte=250" example:"是第一条动态呀"`
}

type createMomentResponse struct {
	common.BaseResponse
	MomentID string `json:"moment_id"`
}

// @Summary 创建动态
// @Produce json
// @Security ApiKeyAuth
// @Param newMoment body createMomentParameter true "参数"
// @Success 200 {object} createMomentResponse 获取成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 未登录
// @Router /moment/create [post]
func (e EndpointsImpl) createMoment(context *gin.Context) {
	parameter := &createMomentParameter{}
	authDetail := common.GetAuthDetail(context)
	if err := context.ShouldBindJSON(parameter); err == nil {
		mid := e.service.CreateNewMoment(parameter.Content, authDetail.UserIdentifier)
		context.JSON(http.StatusCreated, createMomentResponse{
			BaseResponse: common.BaseResponse{
				Code:    2001,
				Message: "创建成功",
			},
			MomentID: mid,
		})
	} else {
		common.HandleHttpErr(err, context)
	}
}
