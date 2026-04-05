package arrays

// MaxArea returns the maximum water a container can hold,
// formed by two lines from the height slice and the x-axis.
//
// Area = min(height[left], height[right]) * (right - left)
//
// Example:
//   MaxArea([]int{1,8,6,2,5,4,8,3,7}) → 49
//
// Constraints:
//   - O(n) time, O(1) space
//
// Hint: Start with left=0, right=len-1 (widest container).
//       Move the pointer pointing to the shorter line inward.
//       Why? The height is limited by the shorter line.
//       Moving the taller one can only reduce width without helping height.
func MaxArea(height []int) int {
	// YOUR CODE HERE

	left, right, maxArea := 0, len(height) - 1, 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}
	}

	return maxArea
}

func min[T int](a, b T) T {
	if a < b {
		return a
	}

	return b
}
