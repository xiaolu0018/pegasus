package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"os"
)

func GenImageFromBase64(src []byte, dst string) error {
	var data []byte = make([]byte, len(src)*4/3+100)
	if _, err := base64.StdEncoding.Decode(data, src); err != nil {
		return err
	}

	m, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	rgbImg := m.(*image.YCbCr)

	f, _ := os.Create(dst)
	defer f.Close()

	return jpeg.Encode(f, rgbImg, nil)
}
