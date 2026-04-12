package binarysearch

// FirstBadVersion returns the first bad version in the range [1, n].
// isBadVersion returns true if the given version is bad.
func FirstBadVersion(n int, isBadVersion func(int) bool) int {
	// Your solution here

	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2

		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
