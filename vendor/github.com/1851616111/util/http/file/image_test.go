package file

import (
	"github.com/1851616111/util/http"
	"os"
	"testing"
)

func TestGetHttpImage(t *testing.T) {
	spec := http.HttpSpec{
		URL:    "https://api.weixin.qq.com/cgi-bin/media/get",
		Method: "GET",
		URLParams: http.NewParams().
			Add("access_token", "Zsbg-4gK8OeB1ZLowEr-pBZSDptM1epuFaYIOdw1yE3fGPM8MDvzeU0qq7woWOdC2s_zet8T_sEDL_3YPfQ54cGE4GNIobT6BJmTEiXPN-UQMBjAAAMIX").
			Add("media_id", "AwP_Ojj8sjxpcUOrUpu1e5nSz9j9AQZTBXMjAW2JellI1jYdubs_gufd_C5wLawE"),
	}

	targetFile := "abc.jpg"
	if err := GetHttpImage(targetFile, &spec); err != nil {
		t.Fatal(err)
	}

	defer func() {
		os.Remove(targetFile)
	}()

}
