package leetcode

// 你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
//
// 示例 1:
//
// 输入: a = 2, b = [3]
// 输出: 8
// 示例 2:
//
// 输入: a = 2, b = [1,0]
// 输出: 1024

// a * b % c = ((a % c) * (b % c)) % c
var base = 1337

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}

	num := b[len(b)-1]

	left := myPow372(a, num)
	right := myPow372(superPow(a, b[:len(b)-1]), 10)

	return left * right % base
}

func myPow372(a, b int) int {
	if b == 0 {
		return 1
	}

	a %= base

	if b&1 == 1 {
		// 奇数
		return (a * myPow372(a, b-1)) % base
	}

	part := myPow372(a, b/2) % base

	return part * part % base
}
