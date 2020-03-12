package nowcoder

// 题目描述
// 我们定义字符串包含关系：字符串A=abc，字符串B=ab，字符串C=ac，则说A包含B，A和C没有包含关系。
// 输入描述:
// 两个字符串，判断这个两个字符串是否具有包含关系，测试数据有多组，请用循环读入。
// 输出描述:
// 如果包含输出1，否则输出0.
// 示例1
// 输入
// abc ab
// 输出
// 1

import (
	"fmt"
)

func main29() {
	for {
		left, right := "", ""
		p, err := fmt.Scan(&left, &right)
		if p == 0 || err != nil {
			return
		}

		if handle29(left, right) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

func handle29(left, right string) bool {
	base := ""
	target := ""
	if len(left) < len(right) {
		base = left
		target = right
	} else if len(left) > len(right) {
		base = right
		target = left
	} else {
		return false
	}

	nextIndex := make([]int, 1)
	for i := 1; i < len(base); i++ {
		left, right := 0, i
		sameCounter := 0
		for left < right && base[left] == base[right] {
			sameCounter++
			left++
			right--
		}
		nextIndex = append(nextIndex, sameCounter)
	}

	copy(nextIndex[1:], nextIndex[:len(nextIndex)-1])
	nextIndex[0] = -1

	targetIndex, baseIndex := 0, 0
	for targetIndex < len(target) {
		if target[targetIndex] == base[baseIndex] {
			targetIndex++
			baseIndex++
			if baseIndex == len(base) {
				return true
			}
			continue
		}

		if nextIndex[baseIndex] == -1 {
			targetIndex++
			baseIndex = 0
		} else {
			baseIndex = nextIndex[baseIndex]
		}
	}

	return false
}
