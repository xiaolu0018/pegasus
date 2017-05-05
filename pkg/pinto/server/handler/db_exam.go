package handler

import (
	"strings"
	"github.com/golang/glog"
	"bjdaos/pegasus/pkg/pinto/server/db"
)
func GetExaminationStatus(m map[string][]string) map[string]int {
	booknos, ok := m["booknos"]
	if !ok {
		return nil
	}

	if len(booknos) == 0 {
		return nil
	}

	rows, err := db.GetReadDB().Query(`SELECT b.bookno, COALESCE(e.status, 0)
	FROM book_record b LEFT JOIN examination e
	ON b.examination_no=e.examination_no
	WHERE b.bookno IN (%s)"`, strings.Join(booknos, ","))
	if err != nil {
		glog.Errorln("pinto.GetExaminationStatus GetExamStatusHandler err ", err)
		return nil
	}
	defer rows.Close()

	ret := map[string]int{}
	for rows.Next() {
		var bookno string
		var status int
		if err = rows.Scan(&bookno, &status); err != nil {
			return nil
		}

		ret[bookno] = status
	}

	return ret
}
