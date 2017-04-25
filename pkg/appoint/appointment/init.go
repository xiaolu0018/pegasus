package appointment

func init() {
	func() {
		go changeAppointmentStatus()
	}()
}
