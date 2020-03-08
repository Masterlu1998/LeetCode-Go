package nowcoder

// 现在有一棵合法的二叉树，树的节点都是用数字表示，现在给定这棵树上所有的父子关系，求这棵树的高度

import "fmt"

type queue24 []int

func newQueue() *queue24 {
	q := queue24(make([]int, 0))
	return &q
}

func (q *queue24) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *queue24) Dequeue() {
	*q = (*q)[1:]
}

func (q *queue24) Get() int {
	return (*q)[0]
}

func (q *queue24) Empty() bool {
	return len(*q) == 0
}

func (q *queue24) Length() int {
	return len(*q)
}

func handle24Main() {

	for {
		total := 0
		p, err := fmt.Scan(&total)
		if p == 0 || err != nil {
			return
		}

		relationship := make([][]int, 0)

		for i := 0; i < total-1; i++ {
			father, child := 0, 0
			p, err := fmt.Scan(&father, &child)
			if p == 0 || err != nil {
				return
			}

			relationship = append(relationship, []int{father, child})
		}

		fmt.Println(handle24(relationship, total))
	}
}

func handle24(relationship [][]int, total int) int {
	if total == 1 {
		return 1
	}

	nodes := make([]int, total)
	cache := make(map[int]bool)
	reMap := make(map[int][]int)

	for _, r := range relationship {
		reMap[r[0]] = append(reMap[r[0]], r[1])
		cache[r[0]] = true
	}

	q := newQueue()
	for _, r := range relationship {
		delete(cache, r[1])
	}

	for index := range cache {
		nodes[index] = 1
		q.Enqueue(index)
	}

	deep := 0
	for !q.Empty() {
		temp := q.Length()
		for i := 0; i < temp; i++ {
			father := q.Get()
			q.Dequeue()
			for _, child := range reMap[father] {
				q.Enqueue(child)
			}
		}
		deep++
	}

	return deep
}
