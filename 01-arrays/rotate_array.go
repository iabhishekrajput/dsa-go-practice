package arrays

// RotateArray rotates the array k steps to the right, in-place.
// Must use O(1) extra space.
func RotateArray(nums []int, k int) {
	// Your solution here

	k = k % len(nums)

	left := 0
	right := len(nums) - 1
	reverse(nums, left, right)

	right = k - 1
	reverse(nums, left, right)

	left = k
	right = len(nums) - 1
	reverse(nums, left, right)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}