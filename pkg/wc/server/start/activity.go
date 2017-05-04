package start

import (
	"fmt"
	"github.com/1851616111/util/weichat/vote"
	"github.com/julienschmidt/httprouter"
)

func (o *ActivityConfig) StartActivity(r *httprouter.Router) error {
	dbI, err := vote.NewDBInterface(o.Ip, o.Port, o.User, o.Passｗord, o.Database)
	if err != nil {
		return err
	}

	if err := dbI.Init(); err != nil {
		return err
	}

	vote.SetDB(dbI)
	vote.AddRouter(r, o.LocalDistPath)
	vote.APPID = o.AppID
	vote.URL_REGISTER_HTML = fmt.Sprintf("%s://%s/dist/activity/regist.html", o.Schema, o.Domain)
	vote.CH_CACHE_IMAGES = make(chan string, 20)
	cachePath, err := o.GetVoteCachedImagePath()
	if err != nil {
		return err
	}

	go vote.StartImageCachedController(cachePath)

	return nil
}
