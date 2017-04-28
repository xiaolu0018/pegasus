package types

//this value should be sync with sql/gen_al.sql function checkNull()
const DEFAULT_DB_NULL = "NULL"

type Report struct {
	Person    `json:"person"`
	Info      `json:"info"`
	Sales     `json:"sales"`
	Healthes  `json:"health"`
	Checkups  []Checkup `json:"checkups"`
	Analyse   Analyse   `json:"analyse"`
	FinalExam `json:"final"`
	Images    []ImageItems `json:"images"`
	Singles   []Single     `json:"singles"`
}

type Person struct {
	Ex_No     *string `json:"examination_no"` //体检单号
	Ex_CkDate *string `json:"checkup_date"`   //检查时间
	Ex_Image  *string `json:"image_url"`      //体检表的图片地址

	Name      *string `json:"name"`      //姓名
	Sex       *int    `json:"sex"`       //性别
	Birthday  *string `json:"birthday"`  //生日
	Ex_Age    *int8   `json:"age"`       //年龄
	Nation    *string `json:"nation"`    //民族
	Married   *bool   `json:"married"`   //婚否
	Email     *string `json:"email"`     //邮件
	CardNo    *string `json:"card_no"`   //身份证号
	Address   *string `json:"address"`   //通讯地址
	Cellphone *string `json:"cellphone"` //手机
	Phone     *string `json:"phone"`     //宅电

	Enterprise *string `json:"enterprise"` //所属公司

}

type Info struct {
	Contact_phone *string `json:"contact_phone"` //门店联系电话
}

type Sale struct {
	Department_name string `json:"department_name"` //科室项目
	Checkup_name    string `json:"checkup_name"`    //项目名称
	Checkup_status  string `json:"checkup_status"`  //检查状态
}

//Sales 是[]byte数组， 需要将Sales 反序列化位 []Sale
type Sales struct {
	Sale_Datail []Sale `json:"sale_details"`
}

type Healthes struct {
	//json不支持map的key为结构体
	Healthes_Detail map[string][]Template `json:"health_details"`
}

type Health struct {
	Type int    `json:"type"`
	Name string `json:"name"`
}

type Template struct {
	Code     string `json:"-"`
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
}

//DEP.department_name,
//I.item_name, EX_I.item_value, EX_I.exception_arrow, EX_I.reference_description, I.examination_unit,
//EX_CK.diagnose_result,
//DEP.doctor_sign, M.previous_name username, mm.previous_name

type Checkup struct {
	Department     string `json:"dept,omitempty"`
	Items          []Item `json:"items,omitempty"`
	ShowDiagnose   bool   `json:"showDiagnose"`
	DiagnoseResult string `json:"diagnose,omitempty"`
	DocterSign     string `json:"docter,omitempty"`
	Username       string `json:"user,omitempty"`
	PreviousName   string `json:"pre,omitempty"`
}

func (r *Report) CleanNull() {
	if r.Checkups == nil {
		return
	}

	for id := range r.Checkups {
		r.Checkups[id].cleanNUll()
	}
}

func (c *Checkup) cleanNUll() {
	if c.Items == nil {
		return
	}

	if c.Department == DEFAULT_DB_NULL {
		c.Department = ""
	}
	if c.DiagnoseResult == DEFAULT_DB_NULL {
		c.DiagnoseResult = ""
	}
	if c.DocterSign == DEFAULT_DB_NULL {
		c.DocterSign = ""
	}
	if c.Username == DEFAULT_DB_NULL {
		c.Username = ""
	}
	if c.PreviousName == DEFAULT_DB_NULL {
		c.PreviousName = ""
	}

	for id, item := range c.Items {
		if item.Name == DEFAULT_DB_NULL {
			c.Items[id].Name = ""
		}
		if item.Value == DEFAULT_DB_NULL {
			c.Items[id].Value = ""
		}
		if item.ExceptionArrow == DEFAULT_DB_NULL {
			c.Items[id].ExceptionArrow = ""
		}
		if item.ReferenceDesc == DEFAULT_DB_NULL {
			c.Items[id].ReferenceDesc = ""
		}
		if item.ExaminationUnit == DEFAULT_DB_NULL {
			c.Items[id].ExaminationUnit = ""
		}
	}
}

type Item struct {
	Name            string `json:"name,omitempty"`
	Value           string `json:"value,omitempty"`
	ExceptionArrow  string `json:"arrow,omitempty"`
	ReferenceDesc   string `json:"ref,omitempty"`
	ExaminationUnit string `json:"unit,omitempty"`
}

type Analyse struct {
	ListSpecs []Content `json:"specs"`
	Docter    string    `json:"doctor"`
}

type Content struct {
	Checkup        string `json:"checkup"`
	Advice         string `json:"advice"`
	DiagnoseResult string `json:"result"`
}

//SUBSTR(updatetime, 0, 11), analyzse_doctor, finalexamination FROM examination_analyse_finalexamination
type FinalExam struct {
	Time   string `json:"time"`
	Final  string `json:"final"`
	Doctor string `json:"doctor"`
}

type ImageItems struct {
	CheckupName    string      `json:"checkup"`
	DiagnoseResult string      `json:"result"`
	Images         []string    `json:"images"`
	Items          []ImageItem `json:"items"`
}

type ImageItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Single struct {
	Image string `json:"image"`
}
