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
	m := map[string]*types.Checkup{}
	for _, s := range sl {
		items := strutil.ClipDBObject(&s)
		if len(items) == 0 {
			continue
		}

		_, ok := m[items[0]]
		if !ok {
			m[items[0]] = &types.Checkup{
				Department:     items[0],
				Items:          []types.Item{},
				DiagnoseResult: unquote(items[6]),
				DocterSign:     unquote(items[7]),
				Username:       items[8],
				PreviousName:   items[9],
			}
		}

		dep := m[items[0]]
		dep.Items = append(dep.Items, types.Item{
			Name:            unquote(items[1]),
			Value:           unquote(items[2]),
			ExceptionArrow:  unquote(items[3]),
			ReferenceDesc:   unquote(items[4]),
			ExaminationUnit: unquote(items[5]),
		})
	}

	for _, v := range m {
		ret = append(ret, *v)
	}

	return ret
}
