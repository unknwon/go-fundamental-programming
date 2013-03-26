package main

import (
	"fmt"
)

var (
	// 使用常规方式
	aaa = "hello"
	// 使用并行方式以及类型推断
	sss, bbb = 1, 2
	// ccc := 3	// 不可以省略 var
)

func main() {
	var a int = 65
	b := string(a)
	fmt.Println(b)
}
