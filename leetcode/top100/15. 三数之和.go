package top100

import "sort"

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
// 满足要求的三元组集合为：
// [
//  [-1, 0, 1],
//  [-1, -1, 2]
// ]

// 思路1：排序+三指针法，目标数放最左边，剩下套用两数之和的算法。注意跳过相同的数字去重
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)

	for k := 0; k < len(nums)-1; k++ {
		if nums[k] > 0 {
			return result
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		left, right := k+1, len(nums)-1

		for left < right {
			if nums[left]+nums[right] < -nums[k] {
				left++
			} else if nums[left]+nums[right] > -nums[k] {
				right--
			} else {
				result = append(result, []int{nums[left], nums[right], nums[k]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return result
}
