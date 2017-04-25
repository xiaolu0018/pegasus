package user

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
	"errors"
)

var (
	ErrIDCardInvalid  = errors.New("Idcard invalid")
	ErrMobileInvalid  = errors.New("user mobile invalid")
	ErrNameInvalid    = errors.New("user name invalid")
	ErrIsMarryInvalid = errors.New("user ismarry invalid")
)

func (u User) CreateValidate() (err error) {
	if len(u.Mobile) != 11 {
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
