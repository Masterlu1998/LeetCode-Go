package nowcoder

import (
	"math"
)

// 题目描述
// 牛牛准备参加学校组织的春游, 出发前牛牛准备往背包里装入一些零食, 牛牛的背包容量为w。
// 牛牛家里一共有n袋零食, 第i袋零食体积为v[i]。
// 牛牛想知道在总体积不超过背包容量的情况下,他一共有多少种零食放法(总体积为0也算一种放法)。
// 输入描述:
// 输入包括两行
// 第一行为两个正整数n和w(1 <= n <= 30, 1 <= w <= 2 * 10^9),表示零食的数量和背包的容量。
// 第二行n个正整数v[i](0 <= v[i] <= 10^9),表示每袋零食的体积。
// 输出描述:
// 输出一个正整数, 表示牛牛一共有多少种零食放法。
// 示例1
// 输入
// 3 10
// 1 2 4
// 输出
// 8
// 说明
// 三种零食总体积小于10,于是每种零食有放入和不放入两种情况，一共有2*2*2 = 8种情况。

func handle_8(n, w int, v []int) int {
	res := 0
	sum := 0
	for _, val := range v {
		sum += val
	}

	if sum <= w {
		return int(math.Pow(float64(2), float64(n)))
	}

	var _handle func(v []int, cur, temp int)
	_handle = func(v []int, cur, temp int) {
		if temp > w {
			return
		}

		res++

		for i := cur; i < len(v); i++ {
			_handle(v, i+1, temp+v[i])
		}
	}

	_handle(v, 0, 0)
	return res
}
