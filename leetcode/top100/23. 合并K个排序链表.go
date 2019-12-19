package top100

// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
// 示例:
//
// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6

// 思路1：逐一遍历数组，找到目前最小的节点，然后让那个节点往后退
func mergeKListsFunc1(lists []*ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead

	for {
		isOver := true
		index := -1
		for i, head := range lists {
			if head == nil {
				continue
			}
			isOver = false
			if index == -1 || head.Val < lists[index].Val {
				index = i
			}
		}

		if isOver {
			break
		}

		cur.Next = lists[index]
		cur = cur.Next
		lists[index] = lists[index].Next
	}

	return preHead.Next
}

// 思路2：降维，一次合并两条链表，最终合并成一个链表
func mergeKListsFunc2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) != 1 {
		newList := make([]*ListNode, 0)
		left, right := 0, len(lists)-1

		for left < right {
			newList = append(newList, mergeTwoLists(lists[left], lists[right]))
			left++
			right--
		}

		if left == right {
			newList = append(newList, lists[left])
		}

		lists = newList
	}

	return lists[0]
}
