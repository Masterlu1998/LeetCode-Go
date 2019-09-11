package main

import (
	"fmt"

	"algorithm-go/algorithm_map"
)

func main() {
	algorithm_map.TestAdjacencyMatrixDFS()
	fmt.Println("//////////////")
	algorithm_map.TestAdjacencyTableDFS()
	fmt.Println("//////////////")
	algorithm_map.TestAdjacencyMatrixBFS()
	fmt.Println("//////////////")
	algorithm_map.TestAdjacencyTableBFS()
}
