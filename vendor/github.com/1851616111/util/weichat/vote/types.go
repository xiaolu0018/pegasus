package vote

import (
	"fmt"
	"github.com/1851616111/util/validator/mobile"
	"github.com/1851616111/util/validator/tel"
)

func ParamNotFoundError(param string) error {
	return fmt.Errorf("param %s not found", param)
}

func ParamInvalidError(param string) error {
	return fmt.Errorf("param %s invalid", param)
}

type Voter struct {
	ID          string `json:"voterid"`
	OpenID      string `json:"openid,omitempty"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Company     string `json:"company,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
	Declaration string `json:"declaration,omitempty"`
	VotedCount  int    `json:"voteCount"`
}

func (v Voter) ValidateRegister() error {
	if len(v.OpenID) == 0 {
		return ParamNotFoundError("openid")
	}

	if len(v.Name) == 0 {
		return ParamNotFoundError("name")
	}

	if len(v.Image) == 0 {
		return ParamNotFoundError("image")
	}

	if len(v.Mobile) == 0 {
		return ParamNotFoundError("mobile")
	}

	err1, err2 := mobile.Validate(v.Mobile), tel.Validate(v.Mobile)
	if err1 != nil && err2 != nil {
		return ParamInvalidError("mobile")
	}

	if len(v.Company) == 0 {
		return ParamNotFoundError("company")
	}

	if len(v.Declaration) == 0 {
		return ParamNotFoundError("decalration")
	}

	return nil
}

func (v *Voter) Complete() {
	v.VotedCount = 0
}

type VoterList struct {
	Index 		int `json:"index"`
	Size 		int `json:"size"`
	TotalPages 	int `json:"total_pages"`
	PageData 	[]Voter `json:"page_data"`
}