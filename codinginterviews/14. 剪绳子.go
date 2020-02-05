package codinginterviews

import "fmt"

// 给定一段绳子，剪成若干段绳子，使得每段绳子的长度相乘结果最大，返回结果

// 思路1：动态规划
func maxProductAfterCutting(length int) int {
	if length <= 1 {
		return 0
	}

	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}

	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= length; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			temp := dp[j] * dp[i-j]
			if temp > max {
				max = temp
			}
		}
		dp[i] = max
	}
	fmt.Println(dp)
	return dp[length]
}

func TestMaxProductAfterCutting() {
	fmt.Println(maxProductAfterCutting(10))
}
