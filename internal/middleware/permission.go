package middleware

import (
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ADMIN = "admin"
	USER  = "user"
)

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, oke := c.Get(CURRENT_USER)
		if !oke {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		userRole := value.(models.User).Role

		hasRole := false
		if userRole == requiredRole {
			hasRole = true
		}
		if !hasRole || (userRole != ADMIN && userRole != USER) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You don't have role " + requiredRole,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
