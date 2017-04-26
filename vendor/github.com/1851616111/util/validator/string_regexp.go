package validator

import (
	"errors"
	"regexp"
)

var Err_Nil_String error = errors.New("string nil")
var Err_Format_Invalid error = errors.New("string format invalid")

func Validate(regexp *regexp.Regexp, str string) error {
	if str == "" {
		return Err_Nil_String
	}

	if !regexp.MatchString(str) {
		return Err_Format_Invalid
	}

	return nil
}
