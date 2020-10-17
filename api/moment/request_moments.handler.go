package moment

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/services/moment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestRecommendMomentsResponse struct {
	common.BaseResponse
	Moments *[]moment.FullMoment `json:"moments"`
}

// @Summary 为路人获取推荐的动态
// @Produce json
// @Success 200 {object} requestRecommendMomentsResponse 获取成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Router /moment/recommend/hot [get]
func (e EndpointsImpl) requestRecommendMoments(context *gin.Context) {
	moments := e.service.GetRecommendMoments()
	context.JSON(http.StatusOK, requestRecommendMomentsResponse{
		BaseResponse: common.BaseResponse{
			Code:    2000,
			Message: "获取成功",
		},
		Moments: moments,
	})
}

// @Summary 获取推荐的动态, 精准推送
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} requestRecommendMomentsResponse 获取成功
// @Failure 401 {object} common.BaseResponse 登陆失败, 未登录
// @Router /moment/recommend [get]
func (e EndpointsImpl) requestRecommendMomentsWithAuth(context *gin.Context) {
	authDetail := helpers.GetAuthDetail(context)
	moments := e.service.GetRecommendMomentsWithUserIdentifier(authDetail.UserIdentifier)
	context.JSON(http.StatusOK, requestRecommendMomentsResponse{
		BaseResponse: common.BaseResponse{
			Code:    2000,
			Message: "获取成功",
		},
		Moments: moments,
	})
}
