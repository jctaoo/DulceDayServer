package helpers

import (
	"DulceDayServer/api/common"
	"github.com/gin-gonic/gin"
)

func IsAuth(context *gin.Context) bool {
	return context.GetBool(common.KIsAuthKey)
}

func AuthUsername(context *gin.Context) string {
	return context.GetString(common.KAuthUsernameContextKey)
}
