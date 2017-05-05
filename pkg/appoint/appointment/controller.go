package appointment

import (
	"fmt"
	"time"

	"bjdaos/pegasus/pkg/appoint/db"
	"github.com/golang/glog"
)

func StartController() {
	appBrokenTicker := time.NewTicker(time.Hour)
	appSyncTicker := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-appBrokenTicker.C:
			if err := breakStatus(); err != nil {
				glog.Errorf("appoint.StartController.breakStatus: err %v\n", err)
			} else {
				glog.Errorln("appoint.StartController.breakStatus: ok")
			}

		case <-appSyncTicker.C:
			if err := syncStatus(); err != nil {
				glog.Errorf("appoint.StartController.syncStatus: err %v\n", err)
			} else {
				glog.Errorln("appoint.StartController.syncStatus: ok")
			}
		}
	}
}

//改变没有应约的状态
func breakStatus() (err error) {
	sqlStr := fmt.Sprintf(`UPDATE %s SET status = '爽约' WHERE status = '%s'`, TABLE_APPOINTMENT, STATUS_SUCCESS)
	_, err = db.GetDB().Exec(sqlStr)
	return
}

func syncStatus() error {
	//只需查体检日期是当天的就行
	orderM, err := getTodayOrgBookOrders()
	if err != nil {
		return err
	}

	for org := range *orderM {
		targetOrg := org
		targetBookNos := (*orderM)[targetOrg]
		if len(targetBookNos) == 0 {
			continue
		}

		bookStatsM, err := listOrgBookStatus(targetOrg.ordAddress, targetBookNos)
		if err != nil {
			glog.Errorf("appoint.appSyncStatus.listOrgBookStatus: err %v\n", err)
			continue
		}

		if err := batchUpdateStatus(bookStatsM); err != nil {
			glog.Errorf("appoint.appSyncStatus.batchUpdateStatus: err %v\n", err)
			continue
		}
	}

	return nil
}
