package user_profile

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getProfileResponse struct {
	common.BaseResponse
	Profile models.UserProfile
}

// @Summary 获取登录用户信息
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} getProfileResponse 获取成功
// @Failure 401 {object} common.BaseResponse 获取失败, 授权失败
// @Router /user/profile [get]
func (e *EndpointsImpl) getSelfProfile(context *gin.Context) {
	// todo 如果 context 里没有对应的值会 panic
	if helpers.IsAuth(context) {
		username := helpers.AuthUsername(context)
		userProfile := e.service.GetProfileByUsername(username)
		if !userProfile.IsEmpty() {
			context.JSON(http.StatusOK, getProfileResponse{
				BaseResponse: common.BaseResponse{
					Code: 1000,
					Message: "获取成功",
				},
				Profile: *userProfile,
			})
		} else {
			// 创建对应 UserProfile
			common.HttpLogger(context, nil, gin.H{
				"username": username,
			}).Info(username + "创建新的 UserProfile")
			userProfile = e.service.CreateNewProfile(username)
			context.JSON(http.StatusOK, getProfileResponse{
				BaseResponse: common.BaseResponse{
					Code: 1000,
					Message: "获取成功",
				},
				Profile: *userProfile,
			})
		}
	} else {
		common.HandleUnAuth(context)
	}
}

type getProfilePathParameter struct {
	Username string `json:"username" binding:"required"`
}

// @Summary 获取用户信息
// @Produce json
// @Param username path string false "用户名"
// @Success 200 {object} getProfileResponse 获取成功
// @Router /user/profile/{username} [get]
func (e *EndpointsImpl) getProfile(context *gin.Context) {
	pathParam := getProfilePathParameter{}
	if err := context.ShouldBindUri(pathParam); err == nil {
		username := pathParam.Username
		userProfile := e.service.GetProfileByUsername(username)
		if !userProfile.IsEmpty() {
			context.JSON(http.StatusOK, getProfileResponse{
				BaseResponse: common.BaseResponse{
					Code: 1000,
					Message: "获取成功",
				},
				Profile: *userProfile,
			})
		} else {
			context.JSON(http.StatusNotFound, getProfileResponse{
				BaseResponse: common.BaseResponse{
					Code: 40001,
					Message: "找不到 " + username + " 的 Profile",
				},
			})
		}
	} else {
		common.HandleHttpErr(err, context)
	}
}