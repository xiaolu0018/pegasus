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

const topHeight = 383
const weight = 757
const baseFile = "./base.png"

func GenPersonPic(path, name, company, pic_person, pic_declaration string) ( string,  error) {
	targetName  := fmt.Sprintf("%d_%s.png", time.Now().UnixNano(), rand.String(20))
	target := fmt.Sprintf("%s/%s", path, targetName)

	base, err := os.Open(baseFile)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "基础图片")
	}
	defer base.Close()

	basePng, err := png.Decode(base)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "基础图片decode")
	}

	file, err := os.Create(target) //需生成的文件名
	if err != nil {
		return "",fmt.Errorf(err.Error() + "create err")
	}
	defer file.Close()

	myWords := fmt.Sprintf("%s/%d_%s.png", path, time.Now().UnixNano(), rand.String(20))
	if err := genMyWords(company, name, myWords); err != nil {
		return "", err
	}

	myWordsFile, err := os.Open(myWords) //文字照片
	if err != nil {
		return "",fmt.Errorf(err.Error() + "我的文字照片")
	}
	defer myWordsFile.Close()

	myWord_png, err := png.Decode(myWordsFile)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "解析我的文字照片")
	}

	personFile, err := os.Open(pic_person) //个人照片
	if err != nil {
		return "",fmt.Errorf(err.Error() + "个人照片错误")
	}
	defer personFile.Close()

	person_png, err := png.Decode(personFile)
	if err != nil {
		return "", err
	}

	declare, err := os.Open(pic_declaration)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "我的声明")
	}
	declare_png, err := png.Decode(declare)
	if err != nil {
		return "",fmt.Errorf(err.Error() + "我的声明 Decode")
	}
	defer declare.Close()

	personWidth := person_png.Bounds().Size()

	person_new_weight := int(float64(personWidth.X) / float64(personWidth.Y) * float64(topHeight))
	newName := fmt.Sprintf("%s/tmp_%d_%s.jpg", path, time.Now().UnixNano(), rand.String(20))
	person_left := (weight - person_new_weight)/2

	if err := go_image.ThumbnailF2F(pic_person, newName, person_new_weight, topHeight); err != nil {
		return "",fmt.Errorf(err.Error() + "我的文字照片生成缩略图")
	}

	newPersonFile, err := os.Open(newName) //个人照片
	if err != nil {
		return "",fmt.Errorf(err.Error() + "个人照片错略图错误")
	}
	defer newPersonFile.Close()

	newPersonPng, err := png.Decode(newPersonFile)
	if err != nil {
		return "", err
	}

	dst := image.NewRGBA(image.Rect(0, 0, weight, 1220))

	draw.Draw(dst, dst.Bounds(), basePng, basePng.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over)
	draw.Draw(dst, dst.Bounds(), newPersonPng, newPersonPng.Bounds().Min.Sub(image.Pt(person_left, 0)), draw.Over)
	draw.Draw(dst, dst.Bounds(), declare_png, declare_png.Bounds().Min.Sub(image.Pt(520, 150)), draw.Over)
	draw.Draw(dst, dst.Bounds(), myWord_png, myWord_png.Bounds().Min.Sub(image.Pt((weight - 380) / 2, 835)), draw.Over)

	if err := png.Encode(file, dst); err != nil {
		return "", err
	}

	return targetName, nil

}

