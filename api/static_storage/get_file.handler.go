package static_storage

import (
	"DulceDayServer/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @Summary 获取静态资源
// @Param key path string true "资源路径"
// @Router /static/{key} [get]
func (e *EndpointsImpl) getFile(context *gin.Context) {
	param := context.Param("key")
	if strings.HasPrefix(param, "/") {
		param = strings.Replace(param, "/", "", 1)
	}
	target, err := e.staticStorageService.GetFileUrl(param)

	common.HttpLogger(context, err, gin.H{
		"key":    param,
		"target": err,
	}).Debug("静态资源重定向，key:"+param, "目标 url:"+target)

	if err != nil {
		context.String(http.StatusNotFound, "未找到相应资源")
	}

	context.Redirect(http.StatusTemporaryRedirect, target)
}
