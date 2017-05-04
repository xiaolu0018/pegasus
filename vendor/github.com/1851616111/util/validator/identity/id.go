package identity

import (
	"errors"
	"github.com/1851616111/util/validator"
	"regexp"
)

var Err_Format_Invalid error = errors.New("person card number format invalid")

var cardNoRegexp_15 *regexp.Regexp = regexp.MustCompile(`^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$`)
var cardNoRegexp_18 *regexp.Regexp = regexp.MustCompile(`^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}([0-9]|X)$`)

func Validate(id string) error {
	switch len(id) {
	case 15:
		return validator.Validate(cardNoRegexp_15, id)
	case 18:
		return validator.Validate(cardNoRegexp_18, id)
	default:
		return Err_Format_Invalid
	}
}
