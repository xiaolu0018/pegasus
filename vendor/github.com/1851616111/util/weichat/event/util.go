package event

import (
	"github.com/1851616111/util/xml"
)

func NewArticleAction(title, desc, pic, url string) *Action {
	it := item{
		Title:       xml.CDATA(title),
		Description: xml.CDATA(desc),
		PicUrl:      xml.CDATA(pic),
		Url:         xml.CDATA(url),
	}

	return &Action{
		Common: Common{
			Type: xml.CDATA(E_News),
		},
		ArticleCount: 1,
		Items:        &[]item{it},
	}
}
