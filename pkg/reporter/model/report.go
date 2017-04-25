package model

import (
	"fmt"

	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	"github.com/golang/glog"
)

const married_code = 2
const unmarried_code = 1

func GetReporterByExNo(exNo string) (*types.Report, error) {
	ret := &types.Report{}

	var married int

	var sales, health_selected, checkupItems, analyse, finalExam, images, singles *string
	row := DB.QueryRow(fmt.Sprintf(`select EX_NO, EX_CHECKUPDATE, EX_IMAGE, EX_AGE, P_NAME, P_SEX, P_CARDNO,
	  P_BIRTHDAY, P_IFMARRIED, P_EMAIL, P_ADDRESS, P_CELLPHONE, P_PHONE, NATION, ENTERPRISE, CONTACT_PHONE,
	  CK_DETAILS, HEALTH_SELECTED, CK_ITEMS, ANALYSE_ADVICE, FINAL_EXAM, IMAGES, SINGLES from go_report where ex_no = '%s'`, exNo))
	if err := row.Scan(&ret.Ex_No, &ret.Ex_CkDate, &ret.Ex_Image, &ret.Ex_Age, &ret.Name, &ret.Sex, &ret.CardNo,
		&ret.Birthday, &married, &ret.Email, &ret.Address, &ret.Cellphone, &ret.Phone, &ret.Nation, &ret.Enterprise, &ret.Contact_phone,
		&sales, &health_selected, &checkupItems, &analyse, &finalExam, &images, &singles); err != nil {
		glog.Errorf("GetReporterByExNo: sql return err %v\n", err)
		return nil, err
	}

	b := married == married_code
	ret.Married = &b

	ret.Sale_Datail = parseSalesData(sales)
	ret.Healthes_Detail = parseHealthSelectedData(health_selected)
	ret.Checkups = parseCheckupItems(checkupItems)
	ret.Analyse = parseAnalyses(analyse)
	ret.FinalExam = parseFinalExam(finalExam)
	ret.Images = parseImageItems(images)
	ret.Singles = parseSingles(singles)

	ret.CleanNull()
	return ret, nil
}
