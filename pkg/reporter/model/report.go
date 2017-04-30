package model

import (
	"fmt"

	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	"github.com/golang/glog"
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/db"
)

const married_code = 2
const unmarried_code = 1

func GetReporterByExNo(exNo string, sync bool) (*types.Report, error) {
	var sql string
	if sync {
		sql = getSyncSQL(exNo)
	} else {
		sql = fmt.Sprintf(`select EX_NO, EX_CHECKUPDATE, EX_IMAGE, EX_AGE, P_NAME, P_SEX, P_CARDNO,
	  P_BIRTHDAY, P_IFMARRIED, P_EMAIL, P_ADDRESS, P_CELLPHONE, P_PHONE, NATION, ENTERPRISE, CONTACT_PHONE,
	  CK_DETAILS, HEALTH_SELECTED, CK_ITEMS, ANALYSE_ADVICE, FINAL_EXAM, IMAGES, SINGLES from go_report where ex_no = '%s'`, exNo)
	}

	ret := &types.Report{}
	var married int
	var sales, health_selected, checkupItems, analyse, finalExam, images, singles *string
	row := DB.QueryRow(sql)
	if err := row.Scan(&ret.Ex_No, &ret.Ex_CkDate, &ret.Ex_Image, &ret.Ex_Age, &ret.Name, &ret.Sex, &ret.CardNo,
		&ret.Birthday, &married, &ret.Email, &ret.Address, &ret.Cellphone, &ret.Phone, &ret.Nation, &ret.Enterprise, &ret.Contact_phone,
		&sales, &health_selected, &checkupItems, &analyse, &finalExam, &images, &singles); err != nil {
		glog.Errorf("GetReporterByExNo: sql return err %v\n", err)
		return nil, err
	}

	*ret.CardNo = encodeCardNo(*ret.CardNo)

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

func getSyncSQL(exam_no string) string {
	return fmt.Sprintf(`select  EX.examination_no EX_NO , EX.checkupdate EX_CHECKUPDATE, replace(EX.image_url,'\','/') EX_IMAGE, EX.age EX_AGE,
            	P.name P_NAME, P.sex P_SEX, P.card_no P_CARDNO, P.bithday P_BIRTHDAY, P.is_marry P_IFMARRIED, P.email P_EMAIL,
              P.address P_ADDRESS, P.cellphone P_CELLPHONE, P.phone P_PHONE,
            	(SELECT nation_name from nation n where n.nation_code = P.nation_code) NATION,
            	se.enterprise_name ENTERPRISE,
            	(SELECT T2.contact_phone FROM examination T1, print_template T2 WHERE T1.checkup_hoscode = T2.hos_code AND T1.examination_no = '%s') CONTACT_PHONE,
              getCheckupStr('%s') CK_DETAILS,
              getSelectedStr('%s') HEALTH_SELECTED,
              getCheckAndItems('%s') CK_ITEMS,
              getFinalDiagoseStr('%s') ANALYSE_ADVICE,
              genFinalExam('%s') FINAL_EXAM,
              getImageStr('%s') IMAGES,
              getSingles('%s') SINGLES
       FROM examination EX
       LEFT JOIN person P ON EX.person_code = P.person_code
       LEFT JOIN organization o ON EX.hos_code = o.org_code
       LEFT JOIN sale_order so ON so.order_code = EX.group_code
       LEFT JOIN sale_enterprise se ON se.enterprise_code = so.enterprise_code
       WHERE
	        EX.examination_no = '%s'`, exam_no, exam_no, exam_no, exam_no, exam_no, exam_no, exam_no, exam_no, exam_no)
}

func encodeCardNo(cardNo string) string {
	switch len(cardNo) {
	case 15:
		return fmt.Sprintf("%s****%s**%s", cardNo[0:2], cardNo[6:12], cardNo[13:])
	case 18:
		return fmt.Sprintf("%s****%s**%s", cardNo[0:2], cardNo[6:14], cardNo[16:])
	default:
		return cardNo
	}
}

func UpdateStatus(exam_no, status string) error {
	var sql string
	switch status {
	case "printed":
		sql = fmt.Sprintf(`UPDATE examination SET status = %d where examination_no = '%s'`, 1090, exam_no)
	default:
		return fmt.Errorf("unknow status %s\n", status)
	}

	_, err := db.GetWriteDB().Exec(sql)
	return err
}
