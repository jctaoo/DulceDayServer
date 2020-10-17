package auth

import (
	"DulceDayServer/api/common"
	"DulceDayServer/api/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sEmailRegisterParameter struct {
	Email      string `json:"email" binding:"email,required" example:"haha@test.com"` // 邮箱
	DeviceName string `json:"device_name,omitempty" example:"bob的iPhone"`             // 登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”
}

type sRegisterResponse struct {
	common.BaseResponse
}

// @Summary 敏感注册，用于生成验证码等, 需要事先登录
// @Produce json
// @Security ApiKeyAuth
// @Param auth body sEmailRegisterParameter true "参数"
// @Success 200 {object} loginResponse 敏感注册成功
// @Failure 400 {object} common.BaseResponse 敏感注册失败, 信息不合规
// @Failure 401 {object} common.BaseResponse 登陆失败, 未登录
// @Router /auth/register/sensitive/email [post]
func (e *EndpointsImpl) registerForSensitiveWithEmail(context *gin.Context) {
	parameter := sEmailRegisterParameter{}
	authDetail := helpers.GetAuthDetail(context)
	if err := context.ShouldBindJSON(&parameter); err == nil {
		ip := context.ClientIP()
		verificationCode, err := e.service.PrepareForAuthForSensitiveVerification(
			authDetail.Username,
			parameter.Email,
			ip,
			parameter.DeviceName,
		)
		if err != nil {
			common.HttpLogger(context, err, gin.H{
				"username":   authDetail.Username,
				"email":      parameter.Email,
				"deviceName": parameter.DeviceName,
			}).Info("使用邮箱敏感注册时发生鉴权错误")
			context.JSON(http.StatusUnauthorized, sRegisterResponse{
				BaseResponse: common.BaseResponse{
					Code:    4000,
					Message: "敏感注册失败",
				},
			})
		} else {
			common.HttpLogger(context, nil, gin.H{
				"username":   authDetail.Username,
				"email":      parameter.Email,
				"deviceName": parameter.DeviceName,
			}).Debugf("用户使用邮箱敏感注册， 邮箱为: %s, 验证码为: %s", parameter.Email, verificationCode)
			// todo 发邮件
			context.JSON(http.StatusOK, loginResponse{
				BaseResponse: common.BaseResponse{
					Code:    2000,
					Message: "敏感注册成功，请查收邮件",
				},
			})
		}
	} else {
		common.HandleHttpErr(err, context)
	}
}
