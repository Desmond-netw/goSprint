package main

// func factorial function with Recursive

func FactorialRecursive(n int) int {
	// return nagative values
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}

	return n * FactorialRecursive(n-1)
}
