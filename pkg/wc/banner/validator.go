package banner

import (
	"errors"
	"strings"
)

var ErrImageNotFound = errors.New("Imageurl not found")
var ErrRedirectNotFound = errors.New("RedirectUrl not found")
var ErrPosInvlid = errors.New("pos invlid param")

func (b Banner) Validate() error {
	if b.Pos == 0 {
		return ErrPosInvlid
	}

	if strings.TrimSpace(b.ImageUrl) == "" {
		return ErrImageNotFound
	}

	if strings.TrimSpace(b.RedirectUrl) == "" {
		return ErrRedirectNotFound
	}
	return nil
}
