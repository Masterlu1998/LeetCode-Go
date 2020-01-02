package top100

// 现在你总共有 n 门课需要选，记为 0 到 n-1。
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
//
// 给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
//
// 示例 1:
//
// 输入: 2, [[1,0]]
// 输出: true
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
// 示例 2:
//
// 输入: 2, [[1,0],[0,1]]
// 输出: false
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
// 说明:
//
// 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
// 提示:
//
// 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
// 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
// 拓扑排序也可以通过 BFS 完成。

// dfs拓扑排序
var flag []int
var metrics [][]bool

func canFinishFunc1(numCourses int, prerequisites [][]int) bool {
	flag = make([]int, numCourses)
	metrics = make([][]bool, numCourses)

	// fill metric
	for i := 0; i < numCourses; i++ {
		if metrics[i] == nil {
			metrics[i] = make([]bool, numCourses)
		}
	}

	for _, pairs := range prerequisites {
		metrics[pairs[1]][pairs[0]] = true
	}

	// dfs200 every node
	for i := 0; i < numCourses; i++ {
		if !dfs207(i) {
			return false
		}
	}

	return true
}

func dfs207(i int) bool {
	if flag[i] == 1 {
		return false
	}

	if flag[i] == -1 {
		return true
	}

	flag[i] = 1
	for j := 0; j < len(metrics); j++ {
		if metrics[i][j] && !(dfs207(j)) {
			return false
		}
	}

	flag[i] = -1
	return true
}

// bfs
type queue []int

func NewQueue() *queue {
	q := queue(make([]int, 0))
	return &q
}

func (q *queue) InQueue(i int) {
	*q = append(*q, i)
}

func (q *queue) DeQueue() int {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *queue) Empty() bool {
	return len(*q) == 0
}

var table []int

func canFinishFunc2(numCourses int, prerequisites [][]int) bool {
	table = make([]int, numCourses)

	for _, pairs := range prerequisites {
		table[pairs[0]]++
	}

	q := NewQueue()

	for i := 0; i < numCourses; i++ {
		if table[i] == 0 {
			q.InQueue(i)
		}
	}

	for !q.Empty() {
		last := q.DeQueue()
		numCourses--
		for _, pairs := range prerequisites {
			if pairs[1] != last {
				continue
			}

			table[pairs[0]]--
			if table[pairs[0]] == 0 {
				q.InQueue(pairs[0])
			}
		}
	}

	return numCourses == 0
}
