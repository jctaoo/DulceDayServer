package user

import (
	"DulceDayServer/database/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type registerParameter struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

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
			successMessage := fmt.Sprintf("注册成功，Hi %s!", user.Username)
			code, data := e.httpStatusPackage.CreatedJSON(successMessage)
			context.JSON(code, data)
		} else {
			// 如果用户已经存在或者不合规 (详见 user.Validate)
			code, data := e.httpStatusPackage.BadRequestJSON("用户已经存在或用户参数不符合标准")
			context.JSON(code, data)
		}
	} else {
		code, data := e.httpStatusPackage.BadRequestJSON("缺少必要的参数用于注册新的用户")
		context.JSON(code, data)
	}
}