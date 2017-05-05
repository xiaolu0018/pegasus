package timeutil

import (
	"fmt"
	"time"
)

func MonthDays(month int, year int) int {
	if month == int(time.February) && isLeap(year) {
		return 29
	}
	return int(daysOfMonth[month+1])
}

//暂时只处理 2006-01-02 这样格式的数据
func GetAllDayFromTimePeriod(startTime, endTime string) ([]string, error) {
	var start, end time.Time
	var err error
	if start, err = time.Parse("2006-01-02", startTime); err != nil {
		return nil, err
	}

	if end, err = time.Parse("2006-01-02", endTime); err != nil {
		return nil, err
	}

	if start.After(end) {
		return nil, fmt.Errorf("startTime cannot after endTime")
	}

	if end.Year()-start.Year() > 1 {
		return nil, fmt.Errorf("endTime is too larger than startTime")
	}
	var days []string

	if end.Year() == start.Year() {
		for i := 0; i < end.YearDay()-start.YearDay()+1; i++ {
			days = append(days, start.Add(time.Duration(i*24)*time.Hour).Format("2006-01-02"))
		}
	} else {
		if isLeap(start.Year()) {
			for i := 0; i < 366-start.YearDay()+1; i++ {
				days = append(days, start.Add(time.Hour*time.Duration(i*24)).Format("2006-01-02"))
			}

		} else {
			for i := 0; i < 365-start.YearDay()+1; i++ {
				days = append(days, start.Add(time.Duration(i*24)*time.Hour).Format("2006-01-02"))
			}

		}
		for i := 0; i < end.YearDay(); i++ {
			days = append(days, start.Add(time.Duration(-i*24)*time.Hour).Format("2006-01-02"))
		}
	}

	return days, nil
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
