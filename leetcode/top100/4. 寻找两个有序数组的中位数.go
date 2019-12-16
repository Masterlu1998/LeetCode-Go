package top100

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// 示例 1:
//
// nums1 = [1, 3]
// nums2 = [2]
//
// 则中位数是 2.0
// 示例 2:
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// 则中位数是 (2 + 3)/2 = 2.5

// 思路1：数学推导起手。nums1长度为m，nums2长度为n，将nums1，nums2分别在i，j处分开，只要满足下列
// 条件i即为中位数分割点。
// 1. nums1[i-1] < nums2[j] && nums2[j-1] < nums1[i]
// 2. i + j = m + n + 1 - i - j （奇数时）
//    i + j = m + n - i - j （偶数时）
// 简化条件2，考虑奇数的情况下得 i = 0 ~ m; j = (m + n + 1) / 2 - i ，为了防止j出现负数，n > m
// 所以根据上述条件寻找i的值，这里可以使用二分法进行查找，如果不满足条件1的前半部分说明i太大，如果
// 不满足条件1的后半部分说明i太小，根据情况二分即可。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if n < m {
		n, m = m, n
		nums1, nums2 = nums2, nums1
	}

	iMin, iMax := 0, m
	halfLength := (m + n + 1) / 2

	for {
		i := (iMin + iMax) / 2
		j := halfLength - i

		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else {
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = float64Max(float64(nums2[j-1]), float64(nums1[i-1]))
			}

			if (m+n)%2 == 1 {
				return maxLeft
			}

			var minRight float64
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = float64Min(float64(nums2[j]), float64(nums1[i]))
			}

			return (maxLeft + minRight) / 2
		}
	}

}

func float64Max(left, right float64) float64 {
	if left > right {
		return left
	}

	return right
}

func float64Min(left, right float64) float64 {
	if left > right {
		return right
	}

	return left
}
