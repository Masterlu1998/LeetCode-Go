package codinginterviews

import "strconv"

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//
// 示例 1:
//
// 输入: 12258
// 输出: 5
// 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

func translateNum(num int) int {
	count := 0
	var _translateNum func(num string)
	_translateNum = func(num string) {
		if len(num) == 0 {
			count++
			return
		}

		for i := 1; i <= len(num); i++ {
			tmpStr := num[:i]
			if len(tmpStr) > 1 && tmpStr[0] == '0' {
				break
			}

			tmp, _ := strconv.Atoi(tmpStr)
			if tmp > 25 {
				break
			}

			_translateNum(num[i:])
		}
	}
	_translateNum(strconv.Itoa(num))
	return count
}
