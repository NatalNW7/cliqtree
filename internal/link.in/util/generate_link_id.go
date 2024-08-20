package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var randSeed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateLinkId() string {
	linkId := make([]byte, 6)

	for i := range linkId {
		linkId[i] = charset[randSeed.Intn(len(charset))]
	}

	return string(linkId)
}