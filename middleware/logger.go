package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools/config"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		if c.Request.Method != "OPTIONS" && config.OSConfig.Logger.EnabledREQ {
			global.ReqLogger.Sugar().Info(fmt.Printf("%s %s %3d %13v %15s",
				reqMethod,
				reqUri,
				statusCode,
				latencyTime,
				clientIP,
			))
		}
	}
}
