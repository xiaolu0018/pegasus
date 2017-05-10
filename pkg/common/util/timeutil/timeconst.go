package timeutil

var daysOfMonth = [...]int{
	31,
	28,
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}

var WeekString = map[string]string{
	"星期日": "Sunday",
	"星期一": "Monday",
	"星期二": "Tuesday",
	"星期三": "Wednesday",
	"星期四": "Thursday",
	"星期五": "Friday",
	"星期六": "Saturday",
}

const (
	FROMAT_DAY  = "2006-01-02"
	FROMAT_YYMMDDHHMMSS = "20060102150405"
	FROMAT_SECOND_= "2006-01-02 15:04:05"
)