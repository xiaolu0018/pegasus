package cache

import (
	"errors"
	"sync"
)

var ErrKeyNotFound error = errors.New("key not found")
var ErrKeyDuplicated error = errors.New("key duplicated")

var c *Cache = NewCache()

type Cache struct {
	sync.Once
	cacheMapping map[string]TTLCache
}

func NewCache() *Cache {
	c := Cache{}

	var cacheMapping map[string]TTLCache = map[string]TTLCache{}
	c.Do(func() {
		c.cacheMapping = cacheMapping
	})

	return &c
}

func Register(tp string, ttlSec int64) error {
	if _, ok := c.cacheMapping[tp]; ok {
		return ErrKeyDuplicated
	}

	cache := NewTTLCache(ttlSec)
	c.cacheMapping[tp] = cache

	return nil
}

func Get(tp, key string) (interface{}, error) {
	return c.cacheMapping[tp].Get(key)
}

func Set(tp, key string, obj interface{}) {
	c.cacheMapping[tp].Set(key, obj)
}
