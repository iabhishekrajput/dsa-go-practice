package recursion

// Power computes x raised to the power n.
// n can be negative. Uses fast exponentiation (O(log n)).
func Power(x float64, n int) float64 {
	// Your solution here

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / Power(x, -1*n)
	}

	if n%2 == 0 {
		halfPower := Power(x, n/2)
		return halfPower * halfPower
	}

	return x * Power(x, (n-1))
}
