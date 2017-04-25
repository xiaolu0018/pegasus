package strings

import (
	"strings"
)

var DB_ARRAY_MID string = `^,^`
var DB_ARRAY_PERFIX string = `[^[`
var DB_ARRAY_SUFFIX string = `]^]`

var DB_OBJECT_MID string = `^:^`
var DB_OBJECT_PERFIX string = `{^{`
var DB_OBJECT_SUFFIX string = `}^}`

var DB_ARRAY_MID_2 string = `%,%`
var DB_ARRAY_PERFIX_2 string = `[%[`
var DB_ARRAY_SUFFIX_2 string = `]%]`

var DB_OBJECT_MID_2 string = `%:%`
var DB_OBJECT_PERFIX_2 string = `{%{`
var DB_OBJECT_SUFFIX_2 string = `}%}`

func Clip(s *string, left, mid, right string) []string {
	if s == nil {
		return nil
	}

	result := strings.TrimSpace(*s)

	if left != "" {
		result = strings.TrimPrefix(result, left)
	}

	if right != "" {
		result = strings.TrimSuffix(result, right)
	}

	if len(result) == 0 {
		return nil
	}

	return strings.Split(result, mid)
}

func ClipDBArray(s *string) []string {
	if s == nil || len(strings.TrimSpace(*s)) == 0 {
		return nil
	}

	return Clip(s, DB_ARRAY_PERFIX, DB_ARRAY_MID, DB_ARRAY_SUFFIX)
}

func ClipDBObject(s *string) []string {
	if s == nil || len(strings.TrimSpace(*s)) == 0 {
		return nil
	}
	return Clip(s, DB_OBJECT_PERFIX, DB_OBJECT_MID, DB_OBJECT_SUFFIX)
}

func ClipDBArray2(s *string) []string {
	if s == nil || len(strings.TrimSpace(*s)) == 0 {
		return nil
	}

	return Clip(s, DB_ARRAY_PERFIX_2, DB_ARRAY_MID_2, DB_ARRAY_SUFFIX_2)
}

func ClipDBObject2(s *string) []string {
	if s == nil || len(strings.TrimSpace(*s)) == 0 {
		return nil
	}
	return Clip(s, DB_OBJECT_PERFIX_2, DB_OBJECT_MID_2, DB_OBJECT_SUFFIX_2)
}
