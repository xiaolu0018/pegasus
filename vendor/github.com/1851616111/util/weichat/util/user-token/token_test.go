package user_token

import "testing"

func TestConfig_Exchange(t *testing.T) {
	cfg := NewTokenConfig("wxd09c7682905819e6", "b9938ddfec045280eba89fab597a0c41")

	_, err := cfg.Exchange("011UR7sc0ruDiv19Tguc0UE3sc0UR7sF")
	if err != nil {
		t.Fatal(err)
	}

}
