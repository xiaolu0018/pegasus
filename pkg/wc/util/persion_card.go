package util

import (
	"strconv"
)

//校验二代身份证

func CheckId(id string) bool {

	if len(id) != 18 {
		return false
	}

	var CheckNo = map[int]int{
		//Y: 0 1 2 3 4 5 6 7 8 9 10
		// 校验码: 1 0 X 9 8 7 6 5 4 3 2
		0:  1,
		1:  0,
		2:  11,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}
	//i位置上的加权因子
	wi := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	var endint, vint, mod11, total int
	var err error
	var i = 0
	for k, v := range id[:17] {
		i++
		if vint, err = strconv.Atoi(string(v)); err != nil {
			return false
		}
		total += vint * wi[k]
	}

	mod11 = total % 11 //前11位之和后mod11的校验数字

	if endint, err = strconv.Atoi(string(id[17:])); err != nil {

	} else {
		if endint == 10 {
			endint = 0
		}
	}
	checkno := CheckNo[mod11]
	if checkno == 11 && id[17:] == "X" {
		return true
	} else if checkno == endint {
		return true
	} else {
		return false
	}
}
