package user_profile

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateProfileParameter struct {
	Username string `json:"username" binding:"gte=1" example:"alen"` // 用户名
	Nickname string `json:"nickname" binding:"gte=3" example:"jc😄taoo"` // 昵称
}

type updateProfileResponse struct {
	common.BaseResponse
	Username string `json:"username"` // 更改后的用户名
	Nickname string `json:"nickname"` // 更改后的昵称
}

// @Summary 更新用户信息
// @Produce json
// @Security ApiKeyAuth
// @Param userProfile body updateProfileParameter true "参数"
// @Success 200 {object} updateProfileResponse 获取成功
// @Failure 401 {object} common.BaseResponse 获取失败, 授权失败
// @Router /user/profile/update [put]
func (e *EndpointsImpl) updateProfile(context *gin.Context) {
	var parameter updateProfileParameter
	var username string
	if username = helpers.AuthUsername(context); !helpers.IsAuth(context) || username == "" {
		common.HandleUnAuth(context)
		return
	}
	if err := context.ShouldBindJSON(&parameter); err == nil {
		newProfile := &models.UserProfile{
			Username: parameter.Username,
			Nickname: parameter.Nickname,
		}
		e.service.UpdateProfile(username, newProfile)
		context.JSON(http.StatusOK, updateProfileResponse{
			BaseResponse: common.BaseResponse{
				Code: 2000,
				Message: "修改成功",
			},
			Username: newProfile.Username,
			Nickname: newProfile.Nickname,
		})
	} else {
		common.HandleHttpErr(err, context)
	}
}
