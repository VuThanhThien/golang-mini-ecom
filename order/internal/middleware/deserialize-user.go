package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/gateway/user/grpc"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/order/utils"
	"github.com/gin-gonic/gin"
)

const (
	CURRENT_USER string = "currentUser"
)

func DeserializeUser(userGateway grpc.IUserGateway) gin.HandlerFunc {
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not logged in"})
			return
		}

		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		userID, err := strconv.ParseUint(fmt.Sprint(sub), 10, 32)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "invalid user ID"})
			return
		}

		user, err := userGateway.Get(ctx, uint(userID))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "the user belonging to this token no longer exists"})
			return
		}

		ctx.Set(CURRENT_USER, user)
		ctx.Next()
	}
}
