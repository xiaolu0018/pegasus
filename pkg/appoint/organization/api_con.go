package organization

import (
	"fmt"

	"github.com/lib/pq"

	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/common/util/timeutil"
	"strings"
	"time"
)

func (o *Config_Basic) Create() error {
	sqlStr := fmt.Sprintf(`INSERT INTO %s (ORG_CODE, CAPACITY, WARNNUM, OFFDAYS, AVOIDNUMBERS) VALUES ($1, $2, $3, $4, $5)ON CONFLICT (org_code) DO UPDATE SET warnnum=EXCLUDED.warnnum,capacity=EXCLUDED.capacity,offdays=EXCLUDED.offdays,avoidnumbers=EXCLUDED.avoidnumbers`, TABLE_ORG_CON_BASIC)
	_, err := db.GetDB().Exec(sqlStr, o.Org_Code, o.Capacity, o.WarnNum, pq.Array(o.OffDays), pq.Array(o.AvoidNumbers))
	return err
}

func GetConfigBasic(orgcode string) (*Config_Basic, error) {
	sqlStr := fmt.Sprintf(`SELECT org_code,capacity,warnnum,offdays,avoidnumbers,ip_address FROM %s WHERE org_code = '%s'`, TABLE_ORG_CON_BASIC, orgcode)
	var org_code string
	var capacity, warnnum int
	var ip string
	offDays := pq.StringArray{}
	avoidNumbers := pq.Int64Array{}
	var err error
	if err = db.GetDB().QueryRow(sqlStr).Scan(&org_code, &capacity, &warnnum, &offDays, &avoidNumbers, &ip); err != nil {
		return nil, err
	}

	return &Config_Basic{
		Org_Code:     org_code,
		Capacity:     capacity,
		WarnNum:      warnnum,
		OffDays:      []string(offDays),
		AvoidNumbers: []int64(avoidNumbers),
		IpAddress:    ip,
	}, nil
}
func (o *Config_Special) Create() error {
	sqlStr := fmt.Sprintf(`INSERT INTO %s (ORG_CODE, checkup_code, CAPACITY) VALUES ($1, $2, $3)`, TABLE_ORG_CON_SPECIAL)
	_, err := db.GetDB().Exec(sqlStr, o.Org_Code, o.CheckupCode, o.Capacity)
	return err
}

func BulkInsertSpecial(specails []Config_Special) error {
	sqlStr := fmt.Sprintf(`INSERT INTO %s (ORG_CODE, checkup_code, CAPACITY) VALUES `, TABLE_ORG_CON_SPECIAL)
	var values_specails []string
	for _, specail := range specails {
		values_specails = append(values_specails, fmt.Sprintf("('%s','%s',%d)", specail.Org_Code, specail.CheckupCode, specail.Capacity))
	}

	values := strings.Join(values_specails, ",")

	sqlStr += values
	_, err := db.GetDB().Exec(sqlStr)
	return err
}

func GetAllGetConfigBasics() ([]Config_Basic, error) {
	sqlStr := fmt.Sprintf("SELECT org_code,offdays FROM %s", TABLE_ORG_CON_BASIC)

	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var basics []Config_Basic
	var offdays pq.StringArray
	var org_code string
	for rows.Next() {
		if err = rows.Scan(&org_code, &offdays); err != nil {
			return nil, err
		}
		basics = append(basics, Config_Basic{
			Org_Code: org_code,
			OffDays:  []string(offdays),
		})
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return basics, nil
}

func deleteOverdueOffdays(offdays []string) []string {
	var result []string
	tm := time.Now()
	for _, offday := range offdays {
		if _, ok := timeutil.WeekString[offday]; ok {
			result = append(result, offday)
			continue
		}
		if len(offday) > 20 {
			if offtime, err := time.Parse("2006-01-02", offday[len(offday)-10:]); err == nil {
				if tm.Before(offtime) {
					result = append(result, offday)
				}
			}
		}

	}
	return result
}

func GetCountGroupByOrgCodeFromConfig_Special() (map[string]int, error) {
	sqlStr := fmt.Sprintf("SELECT org_code,count(*) FROM %s GROUP BY org_code", TABLE_ORG_CON_SPECIAL)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var org string
	var num int
	special_num := make(map[string]int)
	for rows.Next() {
		if err = rows.Scan(&org, &num); err != nil {
			return nil, err
		}
		special_num[org] = num
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return special_num, nil
}
