package codinginterviews

// 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
//
//
// 示例 1：
//
// 输入：nums = [4,1,4,6]
// 输出：[1,6] 或 [6,1]
// 示例 2：
//
// 输入：nums = [1,2,10,4,1,4,3,3]
// 输出：[2,10] 或 [10,2]
//
//
// 限制：
//
// 2 <= nums <= 10000

func singleNumbers(nums []int) []int {
	tp := nums[0]
	for i := 1; i < len(nums); i++ {
		tp = tp ^ nums[i]
	}

	bit := 0
	for i := 1; i != 0; i = i << 1 {
		if tp&i == i {
			bit = i
		}
	}

	if bit == 0 {
		return nil
	}

	left, right := make([]int, 0), make([]int, 0)
	for _, num := range nums {
		if num&bit == bit {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	lr := left[0]
	for i := 1; i < len(left); i++ {
		lr = lr ^ left[i]
	}

	rr := right[0]
	for i := 1; i < len(right); i++ {
		rr = rr ^ right[i]
	}

	return []int{lr, rr}
}
