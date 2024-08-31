package util

import (
	"math/rand"
	"strings"
	"time"
)

func ValidadeRedirectUrl(redirectUrl string) string {
	if !strings.Contains(redirectUrl, "http"){
		return "https://"+redirectUrl
	}

	return redirectUrl
}

func GenerateLinkId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var randSeed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	linkId := make([]byte, 6)

	for i := range linkId {
		linkId[i] = charset[randSeed.Intn(len(charset))]
	}

	return string(linkId)
}