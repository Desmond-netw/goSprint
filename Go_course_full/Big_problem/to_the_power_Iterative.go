package main

// to the power function

func ToThePowerIterative(n int, power int) int {

	// handle special cases
	// treate 0^0 as 1

	if power == 0 {
		return 1
	}

	// handle negative power
	if power < 0 {
		return 0
	}

	// results initila value to 1
	result := 1
	for i := 0; i < power; i++ {
		// check before multiplying result
		result *= n
	}

	return result

}
