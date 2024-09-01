package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		operationID, _ := c.Get(OperationIDKey)
		log.Printf("[%v] [%s] %q %v", operationID, c.Request.Method, c.Request.RequestURI, time.Since(start))
		c.Next()
	}
}
