package nowcoder

import (
	"fmt"
	"strconv"
)

func main26() {
	for {
		n := 0
		p, err := fmt.Scan(&n)
		if p == 0 || err != nil {
			return
		}

		fmt.Println(handle26(n))
	}
}

func handle26(n int) string {
	counter := n
	sum := 0
	for i := 2; i < counter; i++ {
		base := n
		for base >= i {
			sum += base % i
			base /= i
		}

		sum += base
	}

	return simplification(sum, counter-2)
}

func simplification(left, right int) string {
	originLeft, originRight := left, right
	tp := 0
	for right != 0 {
		tp = left % right
		left = right
		right = tp
	}

	return strconv.Itoa(originLeft/left) + "/" + strconv.Itoa(originRight/left)
}
