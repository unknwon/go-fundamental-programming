package main

import (
	"fmt"
)

func main() {
	// 未排序数组
	sort := [...]int{1, 7, 4, 2, 5}
	fmt.Println(sort)

	// 冒泡排序，由大到小
	num := len(sort)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			// 比较大小
			if sort[i] < sort[j] {
				temp := sort[i]
				sort[i] = sort[j]
				sort[j] = temp
			}
		}
	}

	fmt.Println(sort)
}
