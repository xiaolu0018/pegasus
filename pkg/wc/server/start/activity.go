package start

import (
	"github.com/1851616111/util/weichat/vote"
	"github.com/julienschmidt/httprouter"

	apitoken "github.com/1851616111/util/weichat/util/api-token"
)

func (o *ActivityConfig) StartActivity(r *httprouter.Router) error {
	dbI, err := vote.NewDBInterface(o.Ip, o.Port, o.User, o.Password, o.Database)
	if err != nil {
		return err
	}

	if err := dbI.Init(apitoken.TokenCtrl.GetToken()); err != nil {
		return err
	}

	vote.SetDB(dbI)
	vote.AddRouter(r, o.LocalDistPath)
	vote.APPID = o.AppID
	vote.CH_CACHE_IMAGES = make(chan string, 20)
	cachePath, err := o.GetVoteCachedImagePath()
	if err != nil {
		return err
	}

	go vote.StartImageCachedController(cachePath)

	return nil
}
