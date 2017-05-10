package pinto

import (
	"bjdaos/pegasus/pkg/common/util/database"
	"fmt"
	"testing"
)

func TestStatisticsCheckups(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	forstatistics := ForStatistics{
		HosCode:   []string{"0001002"},
		StartDate: "2017-05-04",
		EndDate:   "2017-05-20",
	}

	if result, err := StatisticsCheckups(db, &forstatistics); err != nil {
		t.Fatal(err)
	} else {
		XlsxStatistics(result)
	}
}

func TestFilterStatisticsCheckups(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	forstatistics := ForStatistics{
		HosCode:   []string{"0001002"},
		StartDate: "2017-05-04",
		EndDate:   "2017-05-20",
	}

	if result, err := StatisticsCheckups(db, &forstatistics); err != nil {
		t.Fatal(err)
	} else {
		r, err := FilterStatisticsCheckups(&forstatistics, result)
		if err != nil {
			t.Fatal(err)
		}
		if len(r.Counts) != 1 {
			t.Fatalf("result ___", r)
		}
	}
}

func TestXlsxStatistics(t *testing.T) {
	s_cs := []StatisticsCheckup{}
	XlsxStatistics(s_cs)
}

func TestArrArr(t *testing.T) {
	var arrarr [][]int
	for i := 0; i < 10; i++ {
		arrarr = append(arrarr, []int{1, 2})
	}

	arrarr[5][0] = 3
	fmt.Println(arrarr)

}
