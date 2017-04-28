package event

import (
	"encoding/xml"
	xmlutil "github.com/1851616111/util/xml"
)

var E_Subscribe string = "subscribe"
var E_UnSubscribe string = "unsubscribe"
var E_News string = "news"

//reference https://mp.weixin.qq.com/wiki/9/2c15b20a16019ae613d413e30cac8ea1.html
//reference https://mp.weixin.qq.com/wiki/9/2c15b20a16019ae613d413e30cac8ea1.html#.E5.9B.9E.E5.A4.8D.E5.9B.BE.E6.96.87.E6.B6.88.E6.81.AF
type Action struct {
	XMLName xml.Name `xml:"xml"`
	Common
	ArticleCount int     `xml:"ArticleCount,omitempty"`
	articles     string  `xml:"Articles,omitempty"`
	Items        *[]item `xml:"Articles>item,omitempty"`
}

type Common struct {
	To         xmlutil.CDATA `xml:"ToUserName"`
	From       xmlutil.CDATA `xml:"FromUserName"`
	CreateTime int64         `xml:"CreateTime"`
	Type       xmlutil.CDATA `xml:"MsgType"`
}

type Event struct {
	XMLName xml.Name `xml:"xml"`
	Common
	E    xmlutil.CDATA `xml:"Event"`
	EKey xmlutil.CDATA `xml:"EventKey"`
}

type item struct {
	Title       xmlutil.CDATA `xml:"Title"`
	Description xmlutil.CDATA `xml:"Description"`
	PicUrl      xmlutil.CDATA `xml:"PicUrl"`
	Url         xmlutil.CDATA `xml:"Url"`
}
