package appointment

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/cache"
	"github.com/golang/glog"
	"os"
)

const (
	CACHE_TP      = "appoint_cache"
	CACHE_TTL_SEC = 1800
)

func init() {
	if err := cache.Register(CACHE_TP, CACHE_TTL_SEC); err != nil {
		glog.Errorf("organization register cache error %v\n", err)
		os.Exit(1)
	}
}
