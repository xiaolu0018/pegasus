package error

import "strings"

func ForeignKeyConstraint(err error) bool {
	return strings.Contains(err.Error(), "violates foreign")
}
