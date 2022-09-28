package middleware

import (
	"github.com/CountryMarket/CountryMarket-backend/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	l := logger.LogrusLogger
	return func(c *gin.Context) {
		// 开始的时间
		startTime := time.Now()

		// 处理请求
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()

		// 请求IP
		clientIp := c.ClientIP()

		// 日志的格式
		l.Infof("[GIN] | %3d | %13v | %15s | %s | %s | %s |",
			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUri,
			clientIp,
		)
	}
}
