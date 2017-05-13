package image

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"fmt"
	"path/filepath"
)

var dpi float64 = 100
var fontfile string
var size float64 = 10
var spacing float64 = 1.5

var f *truetype.Font
var text = []string{
	//"                  我是北京的小潘,         ",
	"     我正在参加晒合影赢万元大奖活动,  ",
	"                   请为我投一票,          ",
	"希望有机会为父母赢取15800元健康大奖,",
	"                 让服务老得慢一些        ",
}

func InitFont() error {
	var  err  error
	fontfile, err = filepath.Abs( "./" + "msyh.ttf")
	if err != nil {
		return err
	}

	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		return err
	}

	f, err = freetype.ParseFont(fontBytes)
	return err
}

func genMyWords(company, name, target string) error {
	ruler := color.RGBA{103, 185, 140, 0xff}
	fg, bg := image.White, image.NewUniform(ruler)

	rgba := image.NewRGBA(image.Rect(0, 0, 280, 110))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(size)>>6))

	word := fmt.Sprintf("我是%s的%s,", company, name)
	leftWitheSpace := 24 - (len(word)/3 + len(word)%3)
	var withSpace string
	if  leftWitheSpace > 0 {
		for i := 0 ; i < leftWitheSpace; i ++ {
			withSpace += ` `
		}

		word = withSpace  + word
	}

	if _, err := c.DrawString(word, pt); err != nil {
		return err
	}
	pt.Y += c.PointToFixed(size * spacing)

	for _, s := range text {
		if _, err := c.DrawString(s, pt); err != nil {
			return err
		}
		pt.Y += c.PointToFixed(size * spacing)
	}

	outFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer outFile.Close()

	b := bufio.NewWriter(outFile)

	if err = png.Encode(b, rgba); err != nil {
		return err
	}

	if err = b.Flush(); err != nil {
		return err
	}

	return nil
}
