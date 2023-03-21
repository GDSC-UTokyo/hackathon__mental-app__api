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

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		if !strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Authorization field format is invalid"})
			return
		}

		idToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
		token, err := app.VerifyIDToken(ctx, idToken)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		uid := token.UID

		user := model.User{}
		err = user.GetUserByUId(uid).Error
		if user.UId == uid {
			// ユーザーがfirebaseにもDBにも登録されている場合
			c.Set("userId", user.Id)
			c.Next()
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// ユーザーがfirebaseに登録されているが、DBには登録されていない場合
			client, err := app.Auth(ctx)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			u, err := client.GetUser(ctx, uid)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			newUserId := utils.GenerateId()
			newUser := model.User{
				Id:    newUserId,
				Name:  u.DisplayName,
				Email: u.Email,
				UId:   u.UID,
			}

			if err := newUser.CreateUser().Error; err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			c.Set("userId", newUser.Id)
			c.Next()
		} else {
			// ユーザーがfirebaseに登録されているが、DB検索中にRecordNotFound以外のエラーが起きた場合
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
