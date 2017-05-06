package appointment

import (
	"fmt"
	"testing"
)

func TestGetOffDaye(t *testing.T) {
	dbinit()
	result, err := GetOffDay("0001001")
	fmt.Println("result,err", result, err)
}

func TestDealOffdays(t *testing.T) {
	offday := []string{"星期六", "2017-05-07至2017-05-13", "2017-05-28至2017-06-08"}
	days := DealOffdays(offday)
	fmt.Println("days ,,,", days)
}
