package main

func FactorialIterative(n int) int {
	// check for Negative values

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	// initiate result to 1
	result := 1

	for i := 1; i <= n; i++ {
		// check for overflow
		if result > (1<<63-1)/i {
			return 0
		}
		result *= i
	}

	return result
}
