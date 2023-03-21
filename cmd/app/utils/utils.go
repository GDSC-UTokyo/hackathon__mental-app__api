package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"net/http"
	"time"
)

func GenerateId() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func GetValueFromContext(c *gin.Context, key string) (valueString string) {
	value, exists := c.Get(key)
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cant get value from gin.Context"})
	}
	return value.(string)
}
