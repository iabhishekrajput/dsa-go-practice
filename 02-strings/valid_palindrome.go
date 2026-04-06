package strings

import (
	"unicode"
)

// IsPalindrome checks if a string is a palindrome,
// considering only alphanumeric characters and ignoring case.
//
// Example:
//
//	IsPalindrome("A man, a plan, a canal: Panama") → true
//	IsPalindrome("race a car") → false
//
// Constraints:
//   - O(n) time, O(1) space
//   - Only consider alphanumeric characters (a-z, A-Z, 0-9)
//
// Hint: Two pointers from both ends. Skip non-alphanumeric chars.
//
//	Compare lowercase versions of the characters.
//	Check out unicode.IsLetter, unicode.IsDigit, unicode.ToLower.
func IsPalindrome(s string) bool {
	// YOUR CODE HERE

	left, right := 0, len(s)-1
	for left < right {
		if !isAlphanumeric(rune(s[left])) {
			left++
			continue
		}

		if !isAlphanumeric(rune(s[right])) {
			right--
			continue
		}

		if toLower(rune(s[left])) != toLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func toLower(r rune) rune {
	return unicode.ToLower(r)
}
