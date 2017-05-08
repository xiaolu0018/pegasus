package appointment

import (
	"errors"
	"fmt"
	"strings"

	id "github.com/1851616111/util/validator/identity"
	"github.com/1851616111/util/validator/mobile"

	org "bjdaos/pegasus/pkg/appoint/organization"
	"database/sql"
	"github.com/golang/glog"
	"time"
)

var (
	ErrBasicConfigInvalid        = errors.New("ConfigBasic Invalid")
	ErrAppointChannelInvalid     = errors.New("appoint channel invalid")
	ErrAppointTimeInvalid        = errors.New("appint time invalid")
	ErrAppointMerryStatusInvalid = errors.New("appoint merry status invalid")
)

func (a *Appointment) Validate() error {
	if err := a.validatePersonInfo(); err != nil {
		return err
	}

	if err := a.validateOrg(); err != nil {
		glog.Errorf("appointment.validateOrg: get organization by code %s err %v\n", a.OrgCode, err)
		return err
	}

	if err := a.validatePlan(); err != nil {
		glog.Errorf("appointment.validatePlan: get plan by code %s err %v\n", a.OrgCode, err)
		return err
	}

	if err := a.validateAppointInfo(); err != nil {
		return err
	}

	return nil
}

func (a *Appointment) validateOrg() error {
	if len(strings.TrimSpace(a.OrgCode)) == 0 {
		return FieldEmpty("org_code")
	}

	_, err := org.GetOrgByCode(a.OrgCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("organization code=%s not found", a.OrgCode)
		}
		return err
	}

	return nil
}

func (a *Appointment) validatePlan() error {
	if len(strings.TrimSpace(a.PlanId)) == 0 {
		return nil
	}

	pl, err := GetPlanByID(a.PlanId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("plan id=%s not found", pl.ID)
		}
		return err
	}

	return nil
}

func (a *Appointment) validatePersonInfo() error {
	if len(strings.TrimSpace(a.Appointor)) == 0 {
		return FieldEmpty("appointor")
	}

	if len(strings.TrimSpace(a.CardType)) == 0 {
		return FieldEmpty("cardtype")
	}

	if a.CardType != VALIDATE_CARD_TYPE_ID && a.CardType != VALIDATE_CARD_TYPE_PASSPORT && a.CardType != VALIDATE_CARD_TYPE_OFFICER && a.CardType != VALIDATE_CARD_TYPE_POLICE &&
		a.CardType != VALIDATE_CARD_TYPE_OTHER {
		return fmt.Errorf("plan card type =%s invalid", a.CardType)
	}

	if len(strings.TrimSpace(a.CardNo)) == 0 {
		return FieldEmpty("TrimSpace cardno")
	}
	if a.CardType == VALIDATE_CARD_TYPE_ID {
		if err := id.Validate(a.CardNo); err != nil {
			return FieldInvalid("Validate cardno")
		}
	}

	if len(strings.TrimSpace(a.Mobile)) == 0 {
		return FieldEmpty("mobile")
	}

	if err := mobile.Validate(a.Mobile); err != nil {
		return FieldInvalid("mobile")
	}

	if len(strings.TrimSpace(a.MerryStatus)) == 0 {
		return FieldEmpty("merrystatus")
	}

	if a.MerryStatus != VALIDATE_MERRY_NO && a.MerryStatus != VALIDATE_MERRY_YES {
		return ErrAppointMerryStatusInvalid
	}

	return nil
}

func (a *Appointment) isAppointTimeValid(config *org.Config_Basic) bool {
	now, appointTime := time.Now(), time.Unix(a.AppointTime, 0)

	if a.AppointTime <= now.Unix() {
		return false
	}

	if appointTime.Format("2006-01-02") == now.Format("2006-01-02") {
		return false
	}

	if appointTime.After(now.AddDate(0, 2, 0)) {
		return false
	}

	//是否休假
	for _, v := range config.OffDays {
		if appointTime.Format("2006-01-02") == v {
			return false
		}
	}

	return true
}

func (a *Appointment) validateAppointInfo() error {

	if len(strings.TrimSpace(a.Appoint_Channel)) == 0 {
		return FieldEmpty("appoint_channel")
	}

	if a.Appoint_Channel != VALIDATE_CHANNEL_400 && a.Appoint_Channel != VALIDATE_CHANNEL_WC {
		return ErrAppointChannelInvalid
	}

	orgbasic := &org.Config_Basic{}
	var err error
	if orgbasic, err = org.GetConfigBasic(a.OrgCode); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("organization basic config empty")
		}
		fmt.Errorf("appointment.validateAppointInfo() get basic config err %v\n", err)
		return ErrBasicConfigInvalid
	}

	if !a.isAppointTimeValid(orgbasic) {
		return ErrAppointTimeInvalid
	}

	return nil
}

func FieldEmpty(field string) error {
	return fmt.Errorf("object field %s empty", field)
}

func FieldInvalid(field string) error {
	return fmt.Errorf("object field %s invalid", field)
}
