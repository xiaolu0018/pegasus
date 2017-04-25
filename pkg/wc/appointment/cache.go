package appointment

import (
	//"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
	//"fmt"
	//"time"
	"192.168.199.199/bjdaos/pegasus/pkg/common/cache"
)

//type AppointCache struct {
//	toSyncCh chan struct{}
//
//	appTTLSec                     int64
//	aIDToAppointServerAppointment map[string]interface{}
//	aIDToAppoineForDisplay        map[string]interface{}
//	aIDToUID                      map[string]string
//	appointmentToSecs             map[string]int64 //appointment缓存时间秒数
//}
//
//func (ac *AppointCache) CacheAppintment(a Appointment) error {
//	u, err := user.Get(a.UserID)
//
//	if au, err := CreatAppoint_User(a, u); err != nil {
//		return err
//	} else {
//		ac.aIDToAppoineForDisplay[a.ID.Hex()] = *au
//	}
//	ac.aIDToUID[a.ID.Hex()] = u.ID.Hex()
//	ac.aIDToAppointServerAppointment[a.ID.Hex()] = a
//	ac.appointmentToSecs[a.ID.Hex()] = time.Now().Unix()
//
//	fmt.Println("use", u, err)
//	return nil
//}
//
//func (a *AppointCache) DeleteCache(aid string) {
//	delete(a.aIDToAppoineForDisplay, aid)
//	delete(a.aIDToAppointServerAppointment, aid)
//	delete(a.aIDToUID, aid)
//	delete(a.appointmentToSecs, aid)
//}


var appointCache cache.Cache

