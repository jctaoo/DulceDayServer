package user

import (
	"DulceDayServer/api/common"
	"DulceDayServer/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginWithEmailParameter struct {
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"email,required"`
	DeviceName string `json:"device_name,omitempty"`
}

type loginWithUsernameParameter struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	DeviceName string `json:"device_name,omitempty"`
}

type loginResponse struct {
	common.BaseResponse
	Token string
}

// @Summary 使用邮箱登陆
// @Produce json
// @Param password body string true "密码"
// @Param email body string true "邮箱地址"
// @Param device_name body string false "登陆的设备，如果是浏览器，则 '浏览器(通过IP获取的城市名)'"
// @Success 200 {object} loginResponse 登陆成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 鉴权失败
// @Router /v1/login [post]
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

// @Summary 使用邮箱登陆
// @Produce json
// @Param username body string true "唯一的用户名，类似推特中 @ 后面的以及微信号"
// @Param password body string true "密码"
// @Param device_name body string false "登陆的设备，如果是浏览器，则 '浏览器(通过IP获取的城市名)'"
// @Success 200 {object} loginResponse 登陆成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 鉴权失败
// @Router /v1/login [post]
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
			"email": email,
			"username": username,
			"password": password,
			"deviceName": deviceName,
		}).Info("用户登陆时发生鉴权错误")
		context.JSON(http.StatusUnauthorized, loginResponse{
			BaseResponse: common.BaseResponse{
				Code: 4000,
				Message: "鉴权错误",
			},
		})
	} else {
		context.JSON(http.StatusOK, loginResponse{
			BaseResponse: common.BaseResponse{
				Code: 2000,
				Message: "登陆成功",
			},
			Token: token,
		})
	}
}