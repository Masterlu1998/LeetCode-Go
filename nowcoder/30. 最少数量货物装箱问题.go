package nowcoder

import "fmt"

// 题目描述
// 有重量分别为3，5，7公斤的三种货物，和一个载重量为X公斤的箱子（不考虑体积等其它因素，只计算重量）
// 需要向箱子内装满X公斤的货物，要求使用的货物个数尽可能少（三种货物数量无限）
//
// 输入描述:
// 输入箱子载重量X(1 <= X <= 10000)，一个整数。
// 输出描述:
// 如果无法装满，输出 -1。
// 如果可以装满，输出使用货物的总个数。
// 示例1
// 输入
// 复制
// 4
// 输出
// 复制
// -1
// 说明
// 无法装满
// 示例2
// 输入
// 8
// 输出
// 2
// 说明
// 使用1个5公斤，1个3公斤货物

func main30() {
	x := 0
	p, err := fmt.Scan(&x)
	if p == 0 || err != nil {
		return
	}

	maxint := int(^uint(0) >> 1)
	dp := make([]int, x+1)

	for i := 0; i <= x; i++ {
		dp[i] = maxint
	}

	for i := 0; i <= x; i++ {
		if i == 3 {
			dp[3] = 1
			continue
		}

		if i == 5 {
			dp[5] = 1
			continue
		}

		if i == 7 {
			dp[7] = 1
			continue
		}

		if i-3 >= 0 && dp[i-3] != maxint {
			dp[i] = min(dp[i], dp[i-3]+1)
		}

		if i-5 >= 0 && dp[i-5] != maxint {
			dp[i] = min(dp[i], dp[i-5]+1)
		}

		if i-7 >= 0 && dp[i-7] != maxint {
			dp[i] = min(dp[i], dp[i-7]+1)
		}
	}

	if dp[x] == maxint {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[x])
	}
}
