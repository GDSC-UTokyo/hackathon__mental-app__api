package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

func Hash(token string) string {
	bytes := sha256.Sum256([]byte(token))
	return hex.EncodeToString(bytes[:])
}
