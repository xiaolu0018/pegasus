package organization

import (
	"github.com/golang/glog"
	"os"
	"sync"
	"time"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/cache"
	org "192.168.199.199/bjdaos/pegasus/pkg/appoint/organization"
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	pintoapi "192.168.199.199/bjdaos/pegasus/pkg/pinto/api"
	"192.168.199.199/bjdaos/pegasus/pkg/common/util/database"
)

const CACHE_KEY = "organizations"
const CACHE_TTL_SEC = 1800
const SYNC_WIN_SIZE = time.Second * 200

var once sync.Once

func init() {
	once.Do(func() {
		if err := cache.Register("inner_system", CACHE_TTL_SEC); err != nil {
			glog.Errorf("organization register cache error %v\n", err)
			os.Exit(1)
		}
	})
}

type SyncController struct {
	User     string
	Password string
	IP       string
	Port     string
	Database string
}

func (c *SyncController) Run() {
	if err := c.syncOrganizations(); err != nil {
		glog.Errorf("SyncController.Run: sync pinto organizations err %v\n", err)
		os.Exit(1)
	}
	glog.Errorln("SyncController.Run: sync pinto organizations success")

	for {
		select {
		case <-time.NewTicker(SYNC_WIN_SIZE).C:
			if err := c.syncOrganizations(); err != nil {
				glog.Errorf("SyncController.Run: sync pinto organizations err %v\n")
			}
			glog.Errorln("SyncController.Run: sync pinto organizations success")
		}
	}
}

func (c *SyncController) syncOrganizations() error {
	database, err := database.Init(c.User, c.Password, c.IP, c.Port, c.Database)
	if err != nil {
		return err
	}
	defer database.Close()

	pintoapi.InitPintoDB(database)
	news, err := pintoapi.ListAllOrgs()
	if err != nil {
		return err
	}

	glog.Infof("SyncOrganizations cached orgs %v\n", news)
	cache.Set("inner_system", CACHE_KEY, news)

	olds, err := org.ListDBOrgs()
	if err != nil {
		return err
	}

	adds, dels, changes := types.Diff(news, olds)

	if err := org.CreateOrgs(adds); err != nil {
		glog.Errorf("syncOrganizations() create new orgs err %v\n", err)
	}

	if err := org.DeleteOrgs(dels); err != nil {
		glog.Errorf("syncOrganizations() delete old orgs err %v\n", err)
	}

	if err := org.UpdateOrgs(changes); err != nil {
		glog.Errorf("syncOrganizations() update change orgs err %v\n", err)
	}
	return nil
}
