package model

import "strconv"

func GetExaminationStatus(exam_no string) (int, error) {
	var status string
	if err := DB.QueryRow(`SELECT STATUS FROM EXAMINATION WHERE EXAMINATION_NO = $1`, exam_no).Scan(&status); err != nil {
		return -1, err
	}

	return strconv.Atoi(status)
}
