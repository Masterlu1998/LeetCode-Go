package stack

type stack []int

func newStack() *stack {
	s := stack(make([]int, 0))
	return &s
}

func (s *stack) Push(i int) {
	*s = append(*s, i)
}

func (s *stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack) Get() int {
	return (*s)[len(*s)-1]
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}
