package user_profile

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	globalHelper "DulceDayServer/helpers"
	"bytes"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type updateAvatarResponse struct {
	common.BaseResponse
}

// @Summary 更新头像 (文件传输 go-swagger 无法胜任，请使用 postman 等工具)
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "头像图片"
// @Success 200 {object} updateAvatarResponse 修改成功
// @Failure 401 {object} common.BaseResponse 获取失败, 未登录
// @Router /user/profile/update/avatar [put]
func (e *EndpointsImpl) updateAvatar(context *gin.Context) {
	// 接收文件
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		common.HandleHttpErr(err, context)
		return
	}

	// 校验文件大小
	limit := config.SiteConfig.AvatarSizeMB
	if (header.Size / 1024.0 / 1024.0) > int64(limit) {
		// 文件超过限制大小
		context.JSON(http.StatusBadRequest, updateAvatarResponse{
			BaseResponse: common.BaseResponse{
				Code:    40002,
				Message: "文件过大，请传输 2MB 以下的文件",
			},
		})
		return
	}

	// 校验文件类型
	data, err := ioutil.ReadAll(file)
	contentType, err := globalHelper.GetFileContentType(data)
	if err != nil || !strings.HasPrefix(contentType, "image") {
		context.JSON(http.StatusBadRequest, updateAvatarResponse{
			BaseResponse: common.BaseResponse{
				Code:    40003,
				Message: "错误的文件格式，请上传图片格式的文件，例如 png/jpg",
			},
		})
		return
	}

	// 生成存放该用户头像的 key
	avatarKey := "avatar/" + uuid.NewV4().String() + ".png"

	// 存放头像图片
	err = e.staticStorage.SaveImage(bytes.NewReader(data), avatarKey)
	if err != nil {
		common.HandleHttpErr(err, context)
		return
	}

	// 存放头像路径到持久化数据库中
	authDetail := helpers.GetAuthDetail(context)
	e.service.UpdateProfileByUserIdentifier(authDetail.UserIdentifier, &models.UserProfile{
		AvatarFileKey: avatarKey,
	})

	context.JSON(http.StatusOK, updateAvatarResponse{
		BaseResponse: common.BaseResponse{
			Code:    2000,
			Message: "修改成功",
		},
	})
}