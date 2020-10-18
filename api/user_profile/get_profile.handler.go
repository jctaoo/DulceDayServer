package user_profile

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/services/user_profile"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getProfileResponse struct {
	common.BaseResponse
	Data *user_profile.FullUser `json:"data"`
}

// @Summary 获取登录用户信息
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} getProfileResponse 获取成功
// @Failure 401 {object} common.BaseResponse 获取失败, 授权失败
// @Router /user/profile [get]
func (e *EndpointsImpl) getSelfProfile(context *gin.Context) {
	authDetail := helpers.GetAuthDetail(context)
	userIdentifier := authDetail.UserIdentifier
	fullUser := e.service.GetFullUserByUserIdentifier(userIdentifier)
	context.JSON(http.StatusOK, getProfileResponse{
		BaseResponse: common.BaseResponse{
			Code:    1000,
			Message: "获取成功",
		},
		Data: fullUser,
	})
}

type getProfilePathParameter struct {
	Username string `json:"username" binding:"required"`
}

// @Summary 获取用户信息
// @Produce json
// @Param username path string true "用户名"
// @Success 200 {object} getProfileResponse 获取成功
// @Router /user/profile/{username} [get]
func (e *EndpointsImpl) getProfile(context *gin.Context) {
	pathParam := getProfilePathParameter{}
	if err := context.ShouldBindUri(pathParam); err == nil {
		username := pathParam.Username
		fullUser := e.service.GetFullUserByUsername(username)
		if !fullUser.IsEmpty() {
			context.JSON(http.StatusOK, getProfileResponse{
				BaseResponse: common.BaseResponse{
					Code:    1000,
					Message: "获取成功",
				},
				Data: fullUser,
			})
		} else {
			context.JSON(http.StatusNotFound, getProfileResponse{
				BaseResponse: common.BaseResponse{
					Code:    40001,
					Message: "找不到 " + username + " 的 Profile",
				},
			})
		}
	} else {
		common.HandleHttpErr(err, context)
	}
}
