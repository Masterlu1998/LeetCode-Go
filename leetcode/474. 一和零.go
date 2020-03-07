package leetcode

// 在计算机界中，我们总是追求用有限的资源获取最大的收益。
//
// 现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。
//
// 你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。
//
// 注意:
//
// 给定 0 和 1 的数量都不会超过 100。
// 给定字符串数组的长度不会超过 600。
// 示例 1:
//
// 输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
// 输出: 4
//
// 解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。
// 示例 2:
//
// 输入: Array = {"10", "0", "1"}, m = 1, n = 1
// 输出: 2
//
// 解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。

func findMaxForm(strs []string, m int, n int) int {
	// m -> 0 n -> 1
	dp := make([][]int, m+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zero, one := countZeroOne(strs[i])

		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j][k], 1+dp[j-zero][k-one])
			}
		}
	}

	return dp[m][n]
}

func countZeroOne(s string) (int, int) {
	zero, one := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		} else if s[i] == '0' {
			zero++
		}
	}

	return zero, one
}

func max(left, right int) int {
	if left > right {
		return left
	}

	return right
}
