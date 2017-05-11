package image

import (
	"io/ioutil"
	"testing"
)

func TestGenImageFromBase64(t *testing.T) {
	data, err := ioutil.ReadFile("og0wQszrIJjosjLz7aRFn3FgnQo0")
	if err != nil {
		t.Fatal(err)
	}

	GenImageFromBase64(data, "123.jpg")
}
