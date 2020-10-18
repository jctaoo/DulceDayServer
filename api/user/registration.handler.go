package user

import (
	"DulceDayServer/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerParameter struct {
	Username string `json:"username" binding:"gte=1,required" example:"bob"`              // 用户名
	Password string `json:"password" binding:"required,min=8,max=18" example:"qwerty123"` // 密码
	Email    string `json:"email" binding:"email,required" example:"haha@test.com"`       // 邮箱
}

type registerResponse struct {
	common.BaseResponse
}

// @Summary 注册
// @Produce json
// @Param user body registerParameter true "参数"
// @Success 200 {object} registerResponse 注册成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Router /user/register [post]
func (e *EndpointsImpl) register(context *gin.Context) {
	parameter := registerParameter{}
	if err := context.ShouldBindJSON(&parameter); err == nil {
		email, username, password := parameter.Email, parameter.Username, parameter.Password
		if !e.service.CheckUserExisting(username, email) {
			e.service.Register(username, email, password)
			context.JSON(http.StatusCreated, registerResponse{
				BaseResponse: common.BaseResponse{
					Code:    2001,
					Message: "注册成功",
				},
			})
		} else {
			context.JSON(http.StatusBadRequest, registerResponse{
				BaseResponse: common.BaseResponse{
					Code:    4000,
					Message: "用户已经存在",
				},
			})
		}
	} else {
		common.HandleHttpErr(err, context)
	}
}
