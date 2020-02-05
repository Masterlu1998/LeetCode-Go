package codinginterviews

import "fmt"

// 计算整数二进制中1的数量

func numberOf1Func1(n int) int {
	count := 0
	// 该操作用于去除最右边的1，有几个1做几次操作
	for n != 0 {
		n = (n - 1) & n
		count++
	}

	return count
}

func numberOf1Func2(n int) int {
	count := 0
	tmp := 1
	for tmp != 0 {
		if n&tmp == tmp {
			count++
		}
		tmp = tmp << 1
	}

	return count
}

func TestNumberOf1() {
	num := 2222
	if numberOf1Func1(num) == numberOf1Func2(num) {
		fmt.Println(numberOf1Func1(num), numberOf1Func2(num))
	}
}
