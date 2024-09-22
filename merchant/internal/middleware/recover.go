package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				operationID, _ := c.Get(OperationIDKey)
				log.Error().Msgf("[%v] %v", operationID, err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"Error": "Internal Server Error",
				})
				return
			}
		}()
		c.Next()
	}
}
