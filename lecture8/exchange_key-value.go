package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e",
		6: "f", 7: "g", 8: "h", 9: "i", 0: "j"}
	fmt.Println("m1", m1)
	m2 := make(map[string]int)
	for i, v := range m1 {
		m2[v] = i
	}
	fmt.Println("m2", m2)
}
