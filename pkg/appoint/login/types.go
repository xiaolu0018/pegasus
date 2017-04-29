package login

const TABLE_LOGINUSER = "go_appoint_login_user"

type LoginUser struct {
	LoginAccount string `json:"loginaccount"`
	PassWord     string `json:"password"`
	LoginName    string `json:"loginname"`
	OrgCode      string `json:"orgcode"`
}
