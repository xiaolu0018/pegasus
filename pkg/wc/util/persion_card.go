package util

import (
	"fmt"
	"strconv"
)

//校验二代身份证

//todo 应该加上15位代码补全18位
func CheckId(id string) bool {

	if len(id) != 18 {
		return false
	}

	//i位置上的加权因子
	wi := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	fmt.Println("wi", len(wi))
	var endint, vint, total int
	var err error
	var i = 0
	for k, v := range id[:17] {
		i++
		if vint, err = strconv.Atoi(string(v)); err != nil {
			fmt.Println("zhei")
			return false
		}
		total += vint * wi[k]
	}
	fmt.Println("zheli iiii", i)

	vint = total % 11 //前11位之和后mod11的校验数字

	if endint, err = strconv.Atoi(string(id[17:])); err != nil {

	} else {
		if endint == 10 {
			endint = 0
		}
	}
	fmt.Println(vint, id[17:])
	if vint == 11 && id[17:] == "X" {
		return true
	} else if vint == endint {
		fmt.Println("zhei1")
		return true
	} else {
		fmt.Println("zhei2")
		return false
	}
}
