package menu

import (
	"encoding/json"
	"testing"
)

func TestNewTopButton(t *testing.T) {
	buttons := NewTopButton("体检向导").AddSub(NewViewButton("预约体检", "www.baidu.com")).AddSub(NewViewButton("报告查询", "www.baidu.com"))
	b, err := json.Marshal(buttons)
	if err != nil {
		t.Fatal(err)
	}
	expect := `{"name":"体检向导","sub_button":[{"type":"view","name":"预约体检","url":"www.baidu.com"},{"type":"view","name":"报告查询","url":"www.baidu.com"}]}`
	if string(b) != expect {
		t.Fatal(string(b))
	}
}
