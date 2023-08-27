package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/metricsutil"
)

func PrometheusMiddleware(metricsutil metricsutil.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		metricsutil.IncRequestCounter(c.Request.Method, c.FullPath(), fmt.Sprintf("%d", status))
		metricsutil.ObserveRequestDuration(c.Request.Method, c.FullPath(), duration)
	}
}
