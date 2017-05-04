package file

import (
	"github.com/1851616111/util/http"
	"image/jpeg"
	"os"
)

func GetHttpImage(targetFile string, spec *http.HttpSpec) error {
	rsp, err := http.Send(spec)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(rsp.Body)
	if err != nil {
		return err
	}

	target, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer target.Close()

	return jpeg.Encode(target, img, &jpeg.Options{jpeg.DefaultQuality})
}
