package menu

import (
	"encoding/json"
	"github.com/1851616111/util/http"
	"github.com/1851616111/util/weichat/errors"
)

const NewMenuURL = "https://api.weixin.qq.com/cgi-bin/menu/create"

func CreateMenu(bt *Button, access_token string) error {
	req := &http.HttpSpec{
		URL:         NewMenuURL,
		Method:      "POST",
		ContentType: http.ContentType_JSON,
		URLParams:   http.NewParams().Add("access_token", access_token),
		BodyParams:  http.NewBody().Add("button", []*Button{bt}),
	}

	rsp, err := http.Send(req)
	if err != nil {
		return err
	}

	tmp := errors.Error{}
	if err := json.NewDecoder(rsp.Body).Decode(&tmp); err != nil {
		return err
	}

	if tmp.Code == errors.CODE_SUCCESS {
		return nil
	} else {
		return tmp.Error()
	}
}
