package cache

import "testing"

func TestNewCache(t *testing.T) {
	Register("tp_branch", 1)
	Set("tp_branch", "key_001", "value_001")
	v, err := Get("tp_branch", "key_001")
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := v.(string); !ok {
		t.Fatal()
	}

}
