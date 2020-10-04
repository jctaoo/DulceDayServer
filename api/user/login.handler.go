package user

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
)

type loginParameter struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email string `json:"email,omitempty"`
	DeviceName string `json:"device_name"`
}

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
				code, data := e.httpStatusPackage.BadRequestJSON("鉴权错误")
				context.JSON(code, data)
			} else {
				code, data := e.httpStatusPackage.OkJson("登录成功", gin.H{
					"token": token,
				})
				context.JSON(code, data)
			}
		} else {
			// 如果用户已经存在或者不合规 (详见 user.Validate)
			code, data := e.httpStatusPackage.BadRequestJSON("用户参数不符合标准")
			context.JSON(code, data)
		}
	} else {
		code, data := e.httpStatusPackage.BadRequestJSON("缺少必要的参数用于登陆")
		context.JSON(code, data)
	}
}