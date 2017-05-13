package image

import (
	"testing"
	"fmt"
)

func TestGemPersonPic(t *testing.T) {
	InitFont()
	dist := "/home/michael/IdeaProjects/src/bjdaos/pegasus/dist/activity/"

	var voterImagesPath = fmt.Sprintf("%s/voterimages", dist)
	var genVoterImagePath = fmt.Sprintf("%s/gen", voterImagesPath)
	var comeOnFile = fmt.Sprintf("%s/%s", dist, "img/11.png")
	var _2weimaFile = fmt.Sprintf("%s/%s", dist, "img/22_png.png")
	var yongFile = fmt.Sprintf("%s/%s", dist, "img/333.png")
	var topFile = fmt.Sprintf("%s/%s", dist, "img/topblock.png")

	person := genVoterImagePath + "/" + "888.png"

	var declarationFile string = fmt.Sprintf(`%s/img/10%d.png`, dist, 1)
	path, err := GemPersonPngPic(dist + "voterimages/gen", "北京", "大大泡泡堂", person , comeOnFile, _2weimaFile,
	yongFile, declarationFile, topFile)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path)
}