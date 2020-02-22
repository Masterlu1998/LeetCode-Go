package codinginterviews

func printNumbers(n int) []int {
	result := make([]int, 0)
	if n == 0 {
		return result
	}

	max := 9

	for i := 1; i < n; i++ {
		max = max*10 + 9
	}

	for i := 1; i <= max; i++ {
		result = append(result, i)
	}
	return result
}
