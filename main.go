package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{1, 3, 6, 2}
	sort.Ints(a)
	fmt.Println(a)
}

