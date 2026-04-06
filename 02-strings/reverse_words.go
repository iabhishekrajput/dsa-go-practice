package strings

import "strings"

// ReverseWords reverses the order of words in a string.
// Words are separated by spaces. Leading/trailing spaces are removed.
// Multiple spaces between words are reduced to a single space.
//
// Example:
//   ReverseWords("  the sky  is   blue  ") → "blue is sky the"
//   ReverseWords("hello world") → "world hello"
//
// Hint: Split the string into words, then build the result backwards.
//       Or try the clever way: reverse the entire string, then reverse each word.
func ReverseWords(s string) string {
	// YOUR CODE HERE

	words := strings.Split(s, " ")
	
	left := 0
	for right := range words {
		if words[right] != "" {
			words[left] = words[right]
			left++
		}
	}

	words = words[:left]

	for left, right := 0, len(words) - 1; left < right; left, right = left + 1, right - 1 {
		words[left], words[right] = words[right], words[left]
	}


	return strings.Join(words, " ")
}
