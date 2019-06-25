package main

import (
	"fmt"
)

type Queue []int 

func New() *Queue {
	q := make(Queue, 0)
	return &q
}

func (q *Queue)Push(x int) {
	*q = append(*q, x)
} 

func (q *Queue)Pop() int {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *Queue)Peek() int {
	return (*q)[0]
} 


func main() {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())

}

