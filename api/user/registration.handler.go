package user

import (
	"DulceDayServer/api/base"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerParameter struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email string `json:"email,omitempty"`
}

type registerResponse struct {
	base.Response
}

// @Summary 注册
// @Produce json
// @Param username body string false "唯一的用户名，类似推特中 @ 后面的以及微信号"
// @Param password body string true "密码"
// @Param email body string false "邮箱地址"
// @Success 200 {object} registerResponse 注册成功
// @Failure 400 {object} base.Response 登陆失败, 信息不合规
// @Router /v1/register [post]
func (e *EndpointsImpl) register(context *gin.Context) {
	parameter := registerParameter{}
	if context.BindJSON(&parameter) == nil {
		email, username, password := parameter.Email, parameter.Username, parameter.Password
		user := &models.User{
			Username: username,
			Password: password,
			Email:    email,
		}
		// 检查用户是否存在, 并且检查用户是否合规
		if user.ValidateNew() && !e.service.CheckUserExisting(user) {
			// 生成唯一id，可以由自定义修改
			newIdentifier := e.service.GenerateUserIdentifier()
			user.Identifier = newIdentifier
			// 存储用户
			user = e.service.NewUser(user)
			// 用户创建成功
			context.JSON(http.StatusCreated, registerResponse{
				Response: base.Response{
					Code: 2001,
					Message: "注册成功",
				},
			})
		} else {
			// 如果用户已经存在或者不合规 (详见 user.Validate)
			context.JSON(http.StatusBadRequest, registerResponse{
				Response: base.Response{
					Code: 4000,
					Message: "用户已经存在或用户参数不符合标准",
				},
			})
		}
	} else {
		context.JSON(http.StatusBadRequest, registerResponse{
			Response: base.Response{
				Code: 4001,
				Message: "缺少必要的参数用于注册新的用户",
			},
		})
	}
}