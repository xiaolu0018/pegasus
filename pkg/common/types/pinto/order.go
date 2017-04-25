package types

type Order struct {
	Code string `json:"code"`
	Name string `json:"name"`

	SalesManCode string `json:"salesman_code"`
	SalesManName string `json:"salesman_name"`
	SalesManTel  string `json:"salesman_tel"`
	Comments     string `json:"comments"`
}

type Group struct {
}

type Plan struct {
	Code string `json:"code"`
}

type PlanGroup struct {
	Code string `json:"code"`
	Name string `json:"name"`
	//plan_group_addsale_type
	Allow_AddPerson   string `json:"allow_add_person"`
	ChargeType        string `json:"charge_type"`
	AddItemChargeType string `json:"add_item_charge_type"`
}
