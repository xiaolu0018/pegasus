package user

import (
	"bjdaos/pegasus/pkg/wc/util"
	"errors"
	"github.com/1851616111/util/validator/mobile"
)

var (
	ErrIDCardInvalid  = errors.New("Idcard invalid")
	ErrMobileInvalid  = errors.New("user mobile invalid")
	ErrNameInvalid    = errors.New("user name invalid")
	ErrIsMarryInvalid = errors.New("user ismarry invalid")
)

func (u User) CreateValidate() (err error) {
	if mobile.Validate(u.Mobile) != nil {
		return ErrMobileInvalid
	}
	if u.Name == "" {
		return ErrNameInvalid
	}

	if u.IsMarry == "" {
		return ErrIsMarryInvalid
	}

	if !util.CheckId(u.CardNo) {
		return ErrIDCardInvalid
	}

	return nil
}
