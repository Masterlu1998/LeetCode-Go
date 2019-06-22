package leetcode

// 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 示例:
// 输入: nums = [1,2,3]
// 输出:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

// 思路1：回溯法，定义不同长度的回溯上限然后用一个循环进行长度从0到数组长度的回溯

var resultSub [][]int

func subsetsFunc1(nums []int) [][]int {
	resultSub = make([][]int, 0)
	length := len(nums)
	for i := 0; i <= length; i++ {
		_subsetsFunc1(nums, []int{}, 0, i)
	}
	return resultSub
}

func _subsetsFunc1(nums, temp []int, index, length int) {
	// fmt.Println(temp, index)
	if len(temp) == length {
		final := make([]int, len(temp))
		copy(final, temp)
		resultSub = append(resultSub, final)
		return
	}

	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		_subsetsFunc1(nums, temp, i+1, length)
		temp = temp[:len(temp)-1]
	}
}

// 思路二：递归。每次将一个数字单独剥离出来，寻找剩下数组的子集，并将自己加到剩下数组子集的后面构成新的子集，最后将
// 新的子集和老的自己一起推入数组中

func subsetsFunc2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	res := make([][]int, 0)
	a := nums[0]
	b := subsetsFunc2(nums[1:])
	for _, sub := range b {
		copySub := make([]int, len(sub))
		copy(copySub, sub)
		copySub = append(copySub, a)
		res = append(res, copySub, sub)
	}
	return res
}
