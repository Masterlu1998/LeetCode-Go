package nowcoder

// 题目描述
// 猛兽侠中精灵鼠在利剑飞船的追逐下逃到一个n*n的建筑群中，精灵鼠从（0,0）的位置进入建筑群，建筑群的出口位置为（n-1,n-1），建筑群的每个位置都有阻碍，每个位置上都会相当于给了精灵鼠一个固定值减速，因为精灵鼠正在逃命所以不能回头只能向前或者向下逃跑，现在问精灵鼠最少在减速多少的情况下逃出迷宫？
//
// 输入描述:
// 第一行迷宫的大小: n >=2 & n <= 10000；
// 第2到n+1行，每行输入为以','分割的该位置的减速,减速f >=1 & f < 10。
// 输出描述:
// 精灵鼠从入口到出口的最少减少速度？
// 示例1
// 输入
// 3
// 5,5,7
// 6,7,8
// 2,2,4
// 输出
// 19

func handle21(matrics [][]int, total int) int {
	dp := make([][]int, total)

	tp := 0
	for i := 0; i < total; i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, total)
		}
		tp += matrics[i][0]
		dp[i][0] = tp
	}

	tp = 0
	for i := 0; i < total; i++ {
		tp += matrics[0][i]
		dp[0][i] = tp
	}

	for i := 1; i < total; i++ {
		for j := 1; j < total; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrics[i][j]
		}
	}

	return dp[total-1][total-1]
}

func min(left, right int) int {
	if left < right {
		return left
	}

	return right
}
