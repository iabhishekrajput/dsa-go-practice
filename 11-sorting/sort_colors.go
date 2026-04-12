package sorting

// SortColors sorts an array of 0s, 1s, and 2s in-place in a single pass.
func SortColors(nums []int) {
	// Your solution here
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}
