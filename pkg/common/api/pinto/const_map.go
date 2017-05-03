package pinto

var IdCardToCode = map[string]string{
	"身份证": "01",
	"军官证": "03",
	"警察证": "02",
	"护照":  "05",
	"其他":  "04",
}

var SexToCode = map[string]int{
	"男": 1,
	"女": 2,
}

var MarryToCode = map[string]string{
	"未婚": "1",
	"已婚": "2",
}
