package timeutil

//import "time"

type YearAndMonth struct {
	Year  int
	Month int
}

// 得出确定月份往后n个月的结构
func GetAfterMonths(year, month, n int) []YearAndMonth {
	yearAndMonths := make([]YearAndMonth, 0, n)
	yearandmonth := YearAndMonth{}
	for i := 0; i < n; i++ {
		if (month + i) > 12 {
			y := year + (month+i)/12

			m := (month + i) % 12
			if (month+i)%12 == 0 {
				m = 12
			}
			yearandmonth.Year, yearandmonth.Month = y, m
		} else {
			yearandmonth.Year, yearandmonth.Month = year, (month + i)
		}
		yearAndMonths = append(yearAndMonths, yearandmonth)
	}
	return yearAndMonths
}
