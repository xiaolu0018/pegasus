package appointment

import (
	"fmt"

	"time"

	"bjdaos/pegasus/pkg/wc/user"

	"bjdaos/pegasus/pkg/appoint/appointment"
)

func CreatAppoint_User(a Appointment, u user.User) (*Appoint_User, error) {
	var au Appoint_User
	au.BranchName = a.BranchName
	au.Planname = a.PlanName
	au.AppointDate = a.AppointDate
	au.Name = u.Name
	au.Mobile = u.Mobile
	au.CardID = u.CardNo
	return &au, nil
}

func Get_Appoint_Appointment(u user.User, a Appointment) (*appointment.Appointment, error) {
	var appointtimeint int64
	if appointtime, err := time.Parse("2006-01-02", a.AppointDate); err != nil {
		return nil, err
	} else {
		appointtimeint = appointtime.Unix()
	}
	address := fmt.Sprintf("%s-%s-%s-%s", u.Address.Province, u.Address.City, u.Address.District, u.Address.Details)
	a_a := appointment.Appointment{
		ID:          a.ID,
		PlanId:      a.PlanID,
		AppointTime: appointtimeint,
		OrgCode:     a.BranchID,

		CardNo:          u.CardNo,
		CardType:        u.CardType,
		Mobile:          u.Mobile,
		Appointor:       u.Name,
		Appointorid:     u.ID,
		Address:         address,
		MerryStatus:     u.IsMarry,
		Status:          appointment.STATUS_SUCCESS,
		Appoint_Channel: "微信",
		Company:         "",
		Group:           "",
		Remark:          "",
		Operator:        "微信用户",
		OperateTime:     time.Now().Unix(),
		OrderID:         "",
		CommentID:       "",
		AppointedNum:    0,
		IfSingle:        true,
		IfCancel:        false,
	}
	return &a_a, nil
}
