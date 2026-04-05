package arrays

// LengthOfLongestSubstring returns the length of the longest substring
// without repeating characters.
//
// Example:
//   LengthOfLongestSubstring("abcabcbb") → 3
//   LengthOfLongestSubstring("bbbbb") → 1
//
// Constraints:
//   - O(n) time
//
// Hint: Use a map[byte]int to store the last index where each char was seen.
//       Maintain a 'left' pointer for the window start.
//       When you see a repeat, jump 'left' forward (but never backward!).
func LengthOfLongestSubstring(s string) int {
	// YOUR CODE HERE

	if len(s) < 2 {
		return len(s)
	}

	seenCount := make(map[byte]int)
	left, maxLen := 0, 1

	seenCount[s[left]] = left

	for right := 1; right < len(s); right++ {
		val, ok := seenCount[s[right]]; if ok && val >= left {
			left = val + 1
		}
		seenCount[s[right]] = right
		maxLen = max(maxLen, right - left + 1)
	}

	return maxLen
}
