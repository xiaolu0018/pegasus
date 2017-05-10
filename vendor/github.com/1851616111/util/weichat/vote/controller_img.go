package vote

import (
	"fmt"
	"github.com/1851616111/util/http"
	"github.com/1851616111/util/http/file"
	apitoken "github.com/1851616111/util/weichat/util/api-token"
	"github.com/golang/glog"
)

var CH_CACHE_IMAGES chan string

func StartImageCachedController(cachePath string) {
	for {
		select {
		case imageID := <-CH_CACHE_IMAGES:
			token := apitoken.TokenCtrl.GetToken()
			if len(token) == 0 {
				glog.Errorf("weichat vote image controller get api token null")
				continue
			}

			spec := newImageReqSpec(token, imageID)
			targetFile := fmt.Sprintf("%s/%s.jpg", cachePath, imageID)
			if err := file.GetHttpImage(targetFile, spec); err != nil {
				glog.Errorf("weichat vote image controller get image err %v\n", err)
				continue
			}

			if err := DBI.updateVoterImageStatus(imageID); err != nil {
				glog.Errorf("weichat vote image controller update image status err %v\n", err)
				continue
			}
		}
	}
}

func newImageReqSpec(token, mediaID string) *http.HttpSpec {
	return &http.HttpSpec{
		URL:    "https://api.weixin.qq.com/cgi-bin/media/get",
		Method: "GET",
		URLParams: http.NewParams().
			Add("access_token", token).
			Add("media_id", mediaID),
	}
}
