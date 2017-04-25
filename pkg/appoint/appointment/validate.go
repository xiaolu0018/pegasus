package appointment

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/organization"
	"errors"
	"time"
)

var (
	ErrBasicConfigInvalid = errors.New("ConfigBasic Invalid")
	ErrAppointdateInvalid = errors.New("appointdate Invalid")
)

func (a *Appointment) Validate() error {

	if a.Appointor == "" || a.Mobile == "" || a.CardNo == ""{
		return errors.New("params invalid")
	}
	//所选分院是否存在
	orgbasic := &organization.Config_Basic{}
	var err error
	if orgbasic, err = organization.GetConfigBasic(a.OrgCode); err != nil {
		return ErrBasicConfigInvalid
	}

	//所选日期是否合理

	appointTime := time.Unix(a.AppointTime, 0)
	if a.AppointTime < a.OperateTime {
		return ErrAppointdateInvalid
	}

	if appointTime.Format("2006-01-02") == time.Unix(a.OperateTime, 0).Format("2006-01-02") {
		return ErrAppointdateInvalid
	}

	if appointTime.After(time.Unix(a.OperateTime, 0).AddDate(0, 2, 0)) {
		return ErrAppointdateInvalid
	}

	//是否休假
	for _, v := range orgbasic.OffDays {
		if appointTime.Format("2006-01-02") == v {
			return ErrAppointdateInvalid
		}
	}

	return nil

}
