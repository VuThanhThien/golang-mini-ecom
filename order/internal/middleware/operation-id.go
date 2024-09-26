package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const OperationIDKey = "X-Request-Id"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New()
		c.Set(OperationIDKey, id.String())
		c.Next()
	}
}
