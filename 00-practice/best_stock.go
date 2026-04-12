package practice

// MaxProfit returns the maximum profit from buying and selling a stock once.
// You must buy before you sell. If no profit is possible, return 0.
//
// Example:
//   MaxProfit([]int{7, 1, 5, 3, 6, 4}) → 5
//   MaxProfit([]int{7, 6, 4, 3, 1}) → 0
//
// Constraints:
//   - O(n) time, O(1) space
func MaxProfit(prices []int) int {
	// YOUR CODE HERE

	minSoFar := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minSoFar {
			minSoFar = prices[i]
		}

		profit := prices[i] - minSoFar
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
