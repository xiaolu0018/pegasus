package menu

//curl -XPOST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=YYC9J3PXPJ6vZ_zCwAC9cKA2NgtSGrdAo7U_hLlH8G1CXiO65QUUTYwmzrYux2Qw3LcCpeC8GANM3I1ItfzXb0tipnPCnSeMG5gS9XH4D91HRB8Nquhcd_BDrQ4Q_co5LSKjAAAWDV -d  '{
//"button":[
//{
//"type":"click",
//"name":"今日歌曲d",
//"key":"V1001_TODAY_MUSIC"
//},
//{
//"name":"菜单",
//"sub_button":[
//{
//"type":"view",
//"name":"搜索”,
//"url":"http://www.soso.com/"
//},
//{
//"type":"view",
//"name":"视频",
//"url":"http://v.qq.com/"
//},
//{
//"type":"click",
//"name":"赞一下我们",
//"key":"V1001_GOOD"
//}]
//}]
//}'

//参考 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183&token=&lang=zh_CN

//1、click：点击推事件用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互；
//2、view：跳转URL用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取用户基本信息接口结合，获得用户基本信息。
//3、scancode_push：扫码推事件用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。
//4、scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。
//5、pic_sysphoto：弹出系统拍照发图用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。
//6、pic_photo_or_album：弹出拍照或者相册发图用户点击按钮后，微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。
//7、pic_weixin：弹出微信相册发图器用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。
//8、location_select：弹出地理位置选择器用户点击按钮后，微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。
//9、media_id：下发消息（除文本消息）用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户，永久素材类型可以是图片、音频、视频、图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
//10、view_limited：跳转图文消息URL用户点击view_limited类型按钮后，微信客户端将打开开发者在按钮中填写的永久素材id对应的图文消息URL，永久素材类型只支持图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
type ButtonType int

const (
	Click = iota
	View
	ScanCode_Push
	ScanCode_WaitMsg
	Pic_SysPhoto
	Pic_Photo_or_album
	Pic_WeiXin
	Location_Select
	Media_Id
	View_Limited
)

var buttonTypes []string = []string{
	Click: "click", View: "view", ScanCode_Push: "scancode_push", ScanCode_WaitMsg: "scancode_waitmsg",
	Pic_SysPhoto: "pic_sysphoto", Pic_Photo_or_album: "pic_photo_or_album", Pic_WeiXin: "pic_weixin",
	Location_Select: "location_select", Media_Id: "media_id", View_Limited: "view_limited"}

type Button struct {
	Type      string       `json:"type,omitempty"`
	Name      string       `json:"name"`
	Key       string       `json:"key,omitempty"`
	SubButton *[]SubButton `json:"sub_button,omitempty"`
}

type SubButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
	Key  string `json:"key,omitempty"`
}

func NewViewButton(name, url string) SubButton {
	return SubButton{
		Type: buttonTypes[View],
		Name: name,
		URL:  url,
	}
}

func NewTopButton(name string) *Button {
	subs := []SubButton{}
	return &Button{
		Name:      name,
		SubButton: &subs,
	}
}

func (b *Button) AddSub(sub SubButton) *Button {
	*b.SubButton = append(*b.SubButton, sub)
	return b
}
