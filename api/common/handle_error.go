package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleHttpErr(err error, context *gin.Context) {
	logrus.WithError(err).Info("http 请求处理过程中发生错误")
	if errs, ok := err.(validator.ValidationErrors); ok {
		context.JSON(http.StatusBadRequest, BaseResponse{
			Code: 4001,
			Message: TranslateValidateErr(errs, context),
		})
	} else {
		context.JSON(http.StatusBadRequest, BaseResponse{
			Code: 4002,
			Message: err.Error(),
		})
	}
}
