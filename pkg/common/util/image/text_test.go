package image

import "testing"

func TestGenMyWords(t *testing.T) {
	if err := InitFont(); err != nil {
		t.Fatal(err)
	}

	if err := genMyWords("北京迪安开元", "大潘潘", "out.png"); err != nil {
		t.Fatal(err)
	}
}
