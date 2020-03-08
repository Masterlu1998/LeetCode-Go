package nowcoder

type temp struct {
	candies []int
	val     int
}

type candy struct {
	seq int
	p   int
}

func handle22(c []candy, v int) (int, []int) {
	dp := make([]temp, v+1)

	for i := 0; i < len(c); i++ {
		for j := v; j-c[i].seq >= 0; j-- {
			if dp[j].val < dp[j-c[i].seq].val+c[i].p {
				dp[j].val = dp[j-c[i].seq].val + c[i].p
				dp[j].candies = append(dp[j-c[i].seq].candies, i+1)
			}
		}
	}

	return dp[v].val, dp[v].candies
}
