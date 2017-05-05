package start

import (
	"fmt"
	"github.com/julienschmidt/httprouter"

	"bjdaos/pegasus/pkg/common/util/database"
	"bjdaos/pegasus/pkg/wc/handler"
	"bjdaos/pegasus/pkg/wc/manager"
	"os"
	"path/filepath"
)

func NewActivityConfig() *ActivityConfig {
	return &ActivityConfig{
		DBConfig: database.DBConfig{},
	}
}

type ActivityConfig struct {
	database.DBConfig

	ListenAddress string

	Schema string
	Domain string

	DistBasePath  string
	LocalDistPath string

	AppID     string
	AppSecret string
}

func (o *ActivityConfig) Start(router *httprouter.Router) error {

	if err := o.initApiToken(); err != nil {
		return err
	}

	if err := o.StartActivity(router); err != nil {
		return err
	}

	if err := manager.CreateMenu(o.Schema, o.Domain); err != nil {
		return err

	}

	if err := manager.WatchEvent(); err != nil {
		return err
	}

	mrManager, err := handler.NewMenuRedirectManager(o.GetAbsPath(), "openid")
	if err != nil {
		return err
	}
	mrManager.Redirect("/api/activity/index", "activity/index_activity.html", handler.CompleteOpenidInfo)

	if err := mrManager.AddRedirectToRouter(router); err != nil {
		return err
	}

	return nil
}

func (o *ActivityConfig) GetAbsPath() string {
	return fmt.Sprintf("%s://%s/%s", o.Schema, o.Domain, o.DistBasePath)
}

func (o *ActivityConfig) GetVoteCachedImagePath() (string, error) {
	absPath, err := filepath.Abs(o.LocalDistPath)
	if err != nil {
		return "", err
	}

	target := fmt.Sprintf("%s/voterimages", absPath)
	_, err = os.Stat(target)
	if err != nil {
		//if os.IsNotExist(err) {
		//	if err := os.MkdirAll(target, os.ModeDir); err != nil {
		//		return "", err
		//	}
		//	return target, nil
		//}
		return "", err
	}

	return absPath, nil
}
