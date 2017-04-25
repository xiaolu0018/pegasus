package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	strutil "github.com/1851616111/util/strings"
)

func parseSingles(s *string) []types.Single {
	l := []types.Single{}
	sl := strutil.ClipDBArray(s)
	for _, str := range sl {
		i := types.Single{}
		items := strutil.ClipDBObject(&str)
		i.CheckupName = items[0]
		i.Image = items[1]
		l = append(l, i)
	}

	return l
}
