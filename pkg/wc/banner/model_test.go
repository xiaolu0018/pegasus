package banner

import (
	"bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestBanner_Create(t *testing.T) {
	b := Banner{Pos: 2, RedirectUrl: "www.ccc.com", ImageUrl: "www/image"}
	err := b.CreateOrUpdate(db.Banner())
	fmt.Println("1create", err)
	time.Sleep(time.Second * 20)
	b = Banner{Pos: 2, RedirectUrl: "www.ccc3.com", ImageUrl: "www/image3"}
	err = b.CreateOrUpdate(db.Banner())
	fmt.Println("2create", err)
}

func TestFindBanners(t *testing.T) {
	bannes, err := FindBanners(db.Banner(), bson.M{"hide": false})

	fmt.Println("bannes,", bannes)
	fmt.Println("err", err)
}
