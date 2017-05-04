package appointment

import (
	"time"
)

//
func GetDayFirstSec(t time.Time) int64 {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Unix()
}

func GetDayLastSec(t time.Time) int64 {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999, time.Local).Unix()
}
