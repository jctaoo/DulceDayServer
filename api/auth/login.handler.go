package auth

import (
	"DulceDayServer/api/common"
	"DulceDayServer/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginWithEmailParameter struct {
	Password   string `json:"password" binding:"required" example:"qwerty123"`        // 密码
	Email      string `json:"email" binding:"email,required" example:"haha@test.com"` // 邮箱
	DeviceName string `json:"device_name,omitempty" example:"bob的iPhone"`             // 登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”
}

type loginWithUsernameParameter struct {
	Password   string `json:"password" binding:"required" example:"qwerty123"` // 密码
	Username   string `json:"username" binding:"required" example:"bob"`       // 用户名
	DeviceName string `json:"device_name,omitempty" example:"bob的iPhone"`      // 登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”
}

type loginResponse struct {
	common.BaseResponse
	Token string
}

// @Summary 使用邮箱登陆
// @Produce json
// @Param auth body loginWithEmailParameter true "参数"
// @Success 200 {object} loginResponse 登陆成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 鉴权失败
// @Router /auth/login/email [post]
func (e *EndpointsImpl) loginWithEmail(context *gin.Context) {
	parameter := loginWithEmailParameter{}
	if err := context.ShouldBindJSON(&parameter); err == nil {
		email, password := parameter.Email, parameter.Password
		deviceName := parameter.DeviceName
		e.login(context, email, "", password, deviceName)
	} else {
		common.HandleHttpErr(err, context)
	}
}

// @Summary 使用用户名登陆
// @Produce json
// @Param auth body loginWithUsernameParameter true "参数"
// @Success 200 {object} loginResponse 登陆成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 鉴权失败
// @Router /auth/login/username [post]
func (e *EndpointsImpl) loginWithUsername(context *gin.Context) {
	parameter := loginWithUsernameParameter{}
	if err := context.ShouldBindJSON(&parameter); err == nil {
		username, password := parameter.Username, parameter.Password
		deviceName := parameter.DeviceName
		e.login(context, "", username, password, deviceName)
	} else {
		common.HandleHttpErr(err, context)
	}
}

func (e *EndpointsImpl) login(context *gin.Context, email string, username string, password string, deviceName string) {
	if deviceName == "" {
		deviceName = config.SiteConfig.DefaultDeviceName
	}
	ip := context.ClientIP()
	// 鉴权
	token, err := e.service.AuthenticateWithPassword(username, email, password, ip, deviceName)
	if err != nil {
		common.HttpLogger(context, err, gin.H{
			"email":      email,
			"username":   username,
			"password":   password,
			"deviceName": deviceName,
		}).Info("用户登陆时发生鉴权错误")
		context.JSON(http.StatusUnauthorized, loginResponse{
			BaseResponse: common.BaseResponse{
				Code:    4000,
				Message: "鉴权错误",
			},
		})
	} else {
		context.JSON(http.StatusOK, loginResponse{
			BaseResponse: common.BaseResponse{
				Code:    2000,
				Message: "登陆成功",
			},
			Token: token,
		})
	}
}
