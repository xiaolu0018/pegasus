package model

import (
	"192.168.199.199/bjdaos/totem/pkg/types"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"regexp"
)

func InjectionPrevention(to_match_str string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|
  insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		return false
	}
	return re.MatchString(to_match_str)
}

func GetPagesNumLimitPageNo(page_no int, total_row int) types.PagesNumLimitPageNo {
	var limit int
	page_size := 20

	total_pages_num := total_row / page_size
	if (total_row % page_size) != 0 {
		total_pages_num++
	}
	if total_pages_num == 1 {
		limit = total_row
	} else if page_no == total_pages_num {
		limit = total_row - (total_pages_num-1)*page_size
	} else {
		limit = page_size
	}
	Rets := &types.PagesNumLimitPageNo{
		Total_PagesNum: total_pages_num,
		Limit:          limit,
		Page_No:        page_no,
	}
	return *Rets
}

func GetQueryAll(page int, ex_no string, name string, sex string, alreadyReport bool, begintime string, endtime string, hos_code string) (*types.PageAndData, error) {

	var totalnums int
	var rowsdate []types.PrintInfo

	if len(name) != 0 {
		if InjectionPrevention(name) == true {
			return nil, errors.New("EXIST SQL injection attack")
		}
		name = fmt.Sprintf("AND p.name LIKE '%%%s%%'", name)
	}

	var statusReport string
	if alreadyReport {
		statusReport = "AND e.status > 1080 AND e.status <> 1999"
	} else {
		statusReport = "AND e.status == 1080"
	}

	if len(ex_no) != 0 {
		if InjectionPrevention(ex_no) == true {
			return nil, errors.New("EXIST SQL injection attack")
		}
		ex_no = fmt.Sprintf("AND examination_no = '%s'", ex_no)
	}

	if len(sex) != 0 {
		sex = fmt.Sprintf("AND p.sex = '%s'", sex)
	}

	if len(begintime) != 0 {
		begintime = fmt.Sprintf("AND e.checkupdate >= '%s' ", begintime)
	}

	if len(endtime) != 0 {
		endtime = fmt.Sprintf("AND e.checkupdate <= '%s 23:59:59'", endtime)
	}

	if err := DB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM examination AS e, person AS p
		WHERE e.person_code = p.person_code AND e.hos_code = '%s'
		%s %s %s %s %s %s`,
		hos_code, statusReport, sex, name, ex_no, begintime, endtime)).Scan(&totalnums); err != nil {
		glog.Errorf("GetQueryAll: sql return err %v\n", err)
		return nil, err
	}

	iRet := GetPagesNumLimitPageNo(page, totalnums)
	rows, err := DB.Query(fmt.Sprintf(`SELECT e.examination_no, e.is_group, e.enterprise_name, e.checkupdate, p.name, p.sex, p.card_no, es.examination_status_value
		FROM examination AS e , person AS p, examination_status AS es
		WHERE e.person_code = p.person_code and e.status = CAST(es.examination_status_code AS INTEGER) and e.hos_code= '%s'
		%s %s %s %s %s %s
		order by e.checkupdate desc
		limit '%d' offset '%d'`, hos_code, sex, statusReport, name, ex_no, begintime, endtime, iRet.Limit, (iRet.Page_No-1)*20))
	if err != nil {
		glog.Errorf("GetQueryAll: sql return err %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var row = new(types.PrintInfo)
		if err = rows.Scan(&row.Ex_No, &row.Group, &row.Enterprise, &row.Ex_CkDate, &row.Name, &row.Sex, &row.CardNo, &row.Status); err != nil {
			return nil, err
		}

		rowsdate = append(rowsdate, *row)
	}

	if rows.Err() != nil {
		glog.Errorf("GetQueryAll: sql rows.Err() %v\n", err)
		return nil, rows.Err()
	}

	Rets := &types.PageAndData{
		//Page: iRet.Total_PagesNum,
		Page: totalnums,
		Data: rowsdate,
	}

	return Rets, nil
}
