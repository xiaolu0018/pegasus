package cache

import (
	"sync"
	"time"
)

var CACHE_HOUSEKEEPING_PERIOD = time.Minute * 30

type TTLCache interface {
	Get(string) (interface{}, error)
	Set(string, interface{})
}

func NewTTLCache(sec int64) TTLCache {
	c := impl{
		ttl:           sec,
		objMapping:    map[string]interface{}{},
		objTTlMapping: map[string]int64{},
	}

	go func() {
		ticker := time.NewTicker(CACHE_HOUSEKEEPING_PERIOD)
		for {
			select {
			case <-ticker.C:
				c.DelExpired()
			}
		}
	}()

	return &c
}

type impl struct {
	l             sync.RWMutex
	ttl           int64
	objMapping    map[string]interface{}
	objTTlMapping map[string]int64
}

func (i *impl) Get(key string) (interface{}, error) {
	i.l.RLock()
	defer i.l.RUnlock()

	obj, ok := i.objMapping[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return obj, nil
}

func (i *impl) Set(key string, value interface{}) {
	i.l.Lock()
	defer i.l.Unlock()

	if _, ok := i.objMapping[key]; !ok {
		i.objTTlMapping[key] = time.Now().Unix()
	}

	i.objMapping[key] = value
}

func (i *impl) DelExpired() {
	nowSec := time.Now().Unix()
	for key, keyTime := range i.objTTlMapping {
		if nowSec-keyTime >= i.ttl {
			i.delObj(key)
		}
	}
}

func (i *impl) delObj(key string) {
	i.l.Lock()
	defer i.l.Unlock()
	delete(i.objMapping, key)
	delete(i.objTTlMapping, key)
}
