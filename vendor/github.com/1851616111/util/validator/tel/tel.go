package tel

import (
	"errors"
	"github.com/1851616111/util/validator"
	"regexp"
)

var Err_Format_Invalid error = errors.New("tel number format invalid")

var telRegexp *regexp.Regexp = regexp.MustCompile(`^$|(0[0-9]{2,3}\-)?([2-9][0-9]{6,7})+(\-[0-9]{1,4})?$`)

func Validate(mobile string) error {
	return validator.Validate(telRegexp, mobile)
}
