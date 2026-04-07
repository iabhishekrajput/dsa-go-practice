package hashmaps

import "sort"

// GroupAnagrams groups a list of strings so that anagrams are together.
// The order of groups and order within groups does not matter.
//
// Example:
//   GroupAnagrams(["eat","tea","tan","ate","nat","bat"])
//   → [["eat","tea","ate"], ["tan","nat"], ["bat"]]
//
// Constraints:
//   - O(n * k log k) time where n = number of strings, k = max string length
//
// Hint: Anagrams have the same letters when sorted.
//       Use sorted string as the map key → map[string][]string.
//       sort.Slice works on []byte.
func GroupAnagrams(strs []string) [][]string {
	// YOUR CODE HERE

	anagrams := make(map[string][]string, len(strs))

	for _, v := range strs {
		word := []byte(v)

		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})

		sortedWord := string(word)
		anagram, ok := anagrams[sortedWord]; if !ok {
			anagrams[sortedWord] = []string{v}
		} else {
			anagrams[sortedWord] = append(anagram, v)
		}
	}

	var result [][]string

	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}
