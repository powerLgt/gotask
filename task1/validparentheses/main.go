package main

import "fmt"

/*
   给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
   有效字符串需满足：
   1. 左括号必须用相同类型的右括号闭合。
   2. 左括号必须以正确的顺序闭合。
   3. 每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	stack := make([]byte, 0)
	matchMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := range s {
		match, ok := matchMap[s[i]]
		if ok && (len(stack) == 0 || stack[len(stack)-1] != match) {
			return false
		}

		if ok {
			if len(stack) == 0 || stack[len(stack)-1] != match {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}

	}
	return len(stack) == 0
}

func main() {
	inputs := []string{
		"()", "()[]{}", "(]", "([])",
	}

	for _, input := range inputs {
		fmt.Printf("%s isValid: %v \n", input, isValid(input))
	}

}
