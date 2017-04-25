package user

import (
	tt "github.com/1851616111/util/weichat/manager/test"
	"testing"
)

func TestListUserIDs(t *testing.T) {
	_, err := ListUserIDs(tt.Dev_Basic_Token)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserDetails(t *testing.T) {
	_, err := GetUserDetails(tt.Dev_Basic_Token, tt.Dev_User_OpenID)
	if err != nil {
		t.Fatal(err)
	}
}
