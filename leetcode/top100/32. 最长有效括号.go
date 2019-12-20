package top100

// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
// 示例 1:
//
// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"
// 示例 2:
//
// 输入: ")()())"
// 输出: 4
// 解释: 最长有效括号子串为 "()()"

// 思路1： 用一个栈记录下标，先压入-1（便于正确得出长度），当遇到(时入栈，当遇到)出栈，如果栈不为空，将
// 当前下标与栈顶索引相减为现在的长度，如果栈为空则入栈
type stackInt []int

func newStack() *stackInt {
	s := stackInt(make([]int, 0))
	return &s
}

func (s *stackInt) Push(char int) {
	*s = append(*s, char)
}

func (s *stackInt) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stackInt) Get() int {
	return (*s)[len(*s)-1]
}

func (s *stackInt) IsEmpty() bool {
	return len(*s) == 0
}

func longestValidParenthesesFunc1(s string) int {
	st := newStack()
	result := 0
	st.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st.Push(i)
		} else if s[i] == ')' {

			st.Pop()
			if st.IsEmpty() {
				st.Push(i)
			} else {
				length := i - st.Get()
				if length > result {
					result = length
				}
			}
		}
	}

	return result
}

// 思路2：设置两个计数器Left和Right，分别计数左括号和右括号的个数，从左往右遍历，当右括号大于左
// 括号的时候清零，当右括号等于左括号的时候计数器相加即为有效长度。然后清零计数器再从右往左遍历，
// 当左括号大于右括号的时候清零，当左括号等于右括号的时候计数器相加即为有效长度

func longestValidParenthesesFunc2(s string) int {
	result := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			left++
		case ')':
			right++
		}

		if right > left {
			left, right = 0, 0
		} else if right == left {
			temp := left + right
			if temp > result {
				result = temp
			}
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			left++
		case ')':
			right++
		}

		if right < left {
			left, right = 0, 0
		} else if right == left {
			temp := left + right
			if temp > result {
				result = temp
			}
		}
	}
	return result
}
