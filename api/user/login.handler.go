package user

import (
	"DulceDayServer/api/base"
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginParameter struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email string `json:"email,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
}

type loginResponse struct {
	base.Response
	token string
}

// @Summary 登陆
// @Produce json
// @Param username body string false "唯一的用户名，类似推特中 @ 后面的以及微信号"
// @Param password body string true "密码"
// @Param email body string false "邮箱地址"
// @Param device_name body string false "登陆的设备，如果是浏览器，则 '浏览器(通过IP获取的城市名)'"
// @Success 200 {object} loginResponse 登陆成功
// @Failure 400 {object} base.Response 登陆失败, 信息不合规
// @Failure 401 {object} base.Response 登陆失败, 鉴权失败
// @Router /v1/login [post]
func (e *EndpointsImpl) login(context *gin.Context) {
	parameter := loginParameter{}
	if context.BindJSON(&parameter) == nil && (parameter.Username != "" || parameter.Email != "") {
		email, username, password := parameter.Email, parameter.Username, parameter.Password
		deviceName := parameter.DeviceName
		if deviceName == "" {
			deviceName = config.SiteConfig.DefaultDeviceName
		}
		user := &models.User{
			Username: username,
			Email:    email,
		}
		ip := context.ClientIP()
		// 检查用户是否合规
		if user.Validate() {
			// 鉴权
			token, err := e.service.AuthenticateWithPassword(username, email, password, ip, deviceName)
			if err != nil {
				// todo log
				context.JSON(http.StatusUnauthorized, loginResponse{
					Response: base.Response{
						Code: 4000,
						Message: "鉴权错误",
					},
				})
			} else {
				context.JSON(http.StatusOK, loginResponse{
					Response: base.Response{
						Code: 2000,
						Message: "登陆成功",
					},
					token: token,
				})
			}
		} else {
			// 如果用户已经存在或者不合规 (详见 user.Validate)
			context.JSON(http.StatusBadRequest, loginResponse{
				Response: base.Response{
					Code: 4002,
					Message: "用户参数不符合标准",
				},
			})
		}
	} else {
		context.JSON(http.StatusBadRequest, loginResponse{
			Response: base.Response{
				Code: 4003,
				Message: "缺少必要的参数用于登陆",
			},
		})
	}
}