package xml

import (
	"encoding/xml"
	"testing"
)

type stu struct {
	Name CDATA
}

func TestCDATA_MarshalXML(t *testing.T) {
	stu := stu{"michael"}
	stuStr := `<stu><Name><![CDATA[michael]]></Name></stu>`
	b, err := xml.Marshal(stu)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != stuStr {
		t.Fatal(string(b))
	}

	if err := xml.Unmarshal([]byte(stuStr), &stu); err != nil {
		t.Fatal(err)
	}

	if stu.Name != "michael" {
		t.Fatal()
	}

}
