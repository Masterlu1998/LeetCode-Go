package main

import (
	"fmt"

)

func main() {
	a := []int{1, 2, 3, 4, 56}
	c := append(a[:3], a[4:]...)
	fmt.Println(a, c)
}

