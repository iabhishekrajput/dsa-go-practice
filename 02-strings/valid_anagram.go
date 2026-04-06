package strings


// IsAnagram checks if string t is an anagram of string s.
// Both strings contain only lowercase English letters.
//
// Example:
//   IsAnagram("anagram", "nagaram") → true
//   IsAnagram("rat", "car") → false
//
// Constraints:
//   - O(n) time
//
// Hint: Count character frequencies. You can use a map[byte]int
//       or a [26]int array (since only lowercase letters).
//       Increment for s, decrement for t. All counts should be 0.
func IsAnagram(s string, t string) bool {
	// YOUR CODE HERE

	if (len(s) != len(t)) {
		return false
	}

	counts := make(map[rune]int, 26)
	for _, v := range s {
		counts[v]++
	}

	for _, v := range t {
		counts[v]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
