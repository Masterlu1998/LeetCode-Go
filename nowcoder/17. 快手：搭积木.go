package nowcoder

import "sort"

// 题目描述
// 小明有一袋子长方形的积木，如果一个积木A的长和宽都不大于另外一个积木B的长和宽，则积木A可以搭在积木B的上面。好奇的小明特别想知道这一袋子积木最多可以搭多少层，你能帮他想想办法吗？
// 定义每一个长方形的长L和宽W都为正整数，并且1 <= W <= L <= INT_MAX, 袋子里面长方形的个数为N, 并且 1 <= N <= 1000000.
// 假如袋子里共有5个积木分别为 (2, 2), (2, 4), (3, 3), (2, 5), (4, 5), 则不难判断这些积木最多可以搭成4层, 因为(2, 2) < (2, 4) < (2, 5) < (4, 5)。
// 输入描述:
// 第一行为积木的总个数 N
//
// 之后一共有N行，分别对应于每一个积木的宽W和长L
// 输出描述:
// 输出总共可以搭的层数
// 示例1
// 输入
// 复制
// 5
// 2 2
// 2 4
// 3 3
// 2 5
// 4 5
// 输出
// 复制
// 4

func handle17(matrics [][]int) int {
	sort.Slice(matrics, func(left, right int) bool {
		if matrics[left][0] == matrics[right][0] {
			return matrics[left][1] < matrics[right][1]
		}

		return matrics[left][0] < matrics[right][0]
	})

	dp := make([]int, len(matrics))
	length := 0

	for _, num := range matrics {
		left, right := 0, length
		for left < right {
			mid := (left + right) / 2
			if dp[mid] > num[1] {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == length {
			length++
		}
		dp[left] = num[1]
	}

	return length
}
