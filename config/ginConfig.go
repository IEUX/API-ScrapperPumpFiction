package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	check "goScrapper/modules/utils"
	"net/http"
)

func CustomRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(GinLogsParams()))
	router.Use(GinCustomRecovery())
	err := router.SetTrustedProxies([]string{})
	check.Err(err)
	return router
}

func GinLogsParams() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%s] -> %s by %s in %s. Status[%d]:%s\n",
			param.TimeStamp.Format("15:04:05"),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Latency,
			param.StatusCode,
			param.ErrorMessage,
		)
	}
}

func GinCustomRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
