package cache

import (
	"testing"
	"time"
)

func TestNewTTLCache(t *testing.T) {
	c := NewTTLCache(1)
	CACHE_HOUSEKEEPING_PERIOD = time.Second * 1
	c.Set("key007", "value007")
	if _, err := getValue(c, "key008"); err != ErrKeyNotFound {
		t.Fatal()
	}

	value, err := getValue(c, "key007")
	if err != nil {
		t.Fatal()
	}
	if value != "value007" {
		t.Fatal()
	}

	time.Sleep(time.Second * 3)

	_, err = getValue(c, "key007")
	if err != ErrKeyNotFound {
		t.Fatal(err)
	}

}

func getValue(c TTLCache, key string) (string, error) {
	v, err := c.Get(key)
	if err != nil {
		return "", err
	}
	return v.(string), nil
}
