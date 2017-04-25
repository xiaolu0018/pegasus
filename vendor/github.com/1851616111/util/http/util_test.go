package http

import "testing"

func Test_Send(t *testing.T) {
	_, err := Send(&HttpSpec{
		URL: "https://www.baidu.com",
	})
	if err != nil {
		t.Fatal(err)
	}
}
