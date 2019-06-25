package codinginterviews

import (
	"fmt"
)

// 12. 打印1到最大的n位数

// 输入数字n，按照顺序打印从1到最大的n位十进制数

// 示例1：

// 输入：3
// 输出：1， 2， 3...， 998， 999

// 思路：对n位进行0-9的排列组合，这可能是一个大数问题，需要用字符串来解决。

func printToMaxOfNDigits(n int) {
	_printToMaxOfNDigits(n, "")
}

func _printToMaxOfNDigits(n int, tmp string) {
	if len(tmp) == n {
		start := n-1
		for i := 0; i < len(tmp); i++ {
			if tmp[i] != '0' {
				start = i
				break
			}
		}
		fmt.Println(tmp[start:])
		return
	}
	for i := '0'; i <= '9'; i++ {
		_printToMaxOfNDigits(n, tmp + string(i))
	}
}

func TestPrintToMaxOfNDigits() {
	printToMaxOfNDigits(5)
}