package codinginterviews

type MinStack struct {
	stack1 []int // min存储
	stack2 []int // 数据存储
}

/** initialize your data structure here. */
func Constructor30() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack2 = append(this.stack2, x)
	if len(this.stack1) == 0 || x <= this.stack1[len(this.stack1)-1] {
		this.stack1 = append(this.stack1, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack1) != 0 && len(this.stack2) != 0 && this.stack2[len(this.stack2)-1] == this.stack1[len(this.stack1)-1] {
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	this.stack2 = this.stack2[:len(this.stack2)-1]
}

func (this *MinStack) Top() int {
	return this.stack2[len(this.stack2)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack1[len(this.stack1)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
