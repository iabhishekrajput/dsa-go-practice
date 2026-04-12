package stacks

// NextGreaterElement finds the next greater element for each nums1[i] in nums2.
// nums1 is a subset of nums2. Both have no duplicates.
// Returns -1 if no greater element exists.
func NextGreaterElement(nums1, nums2 []int) []int {
	// Your solution here

	stack := []int{}
	valueMap := make(map[int]int, len(nums2))

	for _, val := range nums2 {
		for len(stack) > 0 && val > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			valueMap[top] = val
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, val)
	}

	result := []int{}
	for _, value := range nums1 {
		nextMax, ok := valueMap[value]; if !ok {
			result = append(result, -1)
		} else {
			result = append(result, nextMax)
		}
	}

	return result
}
