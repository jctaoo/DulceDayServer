package auth

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sEmailLoginParameter struct {
	VerificationCode string `json:"verificationCode" binding:"required" example:"623597"`   // 验证码
	Email            string `json:"email" binding:"email,required" example:"haha@test.com"` // 邮箱
	DeviceName       string `json:"device_name,omitempty" example:"bob的iPhone"`             // 登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”
}

type sLoginResponse struct {
	common.BaseResponse
	Token string
}

// @Summary 使用邮箱进行敏感登陆验证, 需要事先登录
// @Produce json
// @Security ApiKeyAuth
// @Param auth body sEmailLoginParameter true "参数"
// @Success 200 {object} loginResponse 登陆成功
// @Failure 400 {object} common.BaseResponse 登陆失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 未登录
// @Router /auth/login/sensitive/email [post]
func (e *EndpointsImpl) loginForSensitiveWithEmail(context *gin.Context) {
	parameter := sEmailLoginParameter{}
	authDetail := helpers.GetAuthDetail(context)
	if err := context.ShouldBindJSON(&parameter); err == nil {
		ip := context.ClientIP()
		deviceName := parameter.DeviceName
		tokenStr, err := e.service.AuthenticateForSensitiveVerification(parameter.Email, parameter.VerificationCode, ip, deviceName)
		if err != nil {
			common.HttpLogger(context, err, gin.H{
				"verificationCode": parameter.VerificationCode,
				"email":            parameter.Email,
				"username":         authDetail.Username,
				"deviceName":       deviceName,
			}).Info("用户使用邮箱进行敏感登陆验证时发生鉴权错误")
			context.JSON(http.StatusUnauthorized, sLoginResponse{
				BaseResponse: common.BaseResponse{
					Code:    4000,
					Message: "敏感登陆验证错误",
				},
			})
		} else {
			context.JSON(http.StatusOK, sLoginResponse{
				BaseResponse: common.BaseResponse{
					Code:    2000,
					Message: "登陆成功",
				},
				Token: tokenStr,
			})
		}
	} else {
		common.HandleHttpErr(err, context)
	}
}
