package helper

import (
	"errors"
	"fmt"
	"net/http"
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

func MethodAllowed(httpMethod string) error{
	if httpMethod == http.MethodGet || httpMethod == http.MethodPost {
		return nil
	}

	return fmt.Errorf("method '%s' not allowed", httpMethod)
}