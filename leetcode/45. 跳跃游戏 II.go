package leetcode

// 45. 跳跃游戏 II
// 给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 示例:
//
// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 说明:
//
// 假设你总是可以到达数组的最后一个位置。

// 遍历整个数组，每次到尽头了，重新根据最大距离确定边界，以贪心的方式规划出最短的路径
func jump(nums []int) int {
	maxDistance := 0
	count := 0
	end := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]+i > maxDistance {
			maxDistance = nums[i] + i
		}

		if i == end {
			count++
			end = maxDistance
		}

	}

	return count
}
