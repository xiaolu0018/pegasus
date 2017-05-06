package appointment

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/appoint/organization"
	"bjdaos/pegasus/pkg/common/util/timeutil"
)

func GetOffDay(org_code string) (map[string]interface{}, error) {
	cb, err := organization.GetConfigBasic(org_code)
	if err != nil {
		fmt.Println("err", cb)
		return nil, err
	}
	sqlStr := fmt.Sprintf("SELECT date FROM %s WHERE org_code = '%s' AND used = %d", TABLE_CapacityRecords, cb.Org_Code, cb.Capacity)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dates []string
	var date string
	var r_date, r_offday []string
	for rows.Next() {
		if err = rows.Scan(&date); err != nil {
			return nil, err
		}
		dates = append(dates, date)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		r_date = GetDateAfterTimeNow(dates)
		wg.Done()
	}()

	go func() {
		r_offday = GetDateAfterTimeNow(DealOffdays(cb.OffDays))
		wg.Done()
	}()
	wg.Wait()

	capacity_map := make(map[string][]string)
	offday_map := make(map[string][]string)
	for _, v := range r_date {
		if value, ok := capacity_map[v[:7]]; ok {
			value = append(value, v[8:])
			capacity_map[v[:7]] = value
		} else {
			capacity_map[v[:7]] = []string{v[8:]}
		}
	}
	for _, v := range r_offday {
		if value, ok := offday_map[v[:7]]; ok {
			value = append(value, v[8:])
			offday_map[v[:7]] = value
		} else {
			offday_map[v[:7]] = []string{v[8:]}
		}
	}

	result := make(map[string]interface{})

	result["capatityed"] = capacity_map
	result["offdays"] = offday_map

	return result, nil

}

func GetDateAfterTimeNow(dates []string) []string {
	var r_date []string
	for _, d := range dates {
		dtime, err := time.Parse("2006-01-02", d)
		if err != nil {
			continue
		}
		if time.Now().Before(dtime) {
			r_date = append(r_date, d)
		}

	}
	return r_date
}

//处理 休息日的问题，以应对不同的结构 如 星期

func DealOffdays(offdays []string) []string {
	week := make(map[string]bool)
	for _, offday := range offdays {
		if weekstring, ok := timeutil.WeekString[offday]; ok {
			week[weekstring] = true
		}

		if len(offday) > 20 {
			startTime, endTime := offday[:10], offday[len(offday)-10:]
			if periods, err := timeutil.GetAllDayFromTimePeriod(startTime, endTime); err != nil {
				glog.Errorln("appoint.DealOffdays.GetAllDayFromTimePeriod ", err)
				return nil
			} else {
				for _, period := range periods {
					week[period] = true
				}
			}
		}
	}

	months := timeutil.GetAfterMonths(time.Now().Year(), int(time.Now().Month()), 3)

	var resultOffday []string

	for _, ms := range months {
		days := timeutil.MonthDays(ms.Month, ms.Year)
		for i := 1; i < days+1; i++ {
			temp := time.Date(ms.Year, time.Month(ms.Month), i, 0, 0, 0, 0, time.Local)
			if hasWeek, ok := week[temp.Weekday().String()]; ok && hasWeek {
				resultOffday = append(resultOffday, temp.Format("2006-01-02"))
			} else {
				if hasday, ok := week[temp.Format("2006-01-02")]; ok && hasday {
					resultOffday = append(resultOffday, temp.Format("2006-01-02"))
				}
			}
		}
	}

	return resultOffday
}
