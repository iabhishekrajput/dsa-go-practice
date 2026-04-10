package stacks

import (
	"strings"
	"unicode"
)

// DecodeString decodes an encoded string where k[encoded] means
// the encoded string inside brackets is repeated k times.
// Example: "3[a2[c]]" -> "accaccacc"
type StringDecoder struct {
	current string
	count   int
}

func DecodeString(s string) string {
	// Your solution here

	stack := []StringDecoder{}

	count := 0
	current := ""

	for _, currChar := range s {
		if unicode.IsDigit(currChar) {
			count = count*10 + int(currChar - '0')
		} else if unicode.IsLetter(currChar) {
			current += string(currChar)
		} else if currChar == '[' {
			stringDecoder := StringDecoder{
				count:   count,
				current: current,
			}
			stack = append(stack, stringDecoder)
			count = 0
			current = ""
		} else if currChar == ']' {
			top := stack[len(stack) - 1]
			current = top.current + strings.Repeat(current, top.count)
			stack = stack[:len(stack) - 1]
		}
	}

	return current
}
