package user

import (
	models2 "DulceDayServer/database/models"
	"github.com/gin-gonic/gin"
)

type loginParameter struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email string `json:"email,omitempty"`
}

func (e *EndpointsImpl) login(context *gin.Context) {
	parameter := loginParameter{}
	if context.BindJSON(&parameter) == nil && (parameter.Username != "" || parameter.Email != "") {
		email, username, password := parameter.Email, parameter.Username, parameter.Password
		user := &models2.User{
			Username: username,
			Email:    email,
		}
		// 检查用户是否合规
		if user.Validate() {
			// 检查密码是否正确
			if e.genericService.CheckMatchPassword(user, password) {
				// 登陆成功，颁发 jwt token
				token, err := e.authService.GenerateToken(user)
				if err != nil {
					code, data := e.httpStatusPackage.ServerErrorJSON()
					context.JSON(code, data)
				}
				// 返回 token
				code, data := e.httpStatusPackage.OkJson("登陆成功!", gin.H{
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