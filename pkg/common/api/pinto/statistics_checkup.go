package pinto

import (
	"bjdaos/pegasus/pkg/common/util/timeutil"
	"database/sql"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
	"time"
)

type StatisticsCheckup struct {
	CheckupCode string
	CheckupName string
	Date        string
	Count       int
}

func StatisticsCheckups(db *sql.DB, statistics *ForStatistics) ([]StatisticsCheckup, error) {
	var querycheckups, queryhoscodes string
	if len(statistics.Checkups) > 0 {
		querycheckups = fmt.Sprintf(" AND ec.checkup_code in (%s) ", ArrayTostringForSqlIn(statistics.Checkups))
	} else {
		querycheckups = fmt.Sprint("AND ec.checkup_code <> ''")
	}

	if len(statistics.HosCode) > 0 {
		queryhoscodes = fmt.Sprintf(" AND b.bookorg_code in (%s) ", ArrayTostringForSqlIn(statistics.HosCode))
	}

	sqlStr := fmt.Sprintf(`SELECT COALESCE(b.booktimestamp,''), COALESCE(ec.checkup_code,''),count(*), COALESCE(c.checkup_name,'')
	                              FROM book_record b
	                              	LEFT JOIN examination_checkup ec
	                              		ON (b.examination_no=ec.examination_no AND b.bookorg_code=ec.hos_code)
	                              	LEFT JOIN checkup c
	                                	ON ec.checkup_code = c.checkup_code
	                              WHERE b.booktimestamp BETWEEN '%s' AND '%s' %s %s
	                              GROUP BY ec.checkup_code, b.booktimestamp, c.checkup_name
	                              ORDER BY ec.checkup_code,b.booktimestamp`,
		statistics.StartDate, statistics.EndDate, queryhoscodes, querycheckups)
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s_cs []StatisticsCheckup
	var s_c StatisticsCheckup
	for rows.Next() {
		if err = rows.Scan(&s_c.Date, &s_c.CheckupCode, &s_c.Count, &s_c.CheckupName); err != nil {
			return nil, err
		}
		s_cs = append(s_cs, s_c)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return s_cs, nil
}

type S_CSForXlsx struct {
	CheckupNames []string
	Dates        []string
	Counts       [][]int
}

//因为这个s_cs 已经排了序的
func FilterStatisticsCheckups(f_s *ForStatistics, s_cs []StatisticsCheckup) (*S_CSForXlsx, error) {
	days, err := timeutil.GetAllDayFromTimePeriod(f_s.StartDate, f_s.EndDate)
	if err != nil {
		return nil, err
	}
	s_cForXlsx := S_CSForXlsx{}
	checkupNamesMap := map[string]int{}
	countMap := map[string][]StatisticsCheckup{}

	nameIndex := 0

	for _, s_c := range s_cs {
		if _, ok := checkupNamesMap[s_c.CheckupName]; !ok {
			checkupNamesMap[s_c.CheckupName] = nameIndex
			nameIndex++
		}
		if value, ok := countMap[s_c.CheckupName]; ok {
			value = append(value, s_c)
			countMap[s_c.CheckupName] = value
		} else {
			countMap[s_c.CheckupName] = []StatisticsCheckup{s_c}
		}
	}

	nameIndex = 0

	for k := range checkupNamesMap {
		s_cForXlsx.CheckupNames = append(s_cForXlsx.CheckupNames, k)
	}

	daysMap := map[string]int{}

	for k, day := range days {
		initcount := make([]int, len(checkupNamesMap))
		s_cForXlsx.Dates = append(s_cForXlsx.Dates, day)
		daysMap[day] = k
		s_cForXlsx.Counts = append(s_cForXlsx.Counts, initcount)
	}

	for _, name := range s_cForXlsx.CheckupNames {
		if sub_s_cs, ok := countMap[name]; ok {
			for k := range sub_s_cs {
				if index, ok := daysMap[sub_s_cs[k].Date]; ok {
					if nameIndex, ok = checkupNamesMap[sub_s_cs[k].CheckupName]; ok {
						s_cForXlsx.Counts[index][nameIndex] = sub_s_cs[k].Count
					}
				}
			}
		}
	}
	return &s_cForXlsx, err
}

func XlsxStatistics(s_cs S_CSForXlsx) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = ""
	for _, name := range s_cs.CheckupNames {
		cell = row.AddCell()
		cell.Value = name
	}

	for k, rowCounts := range s_cs.Counts {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = s_cs.Dates[k]
		for _, cellCount := range rowCounts {
			cell = row.AddCell()
			cell.SetInt(cellCount)
		}
	}

	err = file.Save(time.Now().Format(timeutil.FROMAT_YYMMDDHHMMSSsss) + "预约统计.xlsx")
	return err
}

func ArrayTostringForSqlIn(arrs []string) string {
	itmeStr := make([]string, len(arrs))
	for k, arr := range arrs {
		itmeStr[k] = fmt.Sprintf(`'%s'`, arr)
	}
	return strings.Join(itmeStr, ",")
}
