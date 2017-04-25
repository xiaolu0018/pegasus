package model

import (
	"strings"
)

var toRemove []string = []string{`"`, `\`, `\"`, `,`}

func unquote(s string) string {
	var i int
	for ; i < len(toRemove); i++ {
		if strings.HasPrefix(s, toRemove[i]) {
			s = strings.TrimPrefix(s, toRemove[i])
			i = -1
		}
	}

	for i = 0; i < len(toRemove); i++ {
		if strings.HasSuffix(s, toRemove[i]) {
			s = strings.TrimSuffix(s, toRemove[i])
			i = -1
		}
	}

	return s
}
