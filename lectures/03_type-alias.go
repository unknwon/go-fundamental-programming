package main

import (
	"fmt"
)

type (
	byte int8
	rune int32
	文本   string
)

func main() {
	var b 文本
	b = "中文啊亲"
	fmt.Println("文本b：" + b)
}
