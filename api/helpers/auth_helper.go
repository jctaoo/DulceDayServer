package helpers

import (
	"DulceDayServer/api/common"
	"github.com/gin-gonic/gin"
)

func AuthUsername(context *gin.Context) string {
	return context.GetString(common.KAuthUsernameContextKey)
}

func IsSensitiveAuth(context *gin.Context) bool {
	return context.GetBool(common.KIsSensitiveAuth)
}