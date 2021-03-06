package pinto

import (
	"bjdaos/pegasus/pkg/common/types"
	"bjdaos/pegasus/pkg/common/util/database"
	"fmt"
	"testing"
)

func TestGetExam(t *testing.T) {
	db, _ := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if _, err := GetExam(db, "0001160001405"); err != nil {
		t.Fatal(err)
	}
}

func TestCreateUpdate(t *testing.T) {
	db, _ := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	ex := types.Examination{
		ExaminationNo:   "",
		HosCode:         "0001001",
		CreateTime:      "2017-05-02 14:43:55",
		UpdateTime:      "",
		Status:          "1040",
		PersonCode:      "20161031150737998",
		CheckupDate:     "",
		GuidePaperState: "0",
		ReportGrantType: "0",
	}
	ex.ExaminationNo = GetExaminationNo(db, ex)
	fmt.Println("exNO ", ex.ExaminationNo)

	if err := CreateUpdate(db, ex); err != nil {
		t.Fatal(err)
	}

}
