package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUnAuth(context *gin.Context) {
	context.JSON(http.StatusUnauthorized, BaseResponse{
		Code: 40001,
		Message: "未登录",
	})
}
