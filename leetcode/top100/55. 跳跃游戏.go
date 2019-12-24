package top100

// 给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个位置。
//
// 示例 1:
//
// 输入: [2,3,1,1,4]
// 输出: true
// 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
// 示例 2:
//
// 输入: [3,2,1,0,4]
// 输出: false
// 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。

// 思路1：填表递归遍历，瓶颈在于遍历distance所能到达的节点，很慢
var cache55 []int

func canJumpFunc1(nums []int) bool {
	cache55 = make([]int, len(nums))
	cache55[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		_canJump(nums, i)
	}

	return cache55[0] == 1
}

func _canJump(nums []int, cur int) {

	distance := nums[cur]
	for i := 1; i <= distance; i++ {
		if cache55[cur+i] == 1 {
			cache55[cur] = 1
			return
		}
	}

	cache55[cur] = -1

}

// 思路2：用一个距离表示目前最近能够到达终点的节点距离，只要比较当前跳跃范围是否比这个距离大就可以了，不需要遍历
func canJumpFunc2(nums []int) bool {
	distance := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= distance {
			distance = 1
		} else {
			distance++
		}
	}

	if distance != 1 {
		return false
	}

	return true
}
