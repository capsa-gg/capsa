package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const hstsSeconds = 60 * 60 * 24 * 365 // 1 year

// SecurityHeadersMiddleware is middleware to add the security headers to each request header.
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Strict-Transport-Security", fmt.Sprintf("max-age=%d; includeSubDomains; preload", hstsSeconds))
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("Referrer-Policy", "strict-origin")

		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Frame-Options", "DENY")

		c.Next()
	}
}
