package pinto

import (
	"github.com/golang/glog"
	"os"
	"sync"
	"time"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/cache"
	org "192.168.199.199/bjdaos/pegasus/pkg/appoint/organization"
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	pintoapi "192.168.199.199/bjdaos/pegasus/pkg/common/api/pinto"
	"192.168.199.199/bjdaos/pegasus/pkg/common/util/database"
)

const CACHE_ORG = "organizations"
const CACHE_CHECKUP = "checkups"

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

type Config struct {
	User     string
	Password string
	IP       string
	Port     string
	Database string
}

func (c *Config) Run() {
	if err := c.sync(); err != nil {
		glog.Errorf("Config.Run: sync pinto organizations err %v\n", err)
		os.Exit(1)
	}
	glog.Errorln("Config.Run: sync pinto organizations success")

	for {
		select {
		case <-time.NewTicker(SYNC_WIN_SIZE).C:
			if err := c.sync(); err != nil {
				glog.Errorf("Config.Run: sync pinto organizations err %v\n")
			}
			glog.Errorln("Config.Run: sync pinto organizations success")
		}
	}
}

func (c *Config) sync() error {
	database, err := database.Init(c.User, c.Password, c.IP, c.Port, c.Database)
	if err != nil {
		return err
	}
	defer database.Close()

	orgs, err := pintoapi.ListOrganizations(database)
	if err != nil {
		return err
	}

	glog.Infof("SyncOrganizations cached orgs %v\n", orgs)
	cache.Set("inner_system", CACHE_ORG, orgs)

	olds, err := org.ListDBOrgs()
	if err != nil {
		return err
	}

	adds, dels, changes := types.Diff(orgs, olds)

	if err := org.CreateOrgs(adds); err != nil {
		glog.Errorf("sync() create new orgs err %v\n", err)
	}

	if err := org.DeleteOrgs(dels); err != nil {
		glog.Errorf("sync () delete old orgs err %v\n", err)
	}

	if err := org.UpdateOrgs(changes); err != nil {
		glog.Errorf("sync () update change orgs err %v\n", err)
	}


	cks, err := pintoapi.ListCheckups(database)
	if err != nil {
		glog.Errorf("Sync checkups cached cks err %v\n", err)
		return err
	}

	glog.Infof("Sync checkups cached cks %v\n", cks)
	cache.Set("inner_system", CACHE_CHECKUP, cks)

	return nil
}

