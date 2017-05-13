package main

import "fmt"

func main() {
	word := fmt.Sprintf("                  我是%s的%s,", "北京", "小潘")
	word2 := fmt.Sprintf("我是%s的%s,", "北京", "小潘")
	fmt.Println(len(word))
	fmt.Println(len(word2))
}
