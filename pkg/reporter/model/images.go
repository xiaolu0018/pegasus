package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	strutil "github.com/1851616111/util/strings"
	"strings"
)

var nilStruct struct{} = struct{}{}

func parseImageItems(s *string) []types.ImageItems {
	l := []types.ImageItems{}
	sl := strutil.ClipDBArray(s)
	for _, str := range sl {
		i := types.ImageItems{}
		items := strutil.ClipDBObject(&str)
		i.CheckupName = items[0]
		i.DiagnoseResult = items[1]
		i.Images = parseImageUrls(items[2])
		i.Items = parseItems(&items[3])
		l = append(l, i)
	}

	return l
}

func parseImageUrls(urls string) []string {
	images := strings.Split(urls, ";")

	//去重
	m := map[string]struct{}{}
	for _, image := range images {
		m[image] = nilStruct
	}

	ret := []string{}
	for image := range m {
		ret = append(ret, image)
	}
	return ret
}

func parseItems(items *string) []types.ImageItem {
	sl := strutil.ClipDBArray2(items)
	ret := make([]types.ImageItem, len(sl))
	for i, str := range sl {
		kv := strutil.ClipDBObject2(&str)
		if len(kv) == 2 {
			ret[i] = types.ImageItem{Name: kv[0], Value: kv[1]}
		}
	}

	return ret
}
