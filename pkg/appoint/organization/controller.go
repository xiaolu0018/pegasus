package organization

import (
	"bjdaos/pegasus/pkg/appoint/db"
	"fmt"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"time"
)

func AutoDeleteOverdueOffdays() {
	for {
		if time.Now().Hour() == 23 {
			basics, err := GetAllGetConfigBasics()
			if err != nil {
				glog.Errorln("appointmen.changeAppointmentStatus err ", err.Error())
			}

			for _, basic := range basics {
				sqlStr := fmt.Sprintf(`UPDATE %s SET offdays = $1 WHERE org_code = '%s'`, TABLE_ORG_CON_BASIC, basic.Org_Code)
				if _, err := db.GetDB().Exec(sqlStr, pq.Array(deleteOverdueOffdays(basic.OffDays))); err != nil {
					glog.Errorln("appointmen.changeAppointmentStatus err ", err.Error())
				}
			}
		}
		time.Sleep(time.Hour)
	}
}
