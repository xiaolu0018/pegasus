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

	person := genVoterImagePath + "/" + "888.png"

	var declarationFile string = fmt.Sprintf(`%s/img/10%d.png`, dist, 1)
	//path, name, company, pic_person, pic_declaration  strin
	path, err := GenPersonPngPic(dist + "voterimages/gen", "北京", "大大泡泡堂", person, declarationFile)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path)
}