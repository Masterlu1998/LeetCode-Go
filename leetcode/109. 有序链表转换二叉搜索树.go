package leetcode

// 109. 有序链表转换二叉搜索树
// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 示例:

// 给定的有序链表： [-10, -3, 0, 5, 9],

// 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

//       0
//      / \
//    -3   9
//    /   /
//  -10  5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路1：二分法，每次寻找子链表的中点作为节点的值，并将链表分为中点左侧和中点右侧两个子链表，并以两个子链表的中
// 点值作为左节点和右节点，递归地构建二叉搜索树。

 func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    midNode, prev := findMid(head)

    root := new(TreeNode)
    root.Val = midNode.Val
    if prev != nil {
        prev.Next = nil
        root.Right = sortedListToBST(midNode.Next)
        root.Left = sortedListToBST(head)
    }

    return root
}

func findMid(node *ListNode) (*ListNode, *ListNode) {
    fast := node
    slow := node
    var prev *ListNode
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        prev = slow
        slow = slow.Next
    }
    return slow, prev
}