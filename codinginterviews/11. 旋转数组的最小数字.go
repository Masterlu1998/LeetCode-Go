package codinginterviews

// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
//
// 示例 1：
//
// 输入：[3,4,5,1,2]
// 输出：1
// 示例 2：
//
// 输入：[2,2,2,0,1]
// 输出：0

func minArray(numbers []int) int {
	length := len(numbers)

	if length == 1 {
		return numbers[0]
	}

	if length == 2 {
		return min(numbers[0], numbers[1])
	}

	if numbers[0] < numbers[length-1] {
		return numbers[0]
	}

	mid := length / 2
	if numbers[mid] < numbers[0] {
		return minArray(numbers[:mid+1])
	}

	if numbers[mid] > numbers[length-1] {
		return minArray(numbers[mid+1:])
	}

	return min(minArray(numbers[:mid+1]), minArray(numbers[mid+1:]))
}

func min(left, right int) int {
	if left > right {
		return right
	}

	return left
}
