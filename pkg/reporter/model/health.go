package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/types"

	strutil "github.com/1851616111/util/strings"
)

var h2tMappings map[types.Health][]types.Template
var t2hMappings map[string]types.Health

//health to template mapping
func getH2TMappings() (map[types.Health][]types.Template, error) {
	rows, err := DB.Query(`SELECT C.health_type, C.health_name, T.template_code, T.template_name
	FROM con_health C LEFT JOIN con_health_template T ON C.health_type = T.health_type
	WHERE C.group_type=1 and C.status=1 and C.health_name<>'心理压力'
	ORDER BY order_position`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tp int
	var name, tpl_Code, tpl_Name string
	var m map[types.Health][]types.Template = make(map[types.Health][]types.Template)
	for rows.Next() {
		if err = rows.Scan(&tp, &name, &tpl_Code, &tpl_Name); err == nil {
			health, template := types.Health{Type: tp, Name: name}, types.Template{Code: tpl_Code, Name: tpl_Name}
			m[health] = append(m[health], template)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

//template to health mapping
func getT2HMappings() {
	t2hMappings = make(map[string]types.Health)
	if h2tMappings == nil {
		return
	}

	for health, templates := range h2tMappings {
		h := health
		for _, template := range templates {
			t2hMappings[template.Code] = h
		}
	}
}

func parseHealthSelectedData(selected *string) map[string][]types.Template {
	newMappings := copyMappings()

	clips := strutil.ClipDBArray(selected)
	for _, clipS := range clips {
		markTemplateSelected(clipS, newMappings)
	}

	return newMappings
}

func copyMappings() map[string][]types.Template {
	new := make(map[string][]types.Template, len(h2tMappings))

	for k, v := range h2tMappings {
		temp := v
		new[k.Name] = copyTemplates(temp)
	}

	return new
}

func copyTemplates(src []types.Template) []types.Template {
	if src == nil {
		return nil
	}

	dst := make([]types.Template, len(src))
	for i, v := range src {
		dst[i] = v
	}

	return dst
}

func markTemplateSelected(templateCode string, target map[string][]types.Template) {
	health, ok := t2hMappings[templateCode]
	if !ok {
		return
	}

	templates := target[health.Name]

	for id, template := range templates {
		if template.Code == templateCode {
			templates[id].Selected = true
			return
		}
	}
}
