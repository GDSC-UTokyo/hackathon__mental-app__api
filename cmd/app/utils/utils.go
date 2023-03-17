package utils

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func GenerateId() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
