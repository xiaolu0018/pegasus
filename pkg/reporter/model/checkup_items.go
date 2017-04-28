package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"
	strutil "github.com/1851616111/util/strings"
)

//DEP.department_name,
//I.item_name, EX_I.item_value, EX_I.exception_arrow, EX_I.reference_description, I.examination_unit,
//EX_CK.diagnose_result,
//DEP.doctor_sign, M.previous_name username, mm.previous_name

func parseCheckupItems(data *string) []types.Checkup {
	sl := strutil.ClipDBArray(data)
	ret := []types.Checkup{}
	var oldItemDep string
	for _, s := range sl {
		items := strutil.ClipDBObject(&s)
		if len(items) == 0 {
			continue
		}

		newItemDep := items[0]
		if oldItemDep != newItemDep {
			ck := types.Checkup{
				Department:   items[0],
				Items:        []types.Item{},
				DocterSign:   unquote(items[7]),
				Username:     items[8],
				PreviousName: items[9],
			}

			isDiagnodeHide := items[10] == "1" || items[11] == "78"
			if !isDiagnodeHide {
				ck.ShowDiagnose = true
				ck.DiagnoseResult = unquote(items[6])
			}

			ret = append(ret, ck)

			oldItemDep = newItemDep
		}

		index := len(ret) - 1

		ret[index].Items = append(ret[index].Items, types.Item{
			Name:            unquote(items[1]),
			Value:           unquote(items[2]),
			ExceptionArrow:  unquote(items[3]),
			ReferenceDesc:   unquote(items[4]),
			ExaminationUnit: unquote(items[5]),
		})
	}

	return ret
}
