package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 20)
	return func(c *gin.Context) {

		if limiter.Allow() {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Limited exceed",
			})
		}

	}
}
