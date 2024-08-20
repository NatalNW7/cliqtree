package helper

import (
	"errors"
	"fmt"
)

type LinkRequest struct {
	Url string
}

func (b *LinkRequest) Validade() error {
	if b.Url == "" {
		return errors.New("param `url` cannot be blank")
	}

	return nil
}

func ValidateLinkId(linkID string) (string, error) {
	if linkID == "" {
		return "", fmt.Errorf("linkId provided: %s", linkID)
	}

	return linkID, nil
}