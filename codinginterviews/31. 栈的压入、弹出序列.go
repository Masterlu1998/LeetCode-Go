package codinginterviews

// 给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的
// 推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。

type stack31 []int

func newStack() *stack31 {
	st := stack31(make([]int, 0))
	return &st
}

func (st *stack31) Push(i int) {
	*st = append(*st, i)
}

func (st *stack31) Pop() int {
	result := (*st)[len(*st)-1]
	(*st) = (*st)[:len(*st)-1]
	return result
}

func (st *stack31) Get() int {
	return (*st)[len(*st)-1]
}

func (st *stack31) Empty() bool {
	return len(*st) == 0
}

// 遍历出栈顺序，如果元素在栈顶，弹出，如果元素不在栈顶，就将推入顺序元素依次推入栈中，如果没找到，就说明
// 有问题，如果找到了就持续上述过程
func validateStackSequences(pushed []int, popped []int) bool {
	st := newStack()
	for _, num := range popped {
		for st.Empty() || num != st.Get() {
			if len(pushed) == 0 {
				return false
			}

			st.Push(pushed[0])
			pushed = pushed[1:]
		}

		st.Pop()
	}

	return true
}
