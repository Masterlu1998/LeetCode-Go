package top100

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 思路1：两次遍历，一次用map记录所有的数，二次遍历的时候用target减当前值在map中查询，要求index与当前
// 不相同即可
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i, n := range nums {
		cache[n] = i
	}

	for cur, n := range nums {
		if index, ok := cache[target-n]; ok {
			if cur != index {
				return []int{index, cur}
			}
		}
	}

	return nil
}
