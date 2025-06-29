package main

// function

/*
-- Methods
step 1. get the numb from numbers (find smallest number among the nonZero numbers )
step 2. substract numb from all elements of non zero element of array
step 3. return the number of operations perform if 2 or 3, or 5 times.
*/

func minimumOperation(numbers []int) int {
	// store uniqnumbers
	uniqueNumbers := make(map[int]struct{})

	// iterate over numbers for individual numbers
	for _, num := range numbers {
		// check if number > 0
		if num > 0 {
			uniqueNumbers[num] = struct{}{}
		}

	}
	return len(uniqueNumbers)
}
