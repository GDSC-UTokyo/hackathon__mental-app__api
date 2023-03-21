package auth

import (
	"cmd/app/model"
	"cmd/app/utils"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

		uid := token.UID

		user := model.User{}
		if err := user.GetUserByUId(utils.Hash(uid)).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				client, _err := app.Auth(ctx)
				if _err != nil {
					c.AbortWithError(http.StatusInternalServerError, _err)
				}

				u, _err := client.GetUser(ctx, uid)
				if _err != nil {
					c.AbortWithError(http.StatusInternalServerError, _err)
				}

				newUserId := utils.GenerateId()
				user = model.User{
					Id:    newUserId,
					Name:  u.DisplayName,
					Email: u.Email,
					UId:   u.UID,
				}

				if _err := user.CreateUser().Error; _err != nil {
					c.AbortWithError(http.StatusInternalServerError, _err)
				}
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not permitted"})
			}
		}

		userId := user.Id
		c.Set("userId", userId)
		c.Next()
	}
}
