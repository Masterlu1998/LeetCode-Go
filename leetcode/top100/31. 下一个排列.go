package top100

// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须原地修改，只允许使用额外常数空间。
//
// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1

// 思路：数学推导如下
// 1. 从右往左找到第一个逆序的数字下标，此时该数字右侧为升序数组，最大；如果没有逆序，说明已经是最大的排序了。
// 2. 从右侧找到第一个比该数字大的数，交换，该位为当前最小值
// 3. 反转右侧数组，最大变为最小。此时数组为当前的下一个排列

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i != -1 {
		j := i + 1
		for j < len(nums) && nums[j] > nums[i] {
			j++
		}

		nums[i], nums[j-1] = nums[j-1], nums[i]

		left, right := i+1, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
		return
	}

	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
