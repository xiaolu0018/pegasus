package image

import "testing"

func TestGenMyWords(t *testing.T) {
	if err := InitFont(); err != nil {
		t.Fatal(err)
	}

	if err := genMyWords("说的的的的的", "大的的水", "out.png"); err != nil {
		t.Fatal(err)
	}
}
