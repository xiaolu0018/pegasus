package util

import "time"

var daysOfMonth = [...]int{
	31,
	28,
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}

func MonthDays(month int, year int) int {
	if month == int(time.February) && isLeap(year) {
		return 29
	}
	return int(daysOfMonth[month+1])
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
