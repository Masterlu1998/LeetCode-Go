package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "  asdf asdfa dd"
	fmt.Println(strings.Split(a, " "))
}
