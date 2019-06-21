package leetcode

/**
 * 接雨水
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 示例
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 */

 import (
	//  "fmt"
 )

/** 
 * 思路1：找到高度最高的数字，并以这个数字为界限从1开始遍历每一层。对于每一层使用跳桩的方法，将超过当前高度的数字
 * 视为一个桩，每找到一个桩，就与上一个桩的下标相减，得到的就是该层能够接到的雨水。（自想，效率低）
 */

 func trapFunc1(height []int) int {
    if len(height) <= 2 {
        return 0
    }
    result := 0
    maxHeight := ^int(^uint(0) >> 1)
    for i := 0; i < len(height); i++ {
        if height[i] > maxHeight {
            maxHeight = height[i]
        }
    }
    for i := 1; i <= maxHeight; i++ {
        count := 0
        prevStub := -1
        for j := 0; j < len(height); j++ {
            if height[j] >= i {
                if prevStub != -1 {
                    count += (j-prevStub-1)
                }
                prevStub = j
            }
        }
        result += count
    }
    return result
}

/**
 * 思路2：用一个栈顺次记录数组的下标。如果遇到高度比当前栈顶低则将当前下表压入栈中。如果遇到高度比当前栈顶
 * 元素高，说明可以盛水，弹栈计算接到水的体积，然后反复与栈顶比较，直到当前元素比栈顶元素小，或者栈空。之后
 * 将当前元素下标压入栈中。
 */

 func trapFunc2(height []int) int {
    result := 0
    cur := 0
    stack := make([]int, 0)
    
    for cur < len(height) {
        for len(stack) != 0 && height[cur] > height[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            distance := cur - stack[len(stack)-1] - 1
            realHeight := min(height[cur], height[stack[len(stack)-1]]) - height[top]
            result += distance * realHeight
        }
        stack = append(stack, cur)
        cur++
    }
    return result
}

