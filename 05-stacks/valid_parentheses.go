package stacks

// IsValidParentheses checks if a string of brackets is valid.
// Each open bracket must be closed by the same type in the correct order.
//
// Example:
//
//	IsValidParentheses("()[]{}") → true
//	IsValidParentheses("([)]") → false
//	IsValidParentheses("{[]}") → true
//
// Constraints:
//   - O(n) time, O(n) space
//
// Hint: Use a slice as a stack.
//
//	Push opening brackets. On closing brackets, pop and check match.
//	A map from closing → opening bracket keeps it clean.
func IsValidParentheses(s string) bool {
	// YOUR CODE HERE

	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0, len(s))

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			lastChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if lastChar != match[char] {
				return false
			}

		}
	}

	return len(stack) == 0
}
