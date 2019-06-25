package leetcode

// 50. Pow(x, n)

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。

// 示例 1:

// 输入: 2.00000, 10
// 输出: 1024.00000
// 示例 2:

// 输入: 2.10000, 3
// 输出: 9.26100
// 示例 3:

// 输入: 2.00000, -2
// 输出: 0.25000
// 解释: 2-2 = 1/22 = 1/4 = 0.25
// 说明:

// -100.0 < x < 100.0
// n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

// 终极思路：动归问题，a的n次方实际上是a的n/2次方的平方，可得到状态转移方程式：n为偶数 a^n = a^(n/2) * a^(n/2)
// n为奇数 a^n = a^(n/2) * a^(n/2) * a。需要注意的是当n为负数的时候，需要将答案变为倒数。代码中还使用了位运算
// 来计算奇偶，提高代码效率。因为范围的问题，这里不能使用常规的方法来加速动归。

func myPow(x float64, n int) float64 {
    if n == 1 {
        return x
    }
    
    if n == 0 {
        return 1
    }
    copyn := n
    
    if n < 0 {
        copyn *= -1
    }
    
    result := myPow(x, copyn / 2)
    result *= result
    if copyn & 1 == 1 {
        result *= x
    }
    
    if n < 0 {
        return 1 / result
    }
    return result
}