package user_profile

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateProfileParameter struct {
	Nickname string `json:"nickname" binding:"gte=3" example:"jc😄taoo"` // 昵称
}

type updateProfileResponse struct {
	common.BaseResponse
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
	authDetail := helpers.GetAuthDetail(context)
	if err := context.ShouldBindJSON(&parameter); err == nil {
		newProfile := &models.UserProfile{
			Nickname: parameter.Nickname,
		}
		e.service.UpdateProfileByUserIdentifier(authDetail.UserIdentifier, newProfile)
		context.JSON(http.StatusOK, updateProfileResponse{
			BaseResponse: common.BaseResponse{
				Code:    2000,
				Message: "修改成功",
			},
			Nickname: newProfile.Nickname,
		})
	} else {
		common.HandleHttpErr(err, context)
	}
}
