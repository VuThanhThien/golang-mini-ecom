package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/VuThanhThien/golang-gorm-postgres/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/utils"
	"github.com/gin-gonic/gin"
)

const (
	CURRENT_USER string = "currentUser"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		//TODO: also can use redis to cache later
		var user models.User
		result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))

		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no longer exists"})
			return
		}

		ctx.Set(CURRENT_USER, user)
		ctx.Next()
	}
}
