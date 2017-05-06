package organization

func Init() {
	go AutoDeleteOverdueOffdays()
}
