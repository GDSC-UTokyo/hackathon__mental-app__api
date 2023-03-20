package auth

import (
	"cmd/app/model"
	"cmd/app/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		if !strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Authorization field format is invalid"})
		}

		idToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
		token, err := app.VerifyIDToken(ctx, idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not permitted"})
		}

		user := model.User{}
		if err := user.GetUserByUId(utils.Hash(token.UID)); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not permitted"})
		}

		userId := user.Id
		c.Set("userId", userId)
		c.Next()
	}
}
