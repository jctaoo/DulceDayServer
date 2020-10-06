package user

import (
	"DulceDayServer/api/common"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerParameter struct {
	Username string `json:"username" binding:"gte=1,required"`
	Password string `json:"password" binding:"required,min=8,max=18"`
	Email string `json:"email" binding:"email,required"`
}

type registerResponse struct {
	common.BaseResponse
}

// @Summary 注册
// @Produce json
// @Param username body string true "唯一的用户名，类似推特中 @ 后面的以及微信号"
// @Param password body string true "密码"
// @Param email body string true "邮箱地址"
// @Success 200 {object} registerResponse 注册成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Router /v1/register [post]
func (e *EndpointsImpl) register(context *gin.Context) {
	parameter := registerParameter{}
	if err := context.ShouldBindJSON(&parameter); err == nil {
		email, username, password := parameter.Email, parameter.Username, parameter.Password
		user := &models.User{
			Username: username,
			Password: password,
			Email:    email,
		}
		// 检查用户是否存在
		if !e.service.CheckUserExisting(user) {
			// 生成唯一id，可以由自定义修改
			newIdentifier := e.service.GenerateUserIdentifier()
			user.Identifier = newIdentifier
			// 存储用户
			user = e.service.NewUser(user)
			// 用户创建成功
			context.JSON(http.StatusCreated, registerResponse{
				BaseResponse: common.BaseResponse{
					Code: 2001,
					Message: "注册成功",
				},
			})
		} else {
			context.JSON(http.StatusBadRequest, registerResponse{
				BaseResponse: common.BaseResponse{
					Code: 4000,
					Message: "用户已经存在",
				},
			})
		}
	} else {
		common.HandleHttpErr(err, context)
	}
}