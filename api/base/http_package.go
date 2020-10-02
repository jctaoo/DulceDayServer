package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpPackage interface {
	// 请求参数解析错误等回传的数据
	BadRequestJSON(message string) (httpCode int, data gin.H)
	// 服务器内部错误
	ServerErrorJSON() (httpCode int, data gin.H)
	// 创建成功
	CreatedJSON(message string) (httpCode int, data gin.H)
	// 成功
	OkJson(message string, attachData map[string]interface{}) (httpCode int, data gin.H)
}

type HttpPackageImpl struct {

}

func NewHttpStatusPackage() *HttpPackageImpl {
	return &HttpPackageImpl{}
}

func (p *HttpPackageImpl) BadRequestJSON(message string) (httpCode int, data gin.H) {
	return http.StatusBadRequest, gin.H{"message": message}
}

func (p *HttpPackageImpl) ServerErrorJSON() (httpCode int, data gin.H) {
	return http.StatusInternalServerError, gin.H{
		"message": "an internal server error occurred, please try again later.",
	}
}

func (p *HttpPackageImpl) CreatedJSON(message string) (httpCode int, data gin.H) {
	return http.StatusCreated, gin.H{"message": message}
}

func (p *HttpPackageImpl) OkJson(message string, attachData map[string]interface{}) (httpCode int, data gin.H) {
	return http.StatusOK, gin.H{"message": message, "data": attachData}
}
