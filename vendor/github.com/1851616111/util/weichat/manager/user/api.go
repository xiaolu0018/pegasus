package user

import (
	"encoding/json"
	"errors"
	httputil "github.com/1851616111/util/http"
	errs "github.com/1851616111/util/weichat/errors"
	"net/http"
)

const URL_ListUserIDs = "https://api.weixin.qq.com/cgi-bin/user/get"
const URL_UserInfo = "https://api.weixin.qq.com/cgi-bin/user/info"

//details https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842&token=&lang=zh_CN
//TODO 目前一批获取10000个用户，超过10000个需要再处理，先这么写
func ListUserIDs(token string) ([]string, error) {
	if len(token) == 0 {
		return nil, errs.ErrInvalidToken
	}

	rsp, err := http.Get(URL_ListUserIDs + "?access_token=" + token)
	if err != nil {
		return nil, err
	}

	tmp := userIDsTmp{}
	if err := json.NewDecoder(rsp.Body).Decode(&tmp); err != nil {
		return nil, err
	}

	if tmp.Error != nil {
		return nil, errors.New(tmp.Error.Msg)
	} else {
		return tmp.Data["openid"], nil
	}
}

func GetUserDetails(token, openID string) (*User, error) {
	if len(token) == 0 || len(openID) == 0 {
		return nil, errs.ErrInvalidToken
	}

	rsp, err := httputil.Send(&httputil.HttpSpec{
		URL:    URL_UserInfo,
		Method: "GET",

		URLParams: httputil.NewParams().Add("access_token", token).Add("openid", openID).Add("lang", "zh_CN"),
	})
	if err != nil {
		return nil, err
	}

	tmp := userTmp{}
	if err := json.NewDecoder(rsp.Body).Decode(&tmp); err != nil {
		return nil, err
	}

	if tmp.Error != nil {
		return nil, errors.New(tmp.Error.Msg)
	} else {
		return tmp.User, nil
	}
}
