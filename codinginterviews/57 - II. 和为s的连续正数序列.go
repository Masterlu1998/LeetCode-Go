package codinginterviews

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//
//
//
// 示例 1：
//
// 输入：target = 9
// 输出：[[2,3,4],[4,5]]
// 示例 2：
//
// 输入：target = 15
// 输出：[[1,2,3,4,5],[4,5,6],[7,8]]
//
//
// 限制：
//
// 1 <= target <= 10^5

func findContinuousSequence(target int) [][]int {
	left, right := 0, 0

	dp := make([]int, target)
	dp[0] = 0
	result := make([][]int, 0)

	right++
	tp := []int{1}
	for right < len(dp) {
		dp[right] = dp[right-1] + right
		if dp[right]-dp[left] < target {
			right++
			tp = append(tp, right)
		} else if dp[right]-dp[left] > target {
			left++
			tp = tp[1:]
		} else {
			right++
			tp = append(tp, right)
			cp := make([]int, len(tp)-1)
			copy(cp, tp)
			result = append(result, cp)

		}
	}

	return result
}
