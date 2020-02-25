package codinginterviews

// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//
//
// 示例:
//
// 输入: [1,2,3,4,5]
// 输出: [120,60,40,30,24]

func constructArr(a []int) []int {
	res := make([]int, len(a))

	last := 1
	for i := 0; i < len(a); i++ {
		res[i] = last
		last *= a[i]
	}

	last = 1
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= last
		last *= a[i]
	}

	return res
}
