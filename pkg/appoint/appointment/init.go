package appointment

import "time"

func Init() {
	func() {
		time.Sleep(3 * time.Second)
		go changeAppointmentStatus()
		go changeAppoitmentStatusToExaming()
	}()
}
