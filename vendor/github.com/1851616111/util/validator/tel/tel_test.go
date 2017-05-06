package tel

import "testing"

func TestValidate(t *testing.T) {
	validTels := []string{"010-87328719", "010-7328719"}
	invalidTels := []string{"", "138000000000", "1380000000a", "aaaaaaaaaaa"}

	for _, tel := range validTels {
		if err := Validate(tel); err != nil {
			t.Fatal(err)
		}
	}

	for _, tel := range invalidTels {
		if err := Validate(tel); err == nil {
			t.Fatal(err)
		}
	}
}
