package algorithm_map

// it will generate algorithm_map like this
//
// 0---3
// | \ |
// 1---2

func NewAdjacencyMatrixUndirectedGraph() ([]int, [][]int) {
	return []int{0, 1, 2, 3}, [][]int{{0, 1, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 1}, {1, 0, 1, 0}}
}

type UndirectedNode struct {
	Index int
	Next  *UndirectedNode
}

func NewAdjacencyTableUndirectedGraph() []*UndirectedNode {
	node3 := &UndirectedNode{3, nil}
	node2 := &UndirectedNode{2, node3}
	node1 := &UndirectedNode{1, node2}
	node5 := &UndirectedNode{2, nil}
	node4 := &UndirectedNode{0, node5}
	node8 := &UndirectedNode{3, nil}
	node7 := &UndirectedNode{1, node8}
	node6 := &UndirectedNode{0, node7}
	node10 := &UndirectedNode{2, nil}
	node9 := &UndirectedNode{0, node10}

	node11 := &UndirectedNode{0, node1}
	node12 := &UndirectedNode{1, node4}
	node13 := &UndirectedNode{2, node6}
	node14 := &UndirectedNode{3, node9}

	return []*UndirectedNode{node11, node12, node13, node14}
}

