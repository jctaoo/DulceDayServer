package moment

import (
	"DulceDayServer/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type toggleMomentStarPathParameter struct {
	MomentID string `json:"moment_id" binding:"required,startswith=M"`
}

type toggleMomentStarResponse struct {
	common.BaseResponse
	IsStarNow bool `json:"is_star_now"`
}

// @Summary 更改点赞👍
// @Produce json
// @Security ApiKeyAuth
// @Produce json
// @Param MomentID path string true "MomentID"
// @Success 200 {object} toggleMomentStarResponse 更改成功
// @Failure 401 {object} common.BaseResponse 登陆失败, 未登录
// @Router /moment/toggle_star/{MomentID} [put]
func (e EndpointsImpl) toggleMomentStar(context *gin.Context) {
	authDetail := common.GetAuthDetail(context)
	pathParameter := &toggleMomentStarPathParameter{}
	if err := context.ShouldBindUri(pathParameter); err == nil {
		isStarNow := e.service.ToggleStarMoment(pathParameter.MomentID, authDetail.UserIdentifier)
		context.JSON(http.StatusOK, toggleMomentStarResponse{
			BaseResponse: common.BaseResponse{
				Code:    2000,
				Message: "更改成功",
			},
			IsStarNow: isStarNow,
		})
	} else {
		common.HandleHttpErr(err, context)
	}
}
