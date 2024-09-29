package middleware

import (
	"fmt"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/pkg/pb"
	"github.com/gin-gonic/gin"
)

// Get user info from middleware, deserialize user info from access token
func GetUserFromMiddleware(c *gin.Context) (*pb.User, error) {
	value, oke := c.Get(CURRENT_USER)
	if !oke {
		return nil, fmt.Errorf("you are not logged in")
	}
	user := value.(*pb.User)
	return user, nil
}
