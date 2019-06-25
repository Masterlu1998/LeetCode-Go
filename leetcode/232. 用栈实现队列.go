package leetcode

// 232. 用栈实现队列

// 使用栈实现队列的下列操作：

// push(x) -- 将一个元素放入队列的尾部。
// pop() -- 从队列首部移除元素。
// peek() -- 返回队列首部的元素。
// empty() -- 返回队列是否为空。
// 示例:

// MyQueue queue = new MyQueue();

// queue.push(1);
// queue.push(2);  
// queue.peek();  // 返回 1
// queue.pop();   // 返回 1
// queue.empty(); // 返回 false
// 说明:

// 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

// 思路1：两个栈，入队时使用其中一个栈，出队的时候判断另外一个栈是否为空，如果为空，就将栈1中的元素出栈，在栈2
// 中入栈，栈顶元素就是出队的元素。

type stack struct {
    nums []int
}

func NewStack() stack {
    return stack{make([]int, 0)}
}

func (s *stack) Push(x int) {
    s.nums = append(s.nums, x)
}

func (s *stack) Pop() int {
    result := s.nums[len(s.nums)-1]
    s.nums = s.nums[:len(s.nums)-1]
    return result
}

func (s *stack) Top() int {
    return s.nums[len(s.nums)-1]
}

func (s *stack) Len() int {
    return len(s.nums)
}

type MyQueue struct {
    stack1 stack
    stack2 stack
}


/** Initialize your data structure here. */
func ConstructorQ() MyQueue {
    return MyQueue{NewStack(), NewStack()}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.stack1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if this.stack2.Len() == 0 {
        for this.stack1.Len() != 0 {
            this.stack2.Push(this.stack1.Pop())
        }
    }
    return this.stack2.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if this.stack2.Len() == 0 {
        for this.stack1.Len() != 0 {
            this.stack2.Push(this.stack1.Pop())
        }
    }
    return this.stack2.Top()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if this.stack1.Len() == 0 && this.stack2.Len() == 0 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
