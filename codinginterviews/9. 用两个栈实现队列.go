package codinginterviews

type CQueue struct {
	stack1 []int // output
	stack2 []int // input
}

func Constructor9() CQueue {
	return CQueue{make([]int, 0), make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.stack2 = append(this.stack2, value)
}

func (this *CQueue) DeleteHead() int {

	if len(this.stack1) == 0 {
		if len(this.stack2) == 0 {
			return -1
		}
		for len(this.stack2) != 0 {
			this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
			this.stack2 = this.stack2[:len(this.stack2)-1]
		}
	}

	result := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[:len(this.stack1)-1]
	return result
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor9();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
