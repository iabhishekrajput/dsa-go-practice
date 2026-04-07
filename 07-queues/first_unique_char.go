package queues

// FirstUniqChar returns the index of the first non-repeating character
// in a string. Returns -1 if no unique character exists.
//
// Example:
//   FirstUniqChar("leetcode") → 0
//   FirstUniqChar("loveleetcode") → 2
//   FirstUniqChar("aabb") → -1
//
// Constraints:
//   - O(n) time
//   - String contains only lowercase English letters
func FirstUniqChar(s string) int {
	// YOUR CODE HERE

	count := make(map[byte]int, len(s))
	queue := make([]int, 0, len(s))
	chars := []byte(s)

	for i, v := range chars {
		c, ok := count[v]; if !ok {
			count[v] = 1
			queue = append(queue, i)
		} else {
			count[v] = c + 1
		}
	}
	
	for len(queue) != 0 {
		if count[chars[queue[0]]] == 1 {
			return queue[0]
		}
		queue = queue[1:]
	}

	return -1
}
