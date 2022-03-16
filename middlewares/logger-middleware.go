package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s\t[%s]\t%s\t%s\t%d\t%s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Latency,
			params.StatusCode,
			params.Path,
		)
	})
}
