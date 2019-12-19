package top100

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
// 输出: true
// 示例 2:
//
// 输入: "()[]{}"
// 输出: true
// 示例 3:
//
// 输入: "(]"
// 输出: false
// 示例 4:
//
// 输入: "([)]"
// 输出: false
// 示例 5:
//
// 输入: "{[]}"
// 输出: true

// 思路：用栈实现，不多说，基础

type stack []byte

func NewStack() *stack {
	s := stack(make([]byte, 0))
	return &s
}

func (s *stack) Get() byte {
	return (*s)[len(*s)-1]
}

func (s *stack) Pop() byte {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack) Push(char byte) {
	*s = append(*s, char)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func isValid(s string) bool {
	st := NewStack()
	for i := 0; i < len(s); i++ {

		char := s[i]

		if char == '(' || char == '[' || char == '{' {
			st.Push(char)
		} else if char == ')' || char == ']' || char == '}' {
			if st.IsEmpty() {
				return false
			}
			switch char {
			case ')':
				if st.Pop() != '(' {
					return false
				}
			case ']':
				if st.Pop() != '[' {
					return false
				}
			case '}':
				if st.Pop() != '{' {
					return false
				}
			}
		}
	}

	if !st.IsEmpty() {
		return false
	}

	return true
}
