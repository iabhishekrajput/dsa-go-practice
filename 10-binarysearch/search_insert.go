package binarysearch

// SearchInsert returns the index of target in a sorted array,
// or the index where it would be inserted to maintain sorted order.
func SearchInsert(nums []int, target int) int {
	// Your solution here

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
