package leetcode

// 输入一个按升序排序的整数数组（可能包含重复数字），你需要将它们分割成几个子序列，其中每个子序列至少包含三个连续整数。返回你是否能做出这样的分割？
//
//
//
// 示例 1：
//
// 输入: [1,2,3,3,4,5]
// 输出: True
// 解释:
// 你可以分割出这样两个连续子序列 :
// 1, 2, 3
// 3, 4, 5
//
//
// 示例 2：
//
// 输入: [1,2,3,3,4,4,5,5]
// 输出: True
// 解释:
// 你可以分割出这样两个连续子序列 :
// 1, 2, 3, 4, 5
// 3, 4, 5
//
//
// 示例 3：
//
// 输入: [1,2,3,4,4,5]
// 输出: False
//
//
// 提示：
//
// 输入的数组长度范围为 [1, 10000]

// 先对所有数字计数，然后优先把当前数加入之前的链中，如果不行，就自己开一个长度为3的链，如果不行就
// 不可能满足题意。
func isPossible(nums []int) bool {
	counter := make(map[int]int, 0)
	tail := make(map[int]int, 0)

	for _, num := range nums {
		counter[num] = counter[num] + 1
	}

	for _, num := range nums {
		if counter[num] == 0 {
			continue
		} else if tail[num] > 0 {
			tail[num] -= 1
			tail[num+1] += 1
		} else if counter[num+1] > 0 && counter[num+2] > 0 {
			counter[num+1] -= 1
			counter[num+2] -= 1
			tail[num+3] += 1
		} else {
			return false
		}
		counter[num] -= 1
	}

	return true
}
