package user

import "github.com/1851616111/util/weichat/errors"

//details https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842&token=&lang=zh_CN
//subscribe	用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
//openid	用户的标识，对当前公众号唯一
//nickname	用户的昵称
//sex	用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
//city	用户所在城市
//country	用户所在国家
//province	用户所在省份
//language	用户的语言，简体中文为zh_CN
//headimgurl	用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
//subscribe_time	用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
//unionid	只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
//remark	公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
//groupid	用户所在的分组ID（兼容旧的用户分组接口）
//tagid_list	用户被打上的标签ID列表

//{"subscribe":1,"openid":"oH4HtwGsY-0JSjhNhJLA7jYYOMsQ","nickname":"话题终结者","sex":1,"language":"zh_CN",
// "city":"海淀","province":"北京","country":"中国","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/PiajxSqBRaELA3P616lZg1b33Vzv8Hhc6GnPEOmUZG7Gl8sL4aNic6WhOmezogpWmoXad53BuoVXmPHTsnVLbbvA\/0",
// "subscribe_time":1487925963,"remark":"","groupid":0,"tagid_list":[]}

type User struct {
	Subscribe     int    `json:"subscribe"`
	OpenID        string `json:"openid"`
	NickName      string `json:"nickname"`
	Sex           int    `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	HeadImgUrl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
	Remark        string `json:"remark"`
	GroupID       int    `json:"groupid"`
	TagID_List    []int  `json:"tagid_list"`
}

type userTmp struct {
	*User
	*errors.Error
}

type userIDsTmp struct {
	*Users
	*errors.Error
}

type Users struct {
	Total int                 `json:"total"`
	Count int                 `json:"count"`
	Data  map[string][]string `json:"data"`
}
