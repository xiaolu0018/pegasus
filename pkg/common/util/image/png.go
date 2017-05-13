package image

import (
	"image/png"
	"fmt"
	"image"
	"os"
	"image/draw"
	"time"
	"github.com/1851616111/util/rand"
	"github.com/hunterhug/go_image"
)

func GenPersonPngPic(path, name, company, pic_person, pic_comeon, pic_erweima, pic_same_yang, pic_declaration, pic_top string) ( string,  error) {
	file7, err := os.Open(pic_top)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "顶层底版错误")
	}
	defer file7.Close()

	png7, err := png.Decode(file7)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "顶层底版错误decide")
	}

	height7 := png7.Bounds().Size().Y
	totalWeight :=  png7.Bounds().Size().X

	myWords := fmt.Sprintf("%s/%d_%s.png", path, time.Now().UnixNano(), rand.String(20))
	if err := genMyWords(company, name, myWords); err != nil {
		return "", err
	}

	targetName  := fmt.Sprintf("%d_%s.png", time.Now().UnixNano(), rand.String(20))
	target := fmt.Sprintf("%s/%s", path, targetName)
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

	img2_left := (totalWeight - size2.X)/2

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

	img4, err := png.Decode(file4)
	if err != nil {
		return "", err
	}
	size4 := img4.Bounds().Size()

	img4_new_weight := int(float64(size4.X) / float64(size4.Y) * float64(height7))
	newName := fmt.Sprintf("%s/tmp_%d_%s.jpg", path, time.Now().UnixNano(), rand.String(20))
	img4_left := (totalWeight - img4_new_weight)/2

	if err := go_image.ThumbnailF2F(pic_person, newName, img4_new_weight, height7); err != nil {
		return "",fmt.Errorf(err.Error() + "我的文字照片生成缩略图")
	}

	file4, err = os.Open(newName) //个人照片
	if err != nil {
		return "",fmt.Errorf(err.Error() + "个人照片错略图错误")
	}
	defer file4.Close()

	img4, err = png.Decode(file4)
	if err != nil {
		return "", err
	}
	size4 = img4.Bounds().Size()

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

	draw.Draw(dst, dst.Bounds(), png7, img4.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over)
	draw.Draw(dst, dst.Bounds(), img4, img4.Bounds().Min.Sub(image.Pt(img4_left, 0)), draw.Over)
	draw.Draw(dst, dst.Bounds(), png6, png6.Bounds().Min.Sub(image.Pt(520, 150)), draw.Over)
	draw.Draw(dst, dst.Bounds(), png3, png3.Bounds().Min.Sub(image.Pt(130, 0+size4.Y)), draw.Over)                  //首先将一个图片信息存入jpg
	draw.Draw(dst, dst.Bounds(), img, img.Bounds().Min.Sub(image.Pt(0, size3.Y-20+size4.Y)), draw.Over)             //首先将一个图片信息存入jpg
	draw.Draw(dst, dst.Bounds(), img2, img2.Bounds().Min.Sub(image.Pt(img2_left, 70+size3.Y+size4.Y)), draw.Over)         //将另外一张图片信息存入jpg
	draw.Draw(dst, dst.Bounds(), png5, png5.Bounds().Min.Sub(image.Pt(245, 90+size3.Y+size4.Y+size2.Y)), draw.Over) //将另外一张图片信息存入jpg

	if err := png.Encode(file, dst); err != nil {
		return "", err
	}

	return targetName, nil
}
