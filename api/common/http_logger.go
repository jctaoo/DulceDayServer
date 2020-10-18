package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HttpLogger(context *gin.Context, err error, params ...interface{}) *logrus.Entry {
	paramsJson, _ := json.Marshal(params)
	return logrus.WithFields(logrus.Fields{
		"ip":     context.ClientIP(),
		"method": context.Request.Method,
		"uri":    context.Request.RequestURI,
		"params": string(paramsJson),
	}).WithError(err)
}
