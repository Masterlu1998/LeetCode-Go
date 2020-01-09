package top100

// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
// 示例 1 :
//
// 输入:nums = [1,1,1], k = 2
// 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
// 说明 :
//
// 数组的长度为 [1, 20,000]。
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

// 利用累计和，计算到每一个索引的和，然后用这个和减去目标值k，如果这个值已经出现过了，说明两者之间的值
// 相加为0,说明sum - k = 0，即k = sum
func subarraySum(nums []int, k int) int {
	cache := map[int]int{0: 1}
	count := 0

	temp := 0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		if total, ok := cache[temp-k]; ok {
			count += total
		}
		cache[temp]++
	}
	return count
}
