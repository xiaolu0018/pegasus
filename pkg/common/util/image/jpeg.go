package image

import (
	"image/png"
	"fmt"
	"image"
	"os"
	"math"
	"image/draw"
	"time"
	"github.com/1851616111/util/rand"
	"image/jpeg"
)

func GenPersonJpgPic(path, name, company, pic_person, pic_comeon, pic_erweima, pic_same_yang, pic_declaration, pic_top string) ( string,  error) {
	myWords := fmt.Sprintf("%s/%d_%s.png", path, time.Now().UnixNano(), rand.String(20))
	if err := genMyWords(company, name, myWords); err != nil {
		return "", err
	}

	target := fmt.Sprintf("%d_%s.png", time.Now().UnixNano(), rand.String(20))
	file, err := os.Create(target) //需生成的文件名
	if err != nil {
		return "",fmt.Errorf(err.Error() + "create err")
	}
	defer file.Close()
	//
	file1, err := os.Open(pic_comeon) //绿色ｃｏｍｅ　ｏｎ
	if err != nil {
		return "",fmt.Errorf(err.Error() + "下面底色出错")
	}
	defer file1.Close()
	img, _ := png.Decode(file1)
	size1 := img.Bounds().Size()
	file2, err := os.Open(pic_erweima) //二维码
	if err != nil {
		return "",fmt.Errorf(err.Error() + "二维码错误")
	}
	defer file2.Close()
	img2, err := png.Decode(file2)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "二维码错误Decode")
	}
	size2 := img2.Bounds().Size()

	file3, err := os.Open(pic_same_yang) //SAME YOUNG
	if err != nil {
		return "",fmt.Errorf(err.Error() + "中间文字错误")
	}
	defer file3.Close()
	png3, err := png.Decode(file3)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "中间文字错误 decode")
	}
	size3 := png3.Bounds().Size()

	file4, err := os.Open(pic_person) //个人照片
	if err != nil {
		return "",fmt.Errorf(err.Error() + "个人照片错误")
	}
	defer file4.Close()

	img4, err := jpeg.Decode(file4)
	if err != nil {
		return "", err
	}
	size4 := img4.Bounds().Size()
	dst := image.NewRGBA(image.Rect(0, 0, 730, size4.Y+size1.Y+size3.Y-20))

	file5, err := os.Open(myWords) //文字照片
	if err != nil {
		return "",fmt.Errorf(err.Error() + "我的文字照片")
	}
	png5, err := png.Decode(file5)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "解析我的文字照片")
	}
	defer file5.Close()


	file6, err := os.Open(pic_declaration)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "文字照片错误")
	}
	png6, err := png.Decode(file6)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "文字照片错误 Decode")
	}
	defer file6.Close()

	file7, err := os.Open(pic_top)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "顶层底版错误")
	}
	png7, err := png.Decode(file7)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "顶层底版错误decide")
	}

	draw.Draw(dst, dst.Bounds(), png7, img4.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over)
	draw.Draw(dst, dst.Bounds(), img4, img4.Bounds().Min.Sub(image.Pt(int(math.Abs(float64(730)-float64(size4.Y))/2), 0)), draw.Over)
	draw.Draw(dst, dst.Bounds(), png6, png6.Bounds().Min.Sub(image.Pt(520, 150)), draw.Over)
	draw.Draw(dst, dst.Bounds(), png3, png3.Bounds().Min.Sub(image.Pt(130, 0+size4.Y)), draw.Over)                  //首先将一个图片信息存入jpg
	draw.Draw(dst, dst.Bounds(), img, img.Bounds().Min.Sub(image.Pt(0, size3.Y-20+size4.Y)), draw.Over)             //首先将一个图片信息存入jpg
	draw.Draw(dst, dst.Bounds(), img2, img2.Bounds().Min.Sub(image.Pt(245, 70+size3.Y+size4.Y)), draw.Over)         //将另外一张图片信息存入jpg
	draw.Draw(dst, dst.Bounds(), png5, png5.Bounds().Min.Sub(image.Pt(245, 90+size3.Y+size4.Y+size2.Y)), draw.Over) //将另外一张图片信息存入jpg

	if err := png.Encode(file, dst); err != nil {
		return "", err
	}

	return target, nil
}