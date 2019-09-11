package algorithm_map

import "fmt"

var (
	MatrixVisited      []bool
	MatrixGlobalTable  [][]int
	MatrixGlobalPoints []int
)

func AdjacencyMatrixDFS(points []int, table [][]int) {
	MatrixVisited = make([]bool, len(points))
	MatrixGlobalPoints = points
	MatrixGlobalTable = table

	for i := 0; i < len(MatrixGlobalPoints); i++ {
		if !MatrixVisited[i] {
			_AdjacencyMatrixDFS(MatrixGlobalPoints[i])
		}
	}

}

func _AdjacencyMatrixDFS(j int) {
	fmt.Println(MatrixGlobalPoints[j])
	MatrixVisited[j] = true
	for i := 0; i < len(MatrixGlobalPoints); i++ {
		if !MatrixVisited[i] && MatrixGlobalTable[j][i] == 1 {
			_AdjacencyMatrixDFS(i)
		}
	}
}

func TestAdjacencyMatrixDFS() {
	points, table := NewAdjacencyMatrixUndirectedGraph()
	AdjacencyMatrixDFS(points, table)
}

var (
	TableVisited []bool
	GlobalTable []*UndirectedNode
)

func AdjacencyTableDFS(table []*UndirectedNode) {
	TableVisited = make([]bool, len(table))
	GlobalTable = table

	for i := 0; i < len(table); i++ {
		if !TableVisited[i] {
			_AdjacencyTableDFS(i)
		}
	}
}

func _AdjacencyTableDFS(i int) {
	TableVisited[i] = true

	cur := GlobalTable[i]
	fmt.Println(cur.Index)

	next := cur.Next
	for next != nil {
		if !TableVisited[next.Index] {
			_AdjacencyTableDFS(next.Index)
		} else {
			next = next.Next
		}
	}
}

func TestAdjacencyTableDFS() {
	table := NewAdjacencyTableUndirectedGraph()
	AdjacencyTableDFS(table)
}

type intQueue []int

func NewQueue() *intQueue {
	return &intQueue{}
}

func (q *intQueue) enQueue(item int) {
	*q = append(*q, item)
}

func (q *intQueue) deQueue() int {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func AdjacencyMatrixBFS(points []int, table [][]int) {
	q := NewQueue()
	visited := make([]bool, len(points))

	for i := 0; i < len(points); i++ {
		if !visited[i] {
			q.enQueue(points[i])
			for len(*q) != 0 {
				cur := q.deQueue()
				if !visited[cur] {
					visited[cur] = true
					fmt.Println(cur)
				}
				for j := 0; j < len(points); j++ {
					if table[i][j] == 1 && !visited[j] {
						q.enQueue(points[j])
					}
				}
			}
		}
	}
}

func TestAdjacencyMatrixBFS() {
	points, table := NewAdjacencyMatrixUndirectedGraph()
	AdjacencyMatrixBFS(points, table)
}

type nodeQueue []*UndirectedNode

func NewNodeQueue() *nodeQueue {
	return &nodeQueue{}
}

func (q *nodeQueue) enNodeQueue(item *UndirectedNode) {
	*q = append(*q, item)
}

func (q *nodeQueue) deQueue() *UndirectedNode {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func AdjacencyTableBFS(table []*UndirectedNode) {
	visited := make([]bool, len(table))
	q := NewNodeQueue()

	for _, node := range table {
		if !visited[node.Index] {
			q.enNodeQueue(node)
			for len(*q) != 0 {
				cur := q.deQueue()
				if !visited[cur.Index] {
					fmt.Println(cur.Index)
					visited[cur.Index] = true
				}

				for cur.Next != nil {
					cur = cur.Next
					q.enNodeQueue(cur)
				}
			}
		}
	}
}

func TestAdjacencyTableBFS() {
	table := NewAdjacencyTableUndirectedGraph()
	AdjacencyTableBFS(table)
}

