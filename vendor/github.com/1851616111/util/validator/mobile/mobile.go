package mobile

import (
	"errors"
	"github.com/1851616111/util/validator"
	"regexp"
)

var Err_Format_Invalid error = errors.New("mobile number format invalid")

var mobileRegexp *regexp.Regexp = regexp.MustCompile(`^[1][0-9]{10}$`)

func Validate(mobile string) error {
	return validator.Validate(mobileRegexp, mobile)
}
