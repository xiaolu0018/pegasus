package user

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
	"errors"
)

var ErrIDCardInvalid = errors.New("Idcard invalid")

func (u User) CreateValidate() (err error) {
	if len(u.Mobile) != 11 {
		return errors.New("user mobile invalid")
	}
	if u.Name == "" {
		return errors.New("user params invalid")
	}

	if !util.CheckId(u.IDCard) {
		return ErrIDCardInvalid
	}

	return nil
}
