package image

import (
	"image/png"
	"fmt"
	"image"
	"os"
	"math"
	"testing"
	"image/draw"
)

func TestGroupPic(t *testing.T) {
	//err := GemPersonPic("lyf.png", "280_110.png")
	//if err != nil {
	//	t.Fatal(err)
	//}
}

func GemPersonPic(target, pic_whoami,  pic_person, pic_comeon, pic_erweima, pic_same_yang, pic_declaration, pic_top string  error {
	file, err := os.Create(target) //需生成的文件名
	if err != nil {
		return fmt.Errorf(err.Error() + "create err")
	}
	defer file.Close()
	//
	file1, err := os.Open(pic_comeon) //绿色ｃｏｍｅ　ｏｎ
	if err != nil {
		return fmt.Errorf(err.Error() + "下面底色出错")
	}
	defer file1.Close()
	img, _ := png.Decode(file1)
	size1 := img.Bounds().Size()
	file2, err := os.Open(pic_erweima) //二维码
	if err != nil {
		return fmt.Errorf(err.Error() + "二维码错误")
	}
	defer file2.Close()
	img2, err := png.Decode(file2)
	if err != nil {
		return fmt.Errorf(err.Error() + "二维码错误Decode")
	}
	size2 := img2.Bounds().Size()

	file3, err := os.Open(pic_same_yang) //SAME YOUNG
	if err != nil {
		return fmt.Errorf(err.Error() + "中间文字错误")
	}
	defer file3.Close()
	png3, err := png.Decode(file3)
	if err != nil {
		return fmt.Errorf(err.Error() + "中间文字错误 decode")
	}
	size3 := png3.Bounds().Size()

	file4, err := os.Open(pic_person) //个人照片
	if err != nil {
		return fmt.Errorf(err.Error() + "个人照片错误")
	}
	defer file4.Close()
	png4, err := png.Decode(file4)
	if err != nil {
		return fmt.Errorf(err.Error() + "个人照片错误 decode")
	}
	size4 := png4.Bounds().Size()
	defer file4.Close()
	png_ := image.NewRGBA(image.Rect(0, 0, 730, size4.Y+size1.Y+size3.Y-20))
	//

	file5, err := os.Open(pic_whoami) //文字照片
	png5, err := png.Decode(file5)

	defer file5.Close()
	file6, err := os.Open(pic_declaration)
	if err != nil {
		return fmt.Errorf(err.Error() + "文字照片错误")
	}
	png6, err := png.Decode(file6)
	if err != nil {
		return fmt.Errorf(err.Error() + "文字照片错误 Decode")
	}
	defer file6.Close()

	file7, err := os.Open(pic_top)
	if err != nil {
		return fmt.Errorf(err.Error() + "顶层底版错误")
	}
	png7, err := png.Decode(file7)
	if err != nil {
		return fmt.Errorf(err.Error() + "顶层底版错误decide")
	}

	draw.Draw(png_, png_.Bounds(), png7, png4.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over)
	draw.Draw(png_, png_.Bounds(), png4, png4.Bounds().Min.Sub(image.Pt(int(math.Abs(float64(730)-float64(size4.Y))/2), 0)), draw.Over)
	draw.Draw(png_, png_.Bounds(), png6, png6.Bounds().Min.Sub(image.Pt(620, 40)), draw.Over)
	draw.Draw(png_, png_.Bounds(), png3, png3.Bounds().Min.Sub(image.Pt(130, 0+size4.Y)), draw.Over)                  //首先将一个图片信息存入jpg
	draw.Draw(png_, png_.Bounds(), img, img.Bounds().Min.Sub(image.Pt(0, size3.Y-20+size4.Y)), draw.Over)             //首先将一个图片信息存入jpg
	draw.Draw(png_, png_.Bounds(), img2, img2.Bounds().Min.Sub(image.Pt(245, 70+size3.Y+size4.Y)), draw.Over)         //将另外一张图片信息存入jpg
	draw.Draw(png_, png_.Bounds(), png5, png5.Bounds().Min.Sub(image.Pt(245, 90+size3.Y+size4.Y+size2.Y)), draw.Over) //将另外一张图片信息存入jpg

	return png.Encode(file, png_)
}
