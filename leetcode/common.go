package leetcode

// leetcode公共函数和结构体

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(left, right int) int {
    if left < right {
        return left
    }
    return right
}