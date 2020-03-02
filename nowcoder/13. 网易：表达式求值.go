package nowcoder

// 题目描述
// 今天上课，老师教了小易怎么计算加法和乘法，乘法的优先级大于加法，但是如果一个运算加了括号，那么它的优先级是最高的。例如：
// 1+2*3=7
// 1*(2+3)=5
// 1*2*3=6
// (1+2)*3=9
// 现在小易希望你帮他计算给定3个数a，b，c，在它们中间添加"+"， "*"， "("， ")"符号，能够获得的最大值。
// 输入描述:
// 一行三个数a，b，c (1 <= a, b, c <= 10)
// 输出描述:
// 能够获得的最大值
// 示例1
// 输入
// 复制
// 1 2 3
// 输出
// 复制
// 9

func handle13(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	res := -1
	for i := 1; i < len(a); i++ {
		tp := handle13(a[:i]) + handle13(a[i:])
		res = max(res, tp)
		tp = handle13(a[:i]) * handle13(a[i:])
		res = max(res, tp)
	}

	return res
}
