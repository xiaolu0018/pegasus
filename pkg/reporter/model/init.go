package model

import (
	"database/sql"
	"github.com/golang/glog"
)

var DB *sql.DB

func Init(db *sql.DB) error {
	DB = db

	var err error
	if h2tMappings, err = getH2TMappings(); err != nil {
		return err
	}

	getT2HMappings()

	glog.Infof("[health] health to template mappings ok, total %d\n", len(h2tMappings))
	glog.Infof("[health] template to health mappings ok, total %d\n", len(t2hMappings))

	//glog.Infof("[health] health to template mappings ok, %v\n", h2tMappings)
	//glog.Infof("[health] template to health mappings ok, %v\n", t2hMappings)

	return nil
}
