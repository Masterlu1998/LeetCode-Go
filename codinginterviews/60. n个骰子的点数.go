package codinginterviews

import "math"

// 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
//
//
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

func twoSum60(n int) []float64 {
	dp := []int{0, 1, 1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		max, min := i*6, i*1
		tp := make([]int, max+1)
		for j := min; j <= max; j++ {
			count := 0
			for k := 1; k <= 6; k++ {
				if j-k > 0 && j-k < len(dp) {
					count += dp[j-k]
				}
			}
			tp[j] = count
		}
		dp = tp
	}

	res := make([]float64, 0)
	all := math.Pow(6.0, float64(n))
	for i := n * 1; i <= n*6; i++ {
		res = append(res, float64(dp[i])/all)
	}

	return res
}
