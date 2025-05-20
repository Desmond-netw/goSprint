package main

// to the power function

func ToThePowerIterative(n int, power int) int {

	if power == 0 {
		return 0
	}

	result := 0

	for i := 0; i <= power; i++ {
		result = n * i

	}
	return result
}
