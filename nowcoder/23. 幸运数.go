package nowcoder

import "fmt"

// 小明同学学习了不同的进制之后，拿起了一些数字做起了游戏。小明同学知道，在日常生活中我们最常用的是十进制数，而在计算机中，二进制数也很常用。现在对于一个数字x，小明同学定义出了两个函数f(x)和g(x)。
// f(x)表示把x这个数用十进制写出后各个数位上的数字之和。如f(123)=1+2+3=6。
// g(x)表示把x这个数用二进制写出后各个数位上的数字之和。如123的二进制表示为1111011，那么g(123)=1+1+1+1+0+1+1=6。
// 小明同学发现对于一些正整数x满足f(x)=g(x)，他把这种数字称为幸运数，现在他想知道，小于等于n的幸运数有多少个。

var (
	cache       map[int]int
	cacheBinary map[int]int
	maxn        = 100009
)

func handle23() {
	answers := make([]int, maxn+1)
	cache, cacheBinary = make(map[int]int), make(map[int]int)

	for i := 1; i < len(answers); i++ {
		if numSum(i) == numSumBinary(i) {
			answers[i] = answers[i-1] + 1
		} else {
			answers[i] = answers[i-1]
		}
	}

	for {
		total := 0
		p, err := fmt.Scan(&total)
		if p == 0 || err != nil {
			return
		}

		for i := 0; i < total; i++ {
			n := 0
			p, err = fmt.Scan(&n)
			if p == 0 || err != nil {
				return
			}
			fmt.Println(answers[n])
		}

	}
}

func numSum(n int) int {
	tp := n
	if val, ok := cache[n]; ok {
		return val
	}

	res := 0
	for n > 0 {
		res += n % 10
		n = n / 10
	}
	cache[tp] = res
	return res
}

func numSumBinary(n int) int {
	tp := n

	if val, ok := cacheBinary[n]; ok {
		return val
	}

	res := 0
	for n != 0 {
		res++
		n = (n - 1) & n
	}

	cacheBinary[tp] = res
	return res
}
