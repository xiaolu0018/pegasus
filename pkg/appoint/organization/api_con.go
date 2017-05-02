package organization

import (
	"fmt"

	"github.com/lib/pq"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
)

func (o *Config_Basic) Create() error {
	sqlStr := fmt.Sprintf(`INSERT INTO %s (ORG_CODE, CAPACITY, WARNNUM, OFFDAYS, AVOIDNUMBERS) VALUES ($1, $2, $3, $4, $5)ON CONFLICT (org_code) DO UPDATE SET warnnum=EXCLUDED.warnnum,capacity=EXCLUDED.capacity,offdays=EXCLUDED.offdays,avoidnumbers=EXCLUDED.avoidnumbers`, TABLE_ORG_CON_BASIC)
	_, err := db.GetDB().Exec(sqlStr, o.Org_Code, o.Capacity, o.WarnNum, pq.Array(o.OffDays), pq.Array(o.AvoidNumbers))
	return err
}

func GetConfigBasic(orgcode string) (*Config_Basic, error) {
	sqlStr := fmt.Sprintf(`SELECT org_code,capacity,warnnum,offdays,avoidnumbers FROM %s WHERE org_code = '%s'`, TABLE_ORG_CON_BASIC, orgcode)
	var org_code string
	var capacity, warnnum int

	offDays := pq.StringArray{}
	avoidNumbers := pq.Int64Array{}
	var err error
	if err = db.GetDB().QueryRow(sqlStr).Scan(&org_code, &capacity, &warnnum, &offDays, &avoidNumbers); err != nil {
		return nil, err
	}

	return &Config_Basic{
		Org_Code:     org_code,
		Capacity:     capacity,
		WarnNum:      warnnum,
		OffDays:      []string(offDays),
		AvoidNumbers: []int64(avoidNumbers),
	}, nil
}
func (o *Config_Special) Create() error {
	sqlStr := fmt.Sprintf(`INSERT INTO %s (ORG_CODE, SALE_CODE, CAPACITY) VALUES ($1, $2, $3)`, TABLE_ORG_CON_SPECIAL)
	_, err := db.GetDB().Exec(sqlStr, o.Org_Code, o.Sale_Code, o.Capacity)
	return err
}
