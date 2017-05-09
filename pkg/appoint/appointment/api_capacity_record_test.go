package appointment

import (
	org "bjdaos/pegasus/pkg/appoint/organization"
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
func TestGetBanners(t *testing.T) {
	dbinit()
	_, err := GetBanners()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAppointedCountAfterTodayByOrgCode(t *testing.T) {
	dbinit()
	result, err := GetAppointedCount("0001001")
	if len(result) == 0 {
		t.Fatalf("result should not equit 0 ,%v", err)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCheckupsAppintedCount(t *testing.T) {
	dbinit()
	config, err := org.GetConfigBasic("0001002")
	if err != nil {
		t.Fatal(err)
	}
	result, err := GetCheckupsAppintedCount(*config, "1")
	if len(result) == 0 {
		t.Fatalf("result should not equit 0 ,%v", err)
	}
	if err != nil {
		t.Fatal(err)
	}
}
