package middleware

import (
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// MaxRequestTimeBeforeLoggingError is the amount in ms requests may take before they are too long and logged by Sentry as taking too long.
	MaxRequestTimeBeforeLoggingError = 2000
)

// GinLogger is Gin middleware to log request metadata.
func GinLogger(rootLogger *zap.Logger) gin.HandlerFunc {
	timeFormat := "02/Jan/2006:15:04:05 -0700"
	httpLog := rootLogger.Named("HttpLogger")

	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		start := time.Now()

		c.Next() // Start handling request

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1_000_000.0))
		status := c.Writer.Status()

		// Generate Zap fields
		reqData := []zap.Field{
			zap.Int("status_code", status),
			zap.Int("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("data_length", c.Writer.Size()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("timestamp", time.Now().Format(timeFormat)),
		}

		// Log if there are errors in the Gin context
		if len(c.Errors) > 0 {
			err := c.Errors.ByType(gin.ErrorTypePrivate).String()
			httpLog.Error(err, reqData...)

			return
		}

		msg := fmt.Sprintf("%d response sent for %s %s", status, method, path)

		// Log server errors
		if status > 499 {
			httpLog.Error(msg, reqData...)

			return
		}

		// Log requests that take too long
		if latency > MaxRequestTimeBeforeLoggingError {
			requestLatencyError := fmt.Sprintf("response sent, but marked as error due to request latency being too long, the latency of %dms exceeded the threshold of %dms", latency, MaxRequestTimeBeforeLoggingError)
			httpLog.Warn(requestLatencyError, reqData...)

			return
		}

		// Log 4xx responses as warnings
		// Check for < 500 not needed because of the > 499 check and return above
		if status > 399 {
			httpLog.Info(msg, reqData...)

			return
		}

		httpLog.Info(msg, reqData...)
	}
}
