package codinginterviews

import "sort"

// 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
//
//
// 示例 1:
//
// 输入: [1,2,3,4,5]
// 输出: True
//
//
// 示例 2:
//
// 输入: [0,0,1,2,5]
// 输出: True
//
//
// 限制：
//
// 数组长度为 5
//
// 数组的数取值为 [0, 13] .

func isStraight(nums []int) bool {
	sort.Ints(nums)

	cur := 0
	zeroCounter, base, missCounter := 0, false, 0
	for cur < len(nums) {
		if nums[cur] == 0 {
			zeroCounter++
			cur++
			continue
		}

		if !base {
			base = true
		} else {
			if nums[cur] == nums[cur-1] {
				return false
			}
			missing := nums[cur] - nums[cur-1]
			if missing > 1 {
				missCounter += missing - 1
			}
		}

		cur++
	}

	if missCounter != 0 {
		return zeroCounter >= missCounter
	}

	return true
}
