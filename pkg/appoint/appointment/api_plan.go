package appointment

import (
	"fmt"

	"database/sql"
	"github.com/lib/pq"

	"bjdaos/pegasus/pkg/appoint/db"
	org "bjdaos/pegasus/pkg/appoint/organization"
	"bjdaos/pegasus/pkg/common/sdk/pinto"
	"errors"
	"github.com/golang/glog"
)

func GetSalesByPlanID(tx *sql.Tx, planid string) ([]string, error) {
	sql := fmt.Sprintf("SELECT sale_codes FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	var itemStr pq.StringArray
	if err := tx.QueryRow(sql).Scan(&itemStr); err != nil {
		return nil, err
	}

	items := []string(itemStr)
	return items, nil
}

func GetPlanByID(planid string) (*Plan, error) {
	pl := Plan{}
	salecodes := pq.StringArray{}
	sql := fmt.Sprintf("SELECT id, name, avatar_img, detail_img, sale_codes, ifshow FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	if err := db.GetDB().QueryRow(sql).Scan(&pl.ID, &pl.Name, &pl.AvatarImg, &pl.DetailImg, &salecodes, &pl.IfShow); err != nil {
		return nil, err
	}

	pl.SaleCodes = []string(salecodes)
	return &pl, nil
}

func GetPlans() ([]Plan, error) {
	ps := make([]Plan, 0)
	sqlStr := fmt.Sprintf("SELECT id,name,avatar_img,detail_img,sale_codes FROM %s", TABLE_PALN)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	salecodes := pq.StringArray{}
	for rows.Next() {
		p := Plan{}
		if err = rows.Scan(&p.ID, &p.Name, &p.AvatarImg, &p.DetailImg, &salecodes); err != nil {
			return nil, err
		}
		p.SaleCodes = []string(salecodes)
		ps = append(ps, p)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return ps, nil
}

func appPlan(tx *sql.Tx, a *Appointment, bc *org.Config_Basic) error {

	sales, err := GetSalesByPlanID(tx, a.PlanId)
	if err != nil {
		glog.Errorf("appoint.addAppointment appPlan: err %v\n", err)
		return err
	}

	if len(sales) == 0 {
		return errors.New("plan sales empty")
	}

	checkups, err := getCheckupsBySales(bc.IpAddress, sales)
	if err != nil {
		glog.Errorf("appoint.addAppointment GetChecupsBySales err ", err.Error())
		return err
	}

	ckLimits, err := GetCheckupLimit(tx, a.OrgCode, checkups)
	if err != nil {
		glog.Errorf("appoint.addAppointment GetCheckupLimit err ", err.Error())
		return err
	}

	checkupUsed, err := GetCheckupsUsed(tx, a.OrgCode, a.AppointDate, checkups)
	if err != nil {
		glog.Errorf("GetItemAppointedNum err ", err.Error())
		return err
	}

	for ck, limit := range ckLimits {
		if used, ok := checkupUsed[ck]; ok {
			if limit <= used {
				return fmt.Errorf(ErrAppointmentString)
			}
		}
	}

	if err := addAppCheckupRecords(tx, a.ID, a.OrgCode, a.AppointDate, checkups); err != nil {
		return err
	}

	a.SaleCodes = sales

	return nil
}

func getCheckupsBySales(ip string, saleCodes []string) ([]string, error) {
	return pinto.NewPintoSDK().GetCheckupCodesBySaleCodes(saleCodes, ip)
}
