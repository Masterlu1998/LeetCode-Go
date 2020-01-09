package top100

// 动态规划：dp[i][j]首先s[i]==s[j]
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	result := 0
	for i := 0; i < len(s); i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]bool, len(s))
		}
		for j := i; j >= 0; j-- {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				result++
			}
		}
	}

	return result
}
