package event

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestXML(t *testing.T) {
	r := Action{}

	s, err := xml.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(s))
}
