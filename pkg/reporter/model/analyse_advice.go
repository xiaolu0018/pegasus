package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	"github.com/1851616111/util/strings"
)

func parseAnalyses(s *string) types.Analyse {
	contexts := []types.Content{}
	var doctor string
	sl := strings.ClipDBArray(s)
	for _, s := range sl {
		ss := strings.ClipDBObject(&s)
		contexts = append(contexts, types.Content{
			Checkup:        unquote(ss[0]),
			Advice:         unquote(ss[2]),
			DiagnoseResult: unquote(ss[3]),
		})

		if doctor == "" {
			doctor = ss[1]
		}
	}

	return types.Analyse{
		ListSpecs: contexts,
		Docter:    doctor,
	}
}
