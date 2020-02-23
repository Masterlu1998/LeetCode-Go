package codinginterviews

// 我们把只包含因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
//
//
// 示例:
//
// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
// 说明:
//
// 1 是丑数。
// n 不超过1690。

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1

	i1, i2, i3 := 0, 0, 0

	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[i1]*2, dp[i2]*3), dp[i3]*5)
		if dp[i] == dp[i1]*2 {
			i1++
		}

		if dp[i] == dp[i2]*3 {
			i2++
		}

		if dp[i] == dp[i3]*5 {
			i3++
		}
	}

	return dp[n-1]
}
