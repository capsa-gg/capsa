package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/constants"
)

// VersionHeaderName is the header used to share the Capsa Server version in each HTTP request.
const VersionHeaderName = "X-Capsa-Server-Version"

// ServerVersionMiddleware is middleware to add the version header to each request header.
func ServerVersionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(VersionHeaderName, constants.Version)
		c.Next()
	}
}
