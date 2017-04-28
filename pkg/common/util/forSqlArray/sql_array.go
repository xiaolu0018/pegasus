package forSqlArray

import (
	"strconv"
	"strings"
)

func SqlstringToStrings(str string) []string {
	if len(str) < 3 {
		return nil
	}
	if str[:1] != "{" && str[(len(str)-1):] != "}" {
		return nil
	}
	return strings.Split(str[1:len(str)-1], ",")
}

func SqlstringToints(str string) ([]int, error) {
	if len(str) < 3 {
		return nil, nil
	}
	if str[:1] != "{" && str[(len(str)-1):] != "}" {
		return nil, nil
	}
	intstrings := strings.Split(str[1:len(str)-1], ",")
	ints := make([]int, len(intstrings))
	var err error
	for k, v := range intstrings {
		if ints[k], err = strconv.Atoi(v); err != nil {
			return nil, err
		}
	}
	return ints, nil
}
