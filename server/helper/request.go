package helper

import "errors"

type LinkRequest struct {
	Url string
}

func (b *LinkRequest) Validade() error {
	if b.Url == "" {
		return errors.New("param `url` cannot be blank")
	}

	return nil
}