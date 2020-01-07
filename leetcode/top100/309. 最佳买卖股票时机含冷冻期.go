package top100

// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 示例:
//
// 输入: [1,2,3,0,2]
// 输出: 3
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

func maxProfitWithCoolFunc1(prices []int) int {
	hold, sold, rest := ^int(^uint(0)>>1), 0, 0

	for i := 0; i < len(prices); i++ {
		temp := sold
		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, temp)
	}

	return max(sold, rest)
}

// 暴力回溯法，超时
func maxProfitWithCoolFunc2(prices []int) int {
	result := ^int(^uint(0) >> 1)
	var _maxProfit func(prices []int, freeze, have bool, temp int)
	_maxProfit = func(prices []int, freeze, have bool, temp int) {
		if len(prices) == 0 {
			if temp > result {
				result = temp
			}
			return
		}

		if !have && !freeze {
			_maxProfit(prices[1:], false, true, temp-prices[0])
		}

		if have {
			_maxProfit(prices[1:], true, false, temp+prices[0])
		}

		_maxProfit(prices[1:], false, have, temp)

	}
	_maxProfit(prices, false, false, 0)
	return result
}
