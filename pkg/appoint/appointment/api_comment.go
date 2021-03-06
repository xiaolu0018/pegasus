package appointment

import (
	"bjdaos/pegasus/pkg/appoint/db"
	"database/sql"
	"fmt"
	"github.com/golang/glog"
)

func (c *Comment) Create(appid string) (err error) {
	var tx *sql.Tx
	if tx, err = db.GetDB().Begin(); err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	var setcommentidToAppointmentSQL string
	setcommentidToAppointmentSQL = fmt.Sprintf("UPDATE %s SET commentid = '%s',status = '已评价' WHERE id = '%s'", TABLE_APPOINTMENT, c.ID, appid)
	var createcomment string
	createcomment = fmt.Sprintf("INSERT INTO %s (id,environment,attitude,breakfase,details,conclusion) VALUES($1,$2,$3,$4,$5,$6)", TABLE_Appoint_Comment)
	if _, err = tx.Exec(setcommentidToAppointmentSQL); err != nil {
		glog.Errorln("Comment Create UPDATE err ", err.Error())
		return
	}
	if _, err = tx.Exec(createcomment, c.ID, c.Environment, c.Attitude, c.Breakfast, c.Details, c.Conclusion); err != nil {
		glog.Errorln("Comment Create INSERT err ", err.Error())
		return
	}
	return
}
