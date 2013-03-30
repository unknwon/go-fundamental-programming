package main

import (
	"fmt"
)

func main() {
	var a int = 1
	var p *int = &a
	fmt.Println(*(p++))
}
