package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	strutil "github.com/1851616111/util/strings"
	"github.com/golang/glog"
)

func parseFinalExam(str *string) types.FinalExam {
	ret := types.FinalExam{}
	sl := strutil.ClipDBObject(str)
	if len(sl) == 3 {
		ret.Time, ret.Final, ret.Doctor = sl[0], sl[1], sl[2]
	} else {
		glog.Errorf("parseFinalExam() with unmatch data %v\n", sl)
	}

	return ret
}
