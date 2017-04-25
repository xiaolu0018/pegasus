package strings

import (
	"strconv"
	"strings"
)

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

func InterceptNumber(str string) []float64 {
	rs := []rune(str)
	ret := []float64{}

	tmp := []rune{}

	for id, r := range rs {
		if isNumberStr(r) {
			tmp = append(tmp, r)
			if id == len(rs)-1 {
				number, _ := strconv.ParseFloat(string(tmp), 64)
				ret = append(ret, number)
			}
		} else {
			if len(tmp) > 0 {
				i, _ := strconv.ParseFloat(string(tmp), 64)
				ret = append(ret, i)
				tmp = []rune{}
			}
		}
	}

	return ret
}

func isNumberStr(r rune) bool {
	return strings.Contains("0123456789.", string(r))
}
