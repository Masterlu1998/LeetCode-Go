package main

import (
	"fmt"
)

func main() {
	p := []int{1, 5, 4, 7, 3, 7}

	a := append(p[:1], p[2:]...)
	
	fmt.Println(p, a)
}
