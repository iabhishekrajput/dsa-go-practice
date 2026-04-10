package stacks

import (
	"strconv"
)

// EvalRPN evaluates an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, /. Division truncates toward zero.
// You can assume the expression is always valid.
func EvalRPN(tokens []string) int {
	// Your solution here

	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
			continue
		}

		num1 := stack[len(stack) - 1]
		num2 := stack[len(stack) - 2]

		stack = stack[:len(stack) - 2]
		
		switch token {
		case "+":
			stack = append(stack, num2 + num1)
		case "-":
			stack = append(stack, num2 - num1)
		case "*":
			stack = append(stack, num2 * num1)
		case "/":
			stack = append(stack, num2 / num1)
		}
	}

	return stack[len(stack) - 1]
}
