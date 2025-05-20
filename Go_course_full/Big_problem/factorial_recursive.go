package main

// func factorial function with Recursive

func FactorialRecursive(n int) int {
	// return nagative values
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1 // return one for factorial of 0
	}

	result := FactorialRecursive(n - 1)

	if result > (1<<63-1)/n {
		return 0
	}

	return n * result
}
