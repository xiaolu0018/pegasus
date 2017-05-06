package pinto

import (
	"fmt"

	"encoding/json"
	"github.com/1851616111/util/http"
)

const JSON_KEY_CHECKUP_CODES = "checkupcodes"

type PintoSDK struct {
}

func NewPintoSDK() *PintoSDK {
	p := PintoSDK{}
	return &p
}

func (p *PintoSDK) GetCheckupCodesBySaleCodes(codes []string, ip string) ([]string, error) {
	rsp, err := http.Send(&http.HttpSpec{
		URL:         ip + "/api/sales/checkups",
		Method:      "GET",
		ContentType: http.ContentType_JSON,
		BodyParams:  http.NewBody().Add("salecodes", &codes),
	})
	if err != nil {
		return nil, err
	}
	result := map[string][]string{}

	err = json.NewDecoder(rsp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if checkups, ok := result[JSON_KEY_CHECKUP_CODES]; ok {
		return checkups, nil
	} else {
		return nil, fmt.Errorf("rep checkupcodes not found")
	}
}
