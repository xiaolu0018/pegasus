package common

import (
	"errors"
	"strings"
)

var ErrProvinceNotFound error = errors.New("address province not found")
var ErrCityNotFound error = errors.New("address city not found")
var ErrDistrictNotFound error = errors.New("address district not found")
var ErrDetailNotFound error = errors.New("address detail not found")

type Address struct {
	Country  string `json:"country,omitempty"`
	Province string `bson:"province" json:"province,omitempty"`
	City     string `bson:"city" json:"city,omitempty"`
	District string `bson:"district" json:"district,omitempty"`
	Details  string `bson:"details" json:"details,omitempty"`
}

func (a Address) Validate() error {
	if strings.TrimSpace(a.Province) == "" {
		return ErrProvinceNotFound
	}
	if strings.TrimSpace(a.City) == "" {
		return ErrCityNotFound
	}
	if strings.TrimSpace(a.District) == "" {
		return ErrDistrictNotFound
	}
	if strings.TrimSpace(a.Details) == "" {
		return ErrDetailNotFound
	}

	return nil
}
