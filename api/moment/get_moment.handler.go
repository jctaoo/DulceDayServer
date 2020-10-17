package moment

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/services/moment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getMomentPathParameter struct {
	MomentID string `json:"moment_id" binding:"required,startswith=M"`
}

type getMomentResponse struct {
	common.BaseResponse
	Moment *moment.FullMoment `json:"moment"`
}

// @Summary 获取某个动态的详细信息
// @Produce json
// @Param MomentID path string true "MomentID"
// @Success 200 {object} getMomentResponse 获取成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Router /moment/get/{MomentID} [get]
func (e EndpointsImpl) getMoment(context *gin.Context) {
	pathParameter := &getMomentPathParameter{}
	authDetail := helpers.GetAuthDetail(context)
	if err := context.ShouldBindUri(pathParameter); err == nil {
		momentId := pathParameter.MomentID
		m := e.service.GetMomentByMomentId(momentId, authDetail.UserIdentifier)
		context.JSON(http.StatusOK, getMomentResponse{
			BaseResponse: common.BaseResponse{
				Code:    2000,
				Message: "获取成功",
			},
			Moment: m,
		})
	} else {
		common.HandleHttpErr(err, context)
	}
}
