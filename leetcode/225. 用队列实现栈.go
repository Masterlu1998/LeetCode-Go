package leetcode

// 225. 用队列实现栈

// 使用队列实现栈的下列操作：

// push(x) -- 元素 x 入栈
// pop() -- 移除栈顶元素
// top() -- 获取栈顶元素
// empty() -- 返回栈是否为空
// 注意:

// 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
// 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
// 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

// 思路：一个队列做主存，另外一个队列做辅存，用一个变量做标记。入栈时往当前主存队列中添加数据；获取栈顶元素时，将
// 主存队列中的n-1个数据加入辅存中，最后一个即为栈顶元素；移除栈顶元素时，将主存队列中的n-1个数据加入辅存队列中，并
// 获取主存队列中最后一个元素作为返回值，最后移除主存队列中的元素，主存队列与辅存队列互换角色。

type Queue struct {
	nums []int
}

func New() Queue {
	return Queue{make([]int, 0)}
}

func (q *Queue) Push(x int) {
	q.nums = append(q.nums, x)
}

func (q *Queue) Pop() int {
	result := q.nums[0]
	q.nums = q.nums[1:]
	return result
}

func (q *Queue) Peek() int {
	result := q.nums[0]
	return result
}

func (q *Queue) Len() int {
	return len(q.nums)
}

type MyStack struct {
	queueA     Queue
	queueB     Queue
	inputQueue int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{New(), New(), 1}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.inputQueue == 0 {
		this.queueA.Push(x)
	}
	this.queueB.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.inputQueue == 1 {
		for this.queueB.Len() != 1 {
			this.queueA.Push(this.queueB.Pop())
		}
		this.inputQueue = 0
		result := this.queueB.Pop()
		return result
	}

	for this.queueA.Len() != 1 {
		this.queueB.Push(this.queueA.Pop())
	}
	this.inputQueue = 1
	result := this.queueA.Pop()
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.inputQueue == 1 {
		for this.queueB.Len() != 1 {
			this.queueA.Push(this.queueB.Pop())
		}
		result := this.queueB.Peek()
		return result
	}

	for this.queueA.Len() != 1 {
		this.queueB.Push(this.queueA.Pop())
	}
	result := this.queueA.Peek()
	return result
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.queueA.Len() == 0 && this.queueB.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
