package model

import (
	"bjdaos/pegasus/pkg/reporter/types"
	strutil "github.com/1851616111/util/strings"
)

func parseSingles(s *string) []types.Single {
	l := []types.Single{}
	reduceM := map[string]struct{}{}
	sl := strutil.ClipDBArray(s)
	for _, str := range sl {
		i := types.Single{}
		items := strutil.ClipDBObject(&str)
		if _, exist := reduceM[items[0]]; exist {
			continue
		}

		i.Image = items[0]
		l = append(l, i)
		reduceM[items[0]] = struct{}{}
	}

	return l
}
