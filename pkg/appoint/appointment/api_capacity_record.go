package appointment

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/organization"
	"fmt"
	"sync"
	"time"
)

func GetOffDay(org_code string) (map[string]interface{}, error) {
	cb, err := organization.GetConfigBasic(org_code)
	if err != nil {
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
		r_offday = GetDateAfterTimeNow(cb.OffDays)
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
