package branch

import (
	"errors"
	"strings"
)

var ErrNameNotFound error = errors.New("branch name not found")
var ErrDescNotFound error = errors.New("branch description not found")
var ErrTelNotFound error = errors.New("branch Tel not found")

func (b *Branch) Validate() error {
	if strings.TrimSpace(b.Name) == "" {
		return ErrNameNotFound
	}
	if strings.TrimSpace(b.Desc) == "" {
		return ErrDescNotFound
	}
	if strings.TrimSpace(b.Tel) == "" {
		return ErrTelNotFound
	}

	if err := b.Address.Validate(); err != nil {
		return err
	}

	return nil
}
